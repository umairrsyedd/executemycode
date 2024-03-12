package handlers

import (
	"executemycode/internal/client"
	"executemycode/internal/container"
	"executemycode/internal/executer"
	"executemycode/pkg/message"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func ConnectionHandler(clientManager *client.ClientRegistry, containerManager *container.ContainerOrc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		client, err := clientManager.NewClient(w, r, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not establish client connection : %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Printf("Client %s connected\n", client.Id)

		go listenForMessages(client, containerManager)
	}
}

func listenForMessages(client *client.Client, containerOrc *container.ContainerOrc) {
	defer client.CloseConnection()
	for {
		_, rawMessage, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err) || websocket.IsUnexpectedCloseError(err) {
				log.Printf("client connection closed for client %s: %v", client.Id, err)
			} else {
				log.Printf("error reading message from client %s: %v", client.Id, err)
			}
			return
		}

		msg, err := message.DecodeMessage(rawMessage)
		if err != nil {
			log.Printf("error decoding message: %s", err)
			return
		}

		err = msg.Validate()
		if err != nil {
			errMsg := message.Message{
				Type:    message.Error,
				Message: err.Error(),
			}

			encodedErrMessage, _ := message.EncodeMessage(errMsg)
			client.Write([]byte(encodedErrMessage))
			continue
		}

		switch msg.Type {
		case message.Code:
			if client.IsExecuting() {
				prevExecution, _ := client.GetExecution()
				prevExecution.Done()
			}

			newExecution := executer.NewExecution(msg.Language, msg.Message, client)
			client.SetExecution(newExecution)
			go containerOrc.ConnectAndExecute(newExecution)
			go newExecution.ListenForOutput()

		case message.Input:
			execution, err := client.GetExecution()
			if err != nil {
				fmt.Println(err)
				continue
			}
			go execution.FeedInput(msg.Message)

		case message.Close:
			execution, err := client.GetExecution()
			if err != nil {
				fmt.Println(err)
				continue
			}
			execution.Done()
		}

	}
}
