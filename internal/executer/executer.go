package executer

import (
	"executemycode/pkg/message"
	"fmt"
	"io"
	"log"
)

type Executer interface {
	StartExecution() error
	FeedInput() error
}

type Execution struct {
	ExecId        int
	ExecutionInfo ExecutionInfo
	InputChan     chan string
	OutputChan    chan string
	OutputWriter  io.Writer
	ExitCode      chan int
}

type ExecutionInfo struct {
	SourceCode      string
	ProgramLanguage ProgramLanguage
	FileExtension   string
	Cmd             []string
}

func NewExecution(execId int, language string, code string, outputWriter io.Writer) *Execution {
	return &Execution{
		ExecId: execId,
		ExecutionInfo: ExecutionInfo{
			ProgramLanguage: ProgramLanguage(language),
			SourceCode:      code,
			FileExtension:   getFileExtension(ProgramLanguage(language)),
			Cmd:             getCmd(ProgramLanguage(language)),
		},
		InputChan:    make(chan string),
		OutputChan:   make(chan string),
		OutputWriter: outputWriter,
		ExitCode:     make(chan int),
	}
}

func (e *Execution) FeedInput(input string) {
	e.InputChan <- input
}

func (e *Execution) Listen() {
	for {
		select {
		case output := <-e.OutputChan:
			msg := message.Message{
				ExecutionId: e.ExecId,
				Type:        message.Output,
				Message:     output,
			}
			e.SendMessage(msg)
		case exitCode := <-e.ExitCode:
			msg := message.Message{
				ExecutionId: e.ExecId,
				Type:        message.Done,
				Message:     fmt.Sprintf("...Program finished with exit code %d", exitCode),
			}
			e.SendMessage(msg)
			return
		}
	}

}

func (e *Execution) SendMessage(msg message.Message) {
	messageToSend, err := message.EncodeMessage(msg)
	if err != nil {
		log.Printf("error encoding message to send to client: %v", err)
	}

	_, err = e.OutputWriter.Write([]byte(messageToSend))
	if err != nil {
		log.Printf("Couldn't Forward Output for Execution Id: %d", e.ExecId)
	}

}

func (e *Execution) Read(p []byte) (n int, err error) {
	select {
	case input, ok := <-e.InputChan:
		if !ok {
			return 0, io.EOF
		}
		return copy(p, []byte(input)), nil
	default:
		return 0, nil
	}
}
