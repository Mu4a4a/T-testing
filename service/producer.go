package service

import (
	"fmt"
	"os"
)

//go:generate go run github.com/vektra/mockery/v2@v2.44.1 --name=Producer
type Producer interface {
	Produce() ([]byte, error)
}

type FileReader struct {
	FilePath string
}

func (f FileReader) Produce() ([]byte, error) {
	fileData, err := os.ReadFile(f.FilePath)

	if err != nil {
		fmt.Println("НЕ УДАЛОСЬ ОТКРЫТЬ, ОШИБКА-", err)
		os.Exit(1)
	}
	return fileData, err
}
