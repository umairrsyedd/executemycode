package container

import (
	"bufio"
	"context"
	"executemycode/internal/executer"
	"fmt"
	"log"
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

func new(ctx context.Context, client *client.Client, containerName string) (createdContainer *Container, err error) {
	containerImageName := os.Getenv("CONTAINER_IMAGE_NAME")
	createResponse, err := client.ContainerCreate(ctx, &container.Config{
		Image: containerImageName,
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
	fileName := "program"
	fileNameWithExt := fmt.Sprintf("%s%s", fileName, execution.ExecutionInfo.LangExecuter.GetFileExt())
	err = c.copyToContainer(ctx, []byte(execution.ExecutionInfo.SourceCode), fileNameWithExt)
	if err != nil {
		return err
	}

	if execution.ExecutionInfo.LangExecuter.IsRunCompiled() {
		err := c.compileCode(ctx, execution.ExecutionInfo.LangExecuter.GetCompileCmd(fileName))
		if err != nil {
			return err
		}

		time.Sleep(2 * time.Second)

		// TODO: Fix this Stuff Later
		// err = c.waitForCompiledFileToBeAvailable(ctx, fileName)
		// if err != nil {
		// 	return err
		// }

	}

	execCreateResp, err := c.Client.ContainerExecCreate(ctx, c.ContainerId, types.ExecConfig{
		Cmd:          execution.ExecutionInfo.LangExecuter.GetRunCmd(fileName),
		Tty:          true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		return err
	}

	execResp, err := c.Client.ContainerExecAttach(ctx, execCreateResp.ID, types.ExecStartCheck{
		Tty:    true,
		Detach: false,
	})
	if err != nil {
		return err
	}
	defer execResp.Close()

	// Input Go Routine
	go func() {
		for {
			input, ok := <-execution.InputChan
			if !ok {
				return
			}
			fmt.Fprintln(execResp.Conn, input)
		}
	}()

	// Output Go Routine
	go func() {
		scanner := bufio.NewScanner(execResp.Conn)
		for scanner.Scan() {
			output := scanner.Text()
			execution.OutputChan <- output
		}
	}()

	for {
		select {
		case stopRequest := <-execution.StopChan:
			if stopRequest {
				ctx.Done()
				return nil
			}
		default:
			execInspect, err := c.Client.ContainerExecInspect(ctx, execCreateResp.ID)
			if err != nil {
				return err
			}
			if !execInspect.Running {
				if !execution.IsDone {
					execution.ExitCode <- execInspect.ExitCode
				}
				return nil
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func (c *Container) compileCode(ctx context.Context, compileCmd []string) error {
	compileResp, err := c.Client.ContainerExecCreate(ctx, c.ContainerId, types.ExecConfig{
		Cmd:          compileCmd,
		Tty:          false,
		AttachStdin:  false,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		return err
	}

	_, err = c.Client.ContainerExecAttach(ctx, compileResp.ID, types.ExecStartCheck{
		Tty:    false,
		Detach: false,
	})
	if err != nil {
		return err
	}

	// Wait for compilation to finish
	err = c.Client.ContainerExecStart(ctx, compileResp.ID, types.ExecStartCheck{})
	if err != nil {
		return err
	}

	// Check compilation result
	compileInspect, err := c.Client.ContainerExecInspect(ctx, compileResp.ID)
	if err != nil {
		return err
	}
	if compileInspect.ExitCode != 0 {
		// Compilation failed
		return fmt.Errorf("compilation failed with exit code: %d", compileInspect.ExitCode)
	}

	return nil
}

func (c *Container) waitForCompiledFileToBeAvailable(ctx context.Context, fileName string) error {
	for {
		_, err := c.Client.ContainerInspect(ctx, c.ContainerId)
		if err != nil {
			return err
		}

		// Check if the file exists in the container
		fileExists, err := c.checkFileExists(ctx, fileName)
		if err != nil {
			return err
		}

		if fileExists {
			break
		}

		// Wait before checking again
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func (c *Container) checkFileExists(ctx context.Context, fileName string) (bool, error) {
	resp, err := c.Client.ContainerStatPath(ctx, c.ContainerId, "app/program")
	if err != nil {
		return false, err
	}

	return resp.Name == fileName, nil
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

func (c *Container) cleanup(ctx context.Context) error {
	execCreateResp, err := c.Client.ContainerExecCreate(ctx, c.ContainerId, types.ExecConfig{
		Cmd:          []string{"sh", "-c", "rm -rf /app/*"},
		Tty:          false,
		AttachStdin:  false,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		return fmt.Errorf("failed to execute cleanup command: %v", err)
	}

	execResp, err := c.Client.ContainerExecAttach(ctx, execCreateResp.ID, types.ExecStartCheck{
		Tty:    false,
		Detach: false,
	})
	if err != nil {
		return fmt.Errorf("failed to attach to exec instance: %v", err)
	}
	defer execResp.Close()

	// Wait for cleanup command to finish
	err = c.Client.ContainerExecStart(ctx, execCreateResp.ID, types.ExecStartCheck{})
	if err != nil {
		return fmt.Errorf("failed to start cleanup command: %v", err)
	}

	// Check cleanup command result
	execInspect, err := c.Client.ContainerExecInspect(ctx, execCreateResp.ID)
	if err != nil {
		return fmt.Errorf("failed to inspect cleanup command: %v", err)
	}
	if execInspect.ExitCode != 0 {
		return fmt.Errorf("cleanup command failed with exit code: %d", execInspect.ExitCode)
	}

	log.Printf("Cleaned Up %s App Directory", c.ContainerName)

	return nil
}
