package executer

import (
	"bufio"
	"executemycode/internal/common"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

type Program struct {
	Id         uuid.UUID
	State      ProgramState
	Language   Language
	Code       string
	InputChan  chan any
	OutputChan chan any
	ErrorChan  chan any
	Completed  chan bool
}

func NewProgram(id uuid.UUID, language Language) *Program {
	return &Program{
		Id:         id,
		State:      Idle,
		Language:   language,
		Code:       "",
		InputChan:  make(chan any),
		OutputChan: make(chan any),
		ErrorChan:  make(chan any),
	}
}

func (p *Program) SetCode(code string) {
	p.Code = code
}

func (p *Program) HandleMessage(message common.Message) {
	switch p.State {
	case Running:
		if message.IsCode() {
			return
		} else if message.IsInput() {
			p.InputChan <- message.Message
			p.OutputChan <- message.Message
		}
	case Idle, Terminated:
		if message.IsCode() {
			p.Code = message.Message
			go p.Execute()
			p.State = Running
		}
	}

}

func (p *Program) Execute() {
	if p.Code == "" {
		p.ErrorChan <- fmt.Errorf("no code set for execution")
		return
	}

	file, cmd, err := p.prepare(p.Language, p.Code)
	if err != nil {
		p.ErrorChan <- err.Error()
		return
	}
	defer os.Remove(file.Name())

	stdinPipe, stdoutPipe, stderrPipe, err := p.setupPipes(cmd)
	if err != nil {
		p.ErrorChan <- err.Error()
		return
	}

	err = cmd.Start()
	if err != nil {
		p.ErrorChan <- err.Error()
		return
	}

	go sendInputToProgram(stdinPipe, p.InputChan)
	go captureOutputFromProgram(stdoutPipe, p.OutputChan)
	go captureOutputFromProgram(stderrPipe, p.ErrorChan)

	err = cmd.Wait()
	if err != nil {
		p.ErrorChan <- err.Error()
		return
	}

	fmt.Printf("\nProgram %v finished executing\n", p.Id)
	p.State = Terminated
}

func (p *Program) prepare(language Language, code string) (file *os.File, cmd *exec.Cmd, err error) {
	execInfo := getExecInfo(language)
	file, err = os.CreateTemp("", fmt.Sprintf("program*.%s", execInfo.FileExtension))
	if err != nil {
		return nil, nil, err
	}

	_, err = file.Write([]byte(code))
	if err != nil {
		return nil, nil, err
	}

	cmd = exec.Command(
		execInfo.ExecName,
		append(execInfo.ExecArgs, file.Name())...,
	)

	return file, cmd, err
}

func (p *Program) setupPipes(cmd *exec.Cmd) (stdinPipe io.WriteCloser, stdoutPipe io.ReadCloser, stderrPipe io.ReadCloser, err error) {
	stdinPipe, err = cmd.StdinPipe()
	if err != nil {
		return nil, nil, nil, err
	}

	stdoutPipe, err = cmd.StdoutPipe()
	if err != nil {
		return nil, nil, nil, err
	}

	stderrPipe, err = cmd.StderrPipe()
	if err != nil {
		return nil, nil, nil, err
	}

	return stdinPipe, stdoutPipe, stderrPipe, nil

}

func sendInputToProgram(pipe io.Writer, channel chan any) {
	for input := range channel {
		_, err := fmt.Fprintln(pipe, input)
		if err != nil {
			fmt.Printf("Error writing input to pipe: %v\n", err)
			break
		}
	}
}

func captureOutputFromProgram(pipe io.ReadCloser, channel chan any) {
	scanner := bufio.NewScanner(pipe)
	scanner.Split(bufio.ScanLines)

	go func() {
		for scanner.Scan() {
			output := scanner.Text()
			channel <- output
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading from pipe: %v\n", err)
		}
	}()
}
