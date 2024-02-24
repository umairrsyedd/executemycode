package container

import (
	"context"
	"fmt"
	"io"
	"os"
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

func NewContainer(ctx context.Context, client *client.Client, containerName string) (createdContainer *Container, err error) {
	createResponse, err := client.ContainerCreate(ctx, &container.Config{
		Image: "golang:latest",
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

func (c *Container) Execute(ctx context.Context, conn io.ReadWriteCloser, filePath string, resultFileName string, Cmd []string) (err error) {
	err = c.CopyToContainer(ctx, filePath, resultFileName)
	if err != nil {
		return err
	}

	execCreateResp, err := c.Client.ContainerExecCreate(ctx, c.ContainerId, types.ExecConfig{
		Cmd:          Cmd,
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
		io.Copy(execResp.Conn, conn) // Forward input to container
	}()

	go func() {
		io.Copy(conn, execResp.Reader) // Read output from container
	}()

	for {
		time.Sleep(2 * time.Second)
		execInspect, err := c.Client.ContainerExecInspect(ctx, execCreateResp.ID)
		if err != nil {
			return err
		}
		if !execInspect.Running {
			break
		}
	}

	return nil
}

func (c *Container) CopyToContainer(ctx context.Context, filePath string, resultFileName string) (err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	err = c.Client.CopyToContainer(ctx, c.ContainerId,
		"./go/",
		GetTarFile(fileContent, resultFileName),
		types.CopyToContainerOptions{})
	if err != nil {
		return fmt.Errorf("error copying file to container: %v", err)
	}

	return nil
}

func (c *Container) Remove(ctx context.Context) error {
	return c.Client.ContainerRemove(ctx, c.ContainerId, container.RemoveOptions{
		Force: true,
	})
}

func (c *Container) SetStatus(status ContainerStatus) {
	c.Status = status
}

func (c *Container) IsActive() bool {
	return c.Status == Active
}
