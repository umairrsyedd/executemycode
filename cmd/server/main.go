package main

import (
	"executemycode/api/handlers"
	_ "executemycode/internal/socket"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func main() {
	http.HandleFunc("/execute", handlers.ExecuteHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), nil)
}
