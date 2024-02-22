package handlers

import (
	"executemycode/internal/bridge"
	"executemycode/internal/executer"
	"executemycode/internal/socket"

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

	program := executer.NewProgram(client.Id, executer.Language(language))
	handler := bridge.New(client, program)

	handler.Start()

	log.Printf("Client %s disconnected", client.Id)
}
