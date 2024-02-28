package program

import (
	"executemycode/utils"
	"fmt"
	"log"
)

type ProgramStatus string

const (
	NotStarted ProgramStatus = "NotStarted"
	Executing  ProgramStatus = "Executing"
	Terminated ProgramStatus = "Terminated"
)

type Program struct {
	Id         int
	Language   Language
	Code       string
	Status     ProgramStatus
	InputChan  chan []byte
	OutputChan chan []byte
	ErrorChan  chan []byte
}

func New(language Language, code string) (*Program, error) {
	if language == "" {
		return nil, fmt.Errorf("language cannot be empty")
	}
	return &Program{
		Id:         utils.GenerateRandom4DigitNumber(),
		Language:   language,
		Code:       code,
		Status:     NotStarted,
		InputChan:  make(chan []byte),
		OutputChan: make(chan []byte),
		ErrorChan:  make(chan []byte),
	}, nil
}

func (pr Program) Read(p []byte) (n int, err error) {
	log.Printf("Output is %v", p)
	select {
	case output := <-pr.OutputChan:
		fmt.Printf("Received Output: %v", output)
		return n, nil
	case err := <-pr.ErrorChan:
		fmt.Printf("Received Error: %v", err)
		return n, nil
	}
}

func (pr Program) Write(p []byte) (n int, err error) {
	pr.InputChan <- p
	fmt.Printf("Wrote Input: %v", p)
	return len(p), nil
}

func (pr Program) Close() error {
	close(pr.InputChan)
	close(pr.OutputChan)
	close(pr.ErrorChan)
	return nil
}
