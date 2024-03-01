package main

import (
	"context"
	"executemycode/api/handlers"
	"executemycode/internal/client"
	"executemycode/internal/container"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	containerCount, _ := strconv.Atoi(os.Getenv("CONTAINER_COUNT"))
	execWait, _ := strconv.Atoi(os.Getenv("MAX_WAIT_TIME_MINS"))
	execWaitTime := execWait * int(time.Second)

	execTimeout, _ := strconv.Atoi(os.Getenv("EXEC_TIMEOUT_MINS"))
	execTimeOutTime := execTimeout * int(time.Second)

	containerOrchestrator, err := container.New(ctx, containerCount, time.Duration(execWaitTime), time.Duration(execTimeOutTime))
	if err != nil {
		log.Fatalf("Error Init Container Manager: %s", err)
	}

	clientRegistry := client.New()

	serverPort := os.Getenv("SERVER_PORT")

	http.HandleFunc("/connect", handlers.ConnectionHandler(clientRegistry, containerOrchestrator))

	server := &http.Server{
		Addr: ":" + serverPort,
	}

	// Graceful shutdown handling
	shutdownComplete := make(chan struct{})
	go func() {
		defer close(shutdownComplete)
		handleGracefulShutdown(ctx, containerOrchestrator, clientRegistry, server)
	}()

	log.Printf("Server is starting on :%s", serverPort)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server ListenAndServe error: %v", err)
	}

	<-shutdownComplete // Wait for shutdown completion
}

func handleGracefulShutdown(ctx context.Context, containerOrchestrator *container.ContainerOrc, clientRegistry *client.ClientRegistry, server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Println("Server is shutting down...")

	containerOrchestrator.Cleanup(ctx)
	clientRegistry.CleanUp()

	server.SetKeepAlivesEnabled(false)

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	log.Println("Server has been gracefully stopped")

	select {
	case <-ctx.Done():
	default:
		ctx.Done()
	}
}
