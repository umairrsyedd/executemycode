package container

import (
	"archive/tar"
	"bytes"
	"io"
	"log"
)

func GetTarFile(file []byte, resultFileName string) io.Reader {
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
