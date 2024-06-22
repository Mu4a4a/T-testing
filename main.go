package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("link.txt")
	if err != nil {
		fmt.Println("Файл пуст")
		return
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		input := reader.Text()
		fmt.Println(spammyMasker(input))
	}
}

func spammyMasker(a string) string {
	link := "http://"
	nlink := len(link)
	inputSlice := []byte(a)
	outputSlice := make([]byte, len(inputSlice))
	copy(outputSlice, inputSlice)

	for i := 0; i <= len(inputSlice)-nlink; i++ {
		if string(inputSlice[i:i+nlink]) == link {
			j := i + nlink
			for j < len(inputSlice) && (libray(inputSlice[j]) || inputSlice[j] == '_' || inputSlice[j] == '.' || inputSlice[j] == '~' || inputSlice[j] == '-') {
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
