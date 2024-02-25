package container

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/docker/docker/client"
)

func TestManager__InitContainers(t *testing.T) {
	ctx := context.Background()
	manager, err := NewManager(ctx, 3, (2 * time.Minute))
	if err != nil {
		fmt.Println(err)
	}
	defer manager.Cleanup(ctx)

	log.Println(manager)
}

func TestContainer__Create(t *testing.T) {
	ctx := context.Background()
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
	}

	container, err := newContainer(ctx, client, "Test_Container")
	if err != nil {
		fmt.Println(err)
	}
	defer container.remove(ctx)
}

type testReadWriteCloser struct {
	input  *bytes.Buffer
	output *bytes.Buffer
}

func (trwc *testReadWriteCloser) Read(p []byte) (n int, err error) {
	return trwc.input.Read(p)
}

func (trwc *testReadWriteCloser) Write(p []byte) (n int, err error) {
	return trwc.output.Write(p)
}

func (trwc *testReadWriteCloser) Close() error {
	return nil
}

func TestContainer__Execute(t *testing.T) {
	ctx := context.Background()
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
	}

	container, err := newContainer(ctx, client, "Test_Container")
	if err != nil {
		fmt.Println(err)
	}

	// Create a testReadWriteCloser to capture output
	testRWClose := &testReadWriteCloser{
		input:  bytes.NewBuffer(nil), // Set your input data here
		output: bytes.NewBuffer(nil),
	}

	err = container.execute(context.TODO(), testRWClose, "./sample.txt", "./sample.go", []string{"go", "run", "sample.go"})
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(5 * time.Second)

	// Print the captured output to stdout
	fmt.Println("Captured Output:", testRWClose.output.String())

	defer container.remove(ctx)
}

func TestFlow_Manager_To_Container_To_User(t *testing.T) {
	ctx := context.Background()
	manager, err := NewManager(ctx, 3, (2 * time.Minute))
	if err != nil {
		fmt.Println(err)
	}
	defer manager.Cleanup(ctx)

	// Create a testReadWriteCloser to capture output
	testRWClose := &testReadWriteCloser{
		input:  bytes.NewBuffer(nil), // Set your input data here
		output: bytes.NewBuffer(nil),
	}

	go manager.ExecuteInAvailableContainer(ctx, testRWClose, "./sample.txt", "./sample.go", []string{"go", "run", "sample.go"})

	go manager.ExecuteInAvailableContainer(ctx, testRWClose, "./sample.txt", "./sample.go", []string{"go", "run", "sample.go"})

	go manager.ExecuteInAvailableContainer(ctx, testRWClose, "./sample.txt", "./sample.go", []string{"go", "run", "sample.go"})

	time.Sleep(10 * time.Second)

	// Print the captured output to stdout
	fmt.Println("Captured Output:", testRWClose.output.String())

}
