package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestProducer(t *testing.T) {
	testTable := []struct {
		nameTest     string
		shouldFail   bool
		fileName     string
		data         string
		expectedData string
	}{
		{
			nameTest:     "Empty name",
			shouldFail:   true,
			fileName:     "",
			data:         "testdata",
			expectedData: "testdata",
		},
		{
			nameTest:     "Empty file",
			shouldFail:   false,
			fileName:     "testdata.txt",
			data:         "",
			expectedData: "",
		},
		{
			nameTest:     "Valid file",
			shouldFail:   false,
			fileName:     "testdata.txt",
			data:         "testdata",
			expectedData: "testdata",
		},
	}
	for _, tc := range testTable {
		t.Run(tc.nameTest, func(t *testing.T) {
			tempFile, err := os.Create(tc.fileName)
			if err != nil {
				if tc.shouldFail {
					fmt.Printf("Не удалось создать файл: %v", err)
					return
				}
				t.Fatalf("Не удалось создать файл: %v", err)
			}
			defer os.Remove(tempFile.Name())

			if _, err := tempFile.Write([]byte(tc.data)); err != nil {
				t.Fatalf("Не удалось записать в файл: %v", err)
			}

			if err := tempFile.Close(); err != nil {
				t.Fatalf("Не удалось закрыть файл: %v", err)
			}

			fileReader := FileReader{FilePath: tempFile.Name()}

			dataProd, err := fileReader.Produce()
			if err != nil {
				t.Fatalf("Ожидалось, что ошибки не будет, но получили: %v", err)
			}
			assert.Equal(t, []byte(tc.expectedData), dataProd, tc.nameTest)
		})
	}
}
