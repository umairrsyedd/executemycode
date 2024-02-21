package handlers

import (
	"executemycode/internal/executer"
	"executemycode/internal/socket"
	"time"

	"fmt"
	"log"
	"net/http"
)

func ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	language := r.URL.Query().Get("language")
	if language == "" {
		http.Error(w, "Language Parameter is required", http.StatusBadRequest)
		return
	}

	client, err := socket.NewClient(w, r, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not establish socket connection : %v", err), http.StatusInternalServerError)
		return
	}
	defer client.CloseConnection()

	log.Printf("Client %s connected", client.Id)

	program := executer.New(client.Id, executer.Language(language))

	go client.ReadMessages(program.InputChan)

	isClosed := <-client.Closed
	if isClosed {
		log.Printf("Client %s disconnected", client.Id)
	}

	time.Sleep(30 * time.Second)
}
