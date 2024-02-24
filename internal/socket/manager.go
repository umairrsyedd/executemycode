package socket

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

var ConnectionManager Manager

type Manager struct {
	Connections map[uuid.UUID]Client
	mutex       sync.Mutex
}

func init() {
	ConnectionManager = Manager{
		Connections: make(map[uuid.UUID]Client),
	}
}

func (m *Manager) ConnectClient(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (client *Client, err error) {
	client, err = newClient(w, r, responseHeader)
	if err != nil {
		return nil, err
	}
	m.addClient(client)
	return client, nil
}

func (m *Manager) DisconnectClient(clientId uuid.UUID) {
	client := m.Connections[clientId]
	client.CloseConnection()

	m.removeClient(&client)
}

func (m *Manager) addClient(client *Client) {
	m.mutex.Lock()
	m.Connections[client.Id] = *client
	m.mutex.Unlock()
}

func (m *Manager) removeClient(client *Client) {
	m.mutex.Lock()
	delete(m.Connections, client.Id)
	m.mutex.Unlock()
}

func (m *Manager) LogActiveConnections() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("\nActive Connections : %d", len(m.Connections))
	}
}

func (m *Manager) CleanUpConnections() {
	for _, client := range m.Connections {
		client.CloseConnection()
	}
}
