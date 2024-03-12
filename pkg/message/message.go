package message

import (
	"encoding/json"
	"fmt"
)

type MessageType string

const (
	// Sent From Client
	Code  MessageType = "code"
	Input MessageType = "input"
	Close MessageType = "close"

	// Sent From Server
	Output MessageType = "output"
	Done   MessageType = "done"
	Error  MessageType = "error"
)

type Message struct {
	Type     MessageType `json:"type"`
	Message  string      `json:"message,omitempty"`
	Language string      `json:"language,omitempty"`
}

func DecodeMessage(rawMessage []byte) (decodedMessage Message, err error) {
	err = json.Unmarshal(rawMessage, &decodedMessage)
	if err != nil {
		return Message{}, fmt.Errorf("error decoding message: %v", err)
	}

	return decodedMessage, nil
}

func EncodeMessage(newMessage Message) (encodedMessage []byte, err error) {
	encodedMessage, err = json.Marshal(newMessage)
	if err != nil {
		return nil, err
	}

	return encodedMessage, nil
}

func (m *Message) Validate() error {
	if m.Type == "" {
		return fmt.Errorf("message type must be present")
	}

	if m.Message == "" {
		return fmt.Errorf("message cannot be empty")
	}

	switch m.Type {
	case Code:
		if m.Language == "" {
			return fmt.Errorf("langauge must be present for code message")
		}
	}

	return nil
}
