package client

import (
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type ClientRegistry struct {
	Connections map[uuid.UUID]Client
	mutex       sync.Mutex
}

func New() (registry *ClientRegistry) {
	return &ClientRegistry{
		Connections: make(map[uuid.UUID]Client),
	}
}

func (m *ClientRegistry) NewClient(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (client *Client, err error) {
	client, err = new(w, r, responseHeader)
	if err != nil {
		return nil, err
	}
	m.addClient(client)
	return client, nil
}

func (m *ClientRegistry) CloseClient(clientId uuid.UUID) {
	client := m.Connections[clientId]
	client.CloseConnection()

	m.removeClient(&client)
}

func (m *ClientRegistry) addClient(client *Client) {
	m.mutex.Lock()
	m.Connections[client.Id] = *client
	m.mutex.Unlock()
}

func (m *ClientRegistry) removeClient(client *Client) {
	m.mutex.Lock()
	delete(m.Connections, client.Id)
	m.mutex.Unlock()
}

func (m *ClientRegistry) CleanUp() {
	for _, client := range m.Connections {
		client.CloseConnection()
	}
}
