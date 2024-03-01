package handlers

import (
	"context"
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
		fmt.Printf("Client %s connected", client.Id)

		go listenForMessages(client, containerManager)
	}
}

func listenForMessages(client *client.Client, containerOrc *container.ContainerOrc) {
	defer client.CloseConnection()
	for {
		_, rawMessage, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err) {
				log.Printf("client connection closed for client %s: %v", client.Id, err)
			} else {
				log.Printf("error reading message from client %s: %v", client.Id, err)
			}
			return
		}

		msgContext := context.TODO()

		msg, err := message.DecodeMessage(rawMessage)
		if err != nil {
			log.Printf("error decoding message: %s", err)
			return
		}

		switch msg.Type {
		case message.Code:
			newExecution := executer.NewExecution(msg.ExecutionId, msg.Language, msg.Message, client)
			client.AddExecution(newExecution)
			go containerOrc.ConnectAndExecute(msgContext, newExecution)
			go newExecution.Listen()

		case message.Input:
			execution, err := client.GetExecution(msg.ExecutionId)
			if err != nil {
				fmt.Println(err)
				continue
			}
			go execution.FeedInput(msg.Message)
		}

	}
}
