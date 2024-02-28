package handlers

import (
	"context"
	"executemycode/internal/client"
	"executemycode/internal/container"
	"executemycode/internal/program"
	"executemycode/pkg/message"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func ConnectionHandler(clientManager *client.ClientRegistry, containerManager *container.ContainerOrc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqContext := r.Context()

		client, err := clientManager.NewClient(w, r, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not establish client connection : %v", err), http.StatusInternalServerError)
			return
		}
		log.Printf("Client %s connected", client.Id)

		go listenForMessages(reqContext, client, containerManager)
	}
}

func listenForMessages(ctx context.Context, client *client.Client, containerOrchestrator *container.ContainerOrc) {
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

		msg, err := message.DecodeMessage(rawMessage)
		if err != nil {
			log.Printf("error decoding message: %s", err)
		}

		switch msg.Type {
		case message.Code:
			ctx := context.TODO()
			newProgram, err := program.New(program.Language(msg.Language), msg.Message)
			if err != nil {
				log.Printf("error creating new program: %s", err)
				break
			}
			newProgram.Status = program.Executing

			container, err := containerOrchestrator.GetAvailableContainer(ctx)
			if err != nil {
				log.Printf("error finding container: %s", err)
				break
			}

			err = container.Execute(ctx, newProgram, msg.Message)
			if err != nil {
				log.Printf("error executing in container: %s", err)
				break
			}

			// containerOrchestrator.ReturnContainer(container)

		}

	}
}
