package message

import (
	"encoding/json"
	"fmt"
)

type MessageType string

const (
	Code   MessageType = "code"
	Input  MessageType = "input"
	Output MessageType = "output"
	Error  MessageType = "error"
	Close  MessageType = "close"
)

type Message struct {
	ProgramId *int        `json:"programId"`
	Type      MessageType `json:"type"`
	Message   string      `json:"message"`
	Language  string      `json:"language"`
}

func DecodeMessage(rawMessage []byte) (decodedMessage Message, err error) {
	err = json.Unmarshal(rawMessage, &decodedMessage)
	if err != nil {
		return Message{}, fmt.Errorf("error decoding message: %v", err)
	}

	return decodedMessage, nil
}

func EncodeMessage(messageType MessageType, message string) (encodedMessage []byte, err error) {
	newMessage := Message{
		Type:    messageType,
		Message: message,
	}

	encodedMessage, err = json.Marshal(newMessage)
	if err != nil {
		return nil, err
	}

	return encodedMessage, nil
}

func (m Message) IsCode() bool {
	return m.Type == Code
}

func (m Message) IsInput() bool {
	return m.Type == Input
}

func (m Message) IsClose() bool {
	return m.Type == Close
}
