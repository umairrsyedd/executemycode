package client

import (
	"executemycode/internal/executer"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id        uuid.UUID
	Conn      *websocket.Conn
	Execution *executer.Execution
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func new(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (client *Client, err error) {
	conn, err := upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		Id:        uuid.New(),
		Conn:      conn,
		Execution: nil,
	}, nil
}

func (c *Client) HasPrevExecution() bool {
	return c.Execution != nil
}

func (c *Client) SetExecution(execution *executer.Execution) {
	c.Execution = execution
}

func (c *Client) GetExecution() (*executer.Execution, error) {
	if c.Execution == nil {
		return nil, fmt.Errorf("no Code is currently being executed")
	}
	return c.Execution, nil
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
