package main

import (
	"fmt"
	"learnGO/service"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Не указан путь к файлу.")
		return
	}
	inputFilePath := args[0]

	var outputFilePath string

	if len(args) >= 2 {
		outputFilePath = args[1]
	} else {
		outputFilePath = "output.txt"
	}

	prod := service.FileReader{FilePath: inputFilePath}
	pres := service.FileWriter{FilePath: outputFilePath}
	i := service.NewService(prod, pres)

	i.Run()
}
