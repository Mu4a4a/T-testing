package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestPresenter(t *testing.T) {
	testTable := []struct {
		nameTest   string
		shouldFail bool
		nameFile   string
		data       []byte
	}{
		{
			nameTest:   "Valid test",
			shouldFail: false,
			nameFile:   "test.txt",
			data:       []byte("11111"),
		},
		{
			nameTest:   "Empty name file",
			shouldFail: true,
			nameFile:   "",
			data:       []byte("some name"),
		},
		{
			nameTest:   "Empty data file",
			shouldFail: false,
			nameFile:   "test.txt",
			data:       []byte(""),
		},
	}
	for _, tt := range testTable {
		t.Run(tt.nameTest, func(t *testing.T) {

			defer os.Remove(tt.nameFile)

			fileWriter := FileWriter{FilePath: tt.nameFile}

			err := fileWriter.Present(tt.data)
			if tt.shouldFail {
				require.Error(t, err, "Ожидалась ошибка, но её не произошло")
				return
			} else {
				require.NoError(t, err, "Ошибка при записи файла: %s", fileWriter.FilePath)
			}

			assert.FileExists(t, fileWriter.FilePath)

			writtenData, err := os.ReadFile(fileWriter.FilePath)
			require.NoError(t, err, "Error reading file: %s", fileWriter.FilePath)

			assert.Equal(t, tt.data, writtenData)
		})
	}
}
