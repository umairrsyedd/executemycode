package container

import (
	"archive/tar"
	"bytes"
	"io"
	"log"
)

func GetTarFile(file []byte, resultFileName string) io.Reader {
	// file, err := os.ReadFile(filePath)
	// if err != nil {
	// 	panic(fmt.Errorf("error reading file: %v", err))
	// }

	// 5. Store the Opened file inside the app directory using tar in the container as sample.go
	var tarBuffer bytes.Buffer
	tarWriter := tar.NewWriter(&tarBuffer)
	err := tarWriter.WriteHeader(&tar.Header{
		Name: resultFileName,
		Mode: 0755,
		Size: int64(len(file)),
	})
	if err != nil {
		log.Fatalf("Error writing TAR header: %v", err)
	}

	_, err = tarWriter.Write(file)
	if err != nil {
		log.Fatalf("Error writing file content to TAR: %v", err)
	}

	err = tarWriter.Close()
	if err != nil {
		log.Fatalf("Error closing TAR writer: %v", err)
	}

	return &tarBuffer
}
