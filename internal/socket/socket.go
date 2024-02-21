package socket

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id     uuid.UUID
	Conn   *websocket.Conn
	Closed chan bool
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewClient(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (client *Client, err error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		Id:     uuid.New(),
		Conn:   conn,
		Closed: make(chan bool),
	}, nil
}

func (c *Client) CloseConnection() {
	c.Conn.Close()
}

func (c *Client) WriteMessage(data []byte) (err error) {
	err = c.Conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReadMessages(receivingChannel chan any) {
	for {
		var message Message
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
				c.Closed <- true
			} else {
				log.Printf("error reading json message %v", err)
			}
		}
	}
}
