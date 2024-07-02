package service

import (
	"fmt"
	"os"
)

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{prod: prod, pres: pres}
}

type Producer interface {
	produce() ([]byte, error)
}

type Presenter interface {
	present([]byte) error
}

type Service struct {
	prod Producer
	pres Presenter
}

type FileReader struct {
	FilePath string
}

func (f FileReader) produce() ([]byte, error) {
	fileData, err := os.ReadFile(f.FilePath)

	if err != nil {
		fmt.Println("НЕ УДАЛОСЬ ОТКРЫТЬ, ОШИБКА-", err)
		os.Exit(1)
	}
	return fileData, err
}

type FileWriter struct {
	FilePath string
}

func (f FileWriter) present(in []byte) error {

	file, err := os.Create(f.FilePath)

	if err != nil {
		fmt.Println("ПРОИЗОШЛА ОШИБКА -", err)
		os.Exit(1)
	}

	defer file.Close()

	file.WriteString(spammyMasker(in))
	return err
}

func (s Service) Run() {

	in, _ := s.prod.produce()

	s.pres.present(in)
}

func spammyMasker(input []byte) string {
	link := "http://"
	nlink := len(link)
	outputSlice := make([]byte, len(input))
	copy(outputSlice, input)

	for i := 0; i <= len(input)-nlink; i++ {
		if string(input[i:i+nlink]) == link {
			j := i + nlink
			for j < len(input) && (libray(input[j]) || input[j] == '_' || input[j] == '.' || input[j] == '~' || input[j] == '-') {
				outputSlice[j] = '*'
				j++
			}
			i = j - 1
		}
	}
	return string(outputSlice)
}

func libray(a byte) bool {
	return (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}
