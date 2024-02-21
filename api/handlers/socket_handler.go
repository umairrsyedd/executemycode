package handlers

import (
	"executemycode/internal/executer"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client represents a WebSocket client with a unique ID.
type Client struct {
	Id   uuid.UUID
	Conn *websocket.Conn
}

func ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	language := r.URL.Query().Get("language")
	if language == "" {
		http.Error(w, "Language Parameter is required", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is empty", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not establish socket connection : %v", err), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := &Client{
		Id:   uuid.New(),
		Conn: conn,
	}

	program := executer.New(client.Id, executer.Language(language), code)
	err = program.Execute()
	if err != nil {
		errMsg := fmt.Sprintf("Error executing program: %v", err)
		if writeErr := conn.WriteMessage(websocket.TextMessage, []byte(errMsg)); writeErr != nil {
			log.Printf("Error sending error message over WebSocket: %v", writeErr)
			return
		}
		return
	}

	log.Printf("Client %s connected", client.Id)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Client %s disconnected", client.Id)
			return
		}
		// Process the message independently for each client.
		processMessage(client, messageType, p)
	}
}

// Process the message for an individual client.
func processMessage(client *Client, messageType int, message []byte) {
	log.Printf("Received message from Client %s: %s", client.Id, string(message))

	// Respond to the client if needed.
	responseMessage := []byte("Message received")
	if err := client.Conn.WriteMessage(messageType, responseMessage); err != nil {
		log.Println(err)
	}
}
