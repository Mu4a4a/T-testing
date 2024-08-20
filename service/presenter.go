package service

import (
	"fmt"
	"os"
)

//go:generate go run github.com/vektra/mockery/v2@v2.44.1 --name=Presenter
type Presenter interface {
	Present([]byte) error
}

type FileWriter struct {
	FilePath string
}

func (f FileWriter) Present(in []byte) error {
	if f.FilePath == "" {
		return fmt.Errorf("the file name can`t be empty")
	}

	file, err := os.Create(f.FilePath)

	if err != nil {
		fmt.Println("ПРОИЗОШЛА ОШИБКА -", err)
		os.Exit(1)
	}

	defer file.Close()

	file.WriteString(spammyMasker(in))

	return err
}
