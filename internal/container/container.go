package container

import (
	"bufio"
	"context"
	"executemycode/internal/executer"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type ContainerStatus string

const (
	Idle   ContainerStatus = "Idle"
	Active ContainerStatus = "Acive"
)

type Container struct {
	Client        *client.Client
	ContainerId   string
	ContainerName string
	Status        ContainerStatus
}

func new(ctx context.Context, client *client.Client, containerName string) (createdContainer *Container, err error) {
	createResponse, err := client.ContainerCreate(ctx, &container.Config{
		Image: "executer-image",
		Cmd:   []string{"sleep", "infinity"},
		Tty:   false,
	}, &container.HostConfig{}, &network.NetworkingConfig{}, &v1.Platform{}, containerName)
	if err != nil {
		return nil, err
	}

	err = client.ContainerStart(ctx, createResponse.ID, container.StartOptions{})
	if err != nil {
		client.ContainerRemove(ctx, createResponse.ID, container.RemoveOptions{
			Force: true,
		})
		return nil, err
	}

	return &Container{
		Client:        client,
		ContainerId:   createResponse.ID,
		ContainerName: containerName,
		Status:        Idle,
	}, nil
}

func (c *Container) execute(ctx context.Context, execution *executer.Execution) (err error) {
	err = c.copyToContainer(ctx, []byte(execution.ExecutionInfo.SourceCode), "sample.go") // TODO: Remove Hardcoding
	if err != nil {
		return err
	}

	execCreateResp, err := c.Client.ContainerExecCreate(ctx, c.ContainerId, types.ExecConfig{
		Cmd:          []string{"go", "run", "sample.go"}, // TODO: Remove Hardcoding
		Tty:          true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		return err
	}

	execResp, err := c.Client.ContainerExecAttach(ctx, execCreateResp.ID, types.ExecStartCheck{
		Tty: true,
	})
	if err != nil {
		return err
	}
	defer execResp.Close()

	go func() {
		for {
			input, ok := <-execution.InputChan
			if !ok {
				return
			}
			fmt.Printf("I'm writing this input to the container: %v\n", input)
			_, err := execResp.Conn.Write([]byte(input))
			if err != nil {
				fmt.Printf("error writing to container: %v", err)
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(execResp.Reader)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			output := scanner.Text()
			execution.OutputChan <- output
		}
	}()

	for {
		execInspect, err := c.Client.ContainerExecInspect(ctx, execCreateResp.ID)
		if err != nil {
			return err
		}
		if !execInspect.Running {
			execution.DoneChan <- true
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (c *Container) copyToContainer(ctx context.Context, content []byte, resultFileName string) (err error) {
	err = c.Client.CopyToContainer(ctx, c.ContainerId,
		"./app/",
		getTarFile(content, resultFileName),
		types.CopyToContainerOptions{})
	if err != nil {
		return fmt.Errorf("error copying file to container: %v", err)
	}

	return nil
}

func (c *Container) remove(ctx context.Context) error {
	return c.Client.ContainerRemove(ctx, c.ContainerId, container.RemoveOptions{
		Force: true,
	})
}

func (c *Container) setStatus(status ContainerStatus) {
	c.Status = status
}
