package main

import (
	"executemycode/api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/execute", handlers.ExecuteHandler)
	http.ListenAndServe(":6364", nil)
}
