package container

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/client"
)

type Manager struct {
	Client               *client.Client
	Containers           []Container
	IdleQueue            chan *Container
	MaxContainerWaitTime time.Duration
}

func NewManager(ctx context.Context, containerCount int, maxWaitTime time.Duration) (manager *Manager, err error) {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return &Manager{}, nil
	}

	manager = &Manager{
		Client:               client,
		Containers:           make([]Container, containerCount),
		IdleQueue:            make(chan *Container, containerCount),
		MaxContainerWaitTime: maxWaitTime,
	}

	err = manager.InitContainers(ctx, containerCount)
	if err != nil {
		return &Manager{}, nil
	}

	return manager, nil
}

func (m *Manager) InitContainers(ctx context.Context, containerCount int) (err error) {
	for i := 0; i < containerCount; i++ {
		container, err := newContainer(ctx, m.Client, fmt.Sprintf("Container-%d", i+1))
		if err != nil {
			return fmt.Errorf("failed to create container %d: %w", i+1, err)
		}
		m.Containers[i] = *container
		m.IdleQueue <- container
	}
	return nil
}

func (m *Manager) ExecuteInAvailableContainer(ctx context.Context, conn io.ReadWriteCloser, filePath string, resultFileName string, Cmd []string) error {
	startTime := time.Now()

	for {
		select {
		case container := <-m.IdleQueue:
			if container == nil {
				return errors.New("no available idle containers")
			}

			container.setStatus(Active)

			defer func(container *Container) {
				container.setStatus(Idle)
				m.IdleQueue <- container
				//TODO: Add a cleanup of the container before returning it to the queue
			}(container)

			fmt.Printf("Picked Container For Execution\nName: %s\nId: %s", container.ContainerName, container.ContainerId)

			return container.execute(ctx, conn, filePath, resultFileName, Cmd)

		case <-time.After(m.MaxContainerWaitTime):
			if time.Since(startTime) > m.MaxContainerWaitTime {
				return errors.New("timed out waiting for an available container")
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (m *Manager) Cleanup(ctx context.Context) {
	close(m.IdleQueue)
	for _, container := range m.Containers {
		container.remove(ctx)
	}
	m.Client.Close()
}
