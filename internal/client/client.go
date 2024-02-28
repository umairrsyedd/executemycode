package client

import (
	"executemycode/internal/program"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id       uuid.UUID
	Conn     *websocket.Conn
	Programs map[int]*program.Program
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func new(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (client *Client, err error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		Id:       uuid.New(),
		Conn:     conn,
		Programs: make(map[int]*program.Program),
	}, nil
}

func (c *Client) AddProgram(programId int, program *program.Program) {
	c.Programs[programId] = program
}

func (c *Client) WriteMessage(data []byte) (err error) {
	err = c.Conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CloseConnection() {
	c.Conn.Close()
}
