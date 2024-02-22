// The bridge service connects the client connection with the executer service

package bridge

import (
	"executemycode/internal/executer"
	"executemycode/internal/socket"
	"fmt"
	"log"
)

type Bridge struct {
	Client  *socket.Client
	Program *executer.Program
}

func New(client *socket.Client, program *executer.Program) Bridge {
	return Bridge{
		Client:  client,
		Program: program,
	}
}

func (b *Bridge) Start() {
	go b.Client.ReadMessages(b.ClientMessageHandler())

	for {
		select {
		case isClosed := <-b.Client.Closed:
			if isClosed {
				log.Printf("Client %s disconnected", b.Client.Id)
				b.Client.CloseConnection()
			}

		case executionComplete := <-b.Program.Completed:
			if executionComplete {
				b.Client.Closed <- true
			}

		case output, ok := <-b.Program.OutputChan:
			if !ok {
				fmt.Println("Output Channel Closed")
				return
			}
			fmt.Printf("%v", output)
			if outputStr, ok := output.(string); ok {
				outputBytes := []byte(outputStr)

				b.Client.WriteMessage(outputBytes)
			} else {
				fmt.Println("Output is not a string")
			}

		case err, ok := <-b.Program.ErrorChan:
			if !ok {
				fmt.Println("Error Channel Closed")
				return
			}
			fmt.Println("Error Channel Received Something")
			fmt.Printf("%v", err)
		}
	}
}

func (b *Bridge) ClientMessageHandler() func(message socket.Message) {
	return func(message socket.Message) {
		switch message.Type {
		case socket.Code:
			b.Program.SetCode(message.Message)
			go b.Program.Execute()
		case socket.Input:
			b.Program.InputChan <- message.Message
		}
	}
}
