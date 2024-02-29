package client

import (
	"executemycode/internal/executer"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id         uuid.UUID
	Conn       *websocket.Conn
	Executions map[int]*executer.Execution
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func new(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (client *Client, err error) {
	conn, err := upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		Id:         uuid.New(),
		Conn:       conn,
		Executions: make(map[int]*executer.Execution),
	}, nil
}

func (c *Client) AddExecution(execution *executer.Execution) {
	c.Executions[execution.ExecId] = execution
}

func (c *Client) GetExecution(execId int) (*executer.Execution, error) {
	execution, exists := c.Executions[execId]
	if !exists {
		return nil, fmt.Errorf("execution with Id: %d doesn't exist", execId)
	}
	return execution, nil
}

func (c *Client) Write(data []byte) (n int, err error) {
	err = c.Conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}

func (c *Client) CloseConnection() {
	c.Conn.Close()
}
