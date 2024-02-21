package executer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

type Program struct {
	Id         uuid.UUID
	Language   Language
	Code       string
	InputChan  chan any
	OutputChan chan any
	ErrorChan  chan any
}

func New(id uuid.UUID, language Language, code string) Program {
	return Program{
		Id:         id,
		Language:   language,
		Code:       code,
		InputChan:  make(chan any),
		OutputChan: make(chan any),
		ErrorChan:  make(chan any),
	}
}

func (p *Program) Execute() (err error) {
	file, cmd, err := p.prepare(p.Language, p.Code)
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	stdinPipe, stdoutPipe, stderrPipe, err := p.setupPipes(cmd)
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	go sendInputToPipe(stdinPipe, p.InputChan)
	go captureOutputFromPipe(stdoutPipe, p.OutputChan)
	go captureOutputFromPipe(stderrPipe, p.ErrorChan)

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error during cmd wait stage " + err.Error())
		return err
	}

	fmt.Printf("\nProgram %v finished executing", p.Id)

	return
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

func sendInputToPipe(pipe io.Writer, channel chan any) {
	for input := range channel {
		_, err := fmt.Fprint(pipe, input)
		if err != nil {
			fmt.Printf("Error writing input to pipe: %v\n", err)
			break
		}
	}
}

func captureOutputFromPipe(pipe io.ReadCloser, channel chan any) {
	scanner := bufio.NewScanner(pipe)
	go func() {
		for scanner.Scan() {
			output := scanner.Text()
			fmt.Printf("\n%v", output)
		}
	}()
}
