package container

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/docker/docker/client"
)

type ContainerOrc struct {
	client          *client.Client
	containers      []Container
	idleQueue       chan *Container
	maxExecWaitTime time.Duration
}

func New(ctx context.Context, containerCount int, maxExecWaitTime time.Duration) (manager *ContainerOrc, err error) {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return &ContainerOrc{}, nil
	}

	manager = &ContainerOrc{
		client:          client,
		containers:      make([]Container, containerCount),
		idleQueue:       make(chan *Container, containerCount),
		maxExecWaitTime: maxExecWaitTime,
	}

	err = manager.initContainers(ctx, containerCount)
	if err != nil {
		return &ContainerOrc{}, err
	}

	return manager, nil
}

func (m *ContainerOrc) initContainers(ctx context.Context, containerCount int) (err error) {
	for i := 0; i < containerCount; i++ {
		container, err := new(ctx, m.client, fmt.Sprintf("executer-%d", i+1))
		if err != nil {
			return fmt.Errorf("failed to create container %d: %w", i+1, err)
		}
		m.containers[i] = *container
		m.idleQueue <- container
	}
	return nil
}

func (m *ContainerOrc) GetAvailableContainer(ctx context.Context) (*Container, error) {
	startTime := time.Now()

	for {
		select {
		case container := <-m.idleQueue:
			if container == nil {
				return nil, errors.New("no available idle containers")
			}
			container.setStatus(Active)
			return container, nil

		case <-time.After(m.maxExecWaitTime):
			if time.Since(startTime) > m.maxExecWaitTime {
				return nil, errors.New("timed out waiting for an available container")
			}

		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func (m *ContainerOrc) ReturnContainer(container *Container) {
	container.setStatus(Idle)
	m.idleQueue <- container
	//TODO: Add a cleanup of the container before returning it to the queue
}

func (m *ContainerOrc) Cleanup(ctx context.Context) {
	close(m.idleQueue)
	for _, container := range m.containers {
		container.remove(ctx)
	}
	m.client.Close()
}
