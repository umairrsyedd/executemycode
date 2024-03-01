package container

import (
	"context"
	"executemycode/internal/executer"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/client"
)

type ContainerOrc struct {
	client          *client.Client
	containers      []Container
	idleQueue       chan *Container
	maxExecWaitTime time.Duration
	execTimeout     time.Duration
}

func New(ctx context.Context, containerCount int, maxExecWaitTime time.Duration, execTimeout time.Duration) (manager *ContainerOrc, err error) {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return &ContainerOrc{}, nil
	}

	manager = &ContainerOrc{
		client:          client,
		containers:      make([]Container, containerCount),
		idleQueue:       make(chan *Container, containerCount),
		maxExecWaitTime: maxExecWaitTime,
		execTimeout:     execTimeout,
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

func (m *ContainerOrc) ConnectAndExecute(execution *executer.Execution) error {
	startTime := time.Now()

	executionCtx, cancel := context.WithTimeout(context.Background(), m.execTimeout)
	defer cancel()
	for {
		select {
		case container := <-m.idleQueue:
			if container == nil {
				return fmt.Errorf("no available idle containers")
			}
			container.setStatus(Active)

			defer func(container *Container) {
				container.setStatus(Idle)
				m.idleQueue <- container
			}(container)

			err := container.execute(executionCtx, execution)
			if err != nil {
				log.Printf("error from container execute: %s", err)
				return err
			}
			return nil

		case <-time.After(m.maxExecWaitTime):
			if time.Since(startTime) > m.maxExecWaitTime {
				log.Printf("timed out waiting for an available container")
				return fmt.Errorf("timed out waiting for an available container")
			}

		case <-executionCtx.Done():
			log.Printf("execution context is done")
			return executionCtx.Err()
		}
	}
}

func (m *ContainerOrc) Cleanup(ctx context.Context) {
	close(m.idleQueue)
	for _, container := range m.containers {
		container.remove(ctx)
	}
	m.client.Close()
}
