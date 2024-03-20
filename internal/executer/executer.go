package executer

import (
	execlangauge "executemycode/internal/executer/exec_language"
	"executemycode/pkg/language"
	"executemycode/pkg/message"
	"fmt"
	"io"
	"log"
)

type Execution struct {
	ExecutionInfo ExecutionInfo
	InputChan     chan string
	OutputChan    chan string
	StopChan      chan bool
	OutputWriter  io.Writer
	ExitCode      chan int
	IsDone        bool
}

type ExecutionInfo struct {
	SourceCode   string
	LangExecuter execlangauge.LanguageExecuter
}

func NewExecution(lang string, code string, outputWriter io.Writer) *Execution {
	return &Execution{
		ExecutionInfo: ExecutionInfo{
			SourceCode:   code,
			LangExecuter: execlangauge.New(language.ProgramLanguage(lang)),
		},
		OutputWriter: outputWriter,
		InputChan:    make(chan string),
		OutputChan:   make(chan string),
		ExitCode:     make(chan int),
		StopChan:     make(chan bool),
	}
}

func (e *Execution) FeedInput(input string) {
	e.InputChan <- input
}

func (e *Execution) ListenForOutput() {
	for {
		select {
		case output := <-e.OutputChan:
			msg := message.Message{
				Type:    message.Output,
				Message: output,
			}
			e.SendMessage(msg)
		case exitCode := <-e.ExitCode:
			msg := message.Message{
				Type:    message.Done,
				Message: fmt.Sprintf("...Program finished with exit code %d", exitCode),
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
		log.Printf("Couldn't Forward Output for Execution")
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

func (e *Execution) Done() {
	e.StopChan <- true
	close(e.InputChan)
	close(e.OutputChan)
	close(e.ExitCode)
	e.IsDone = true
}

func (e *Execution) Clear() {
	if e.IsDone {
		return
	}
	close(e.InputChan)
	close(e.OutputChan)
	close(e.ExitCode)
	close(e.StopChan)
}
