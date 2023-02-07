package main

import (
	"fmt"
	"os"

	"github.com/local/go-test/models"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Println(help())
		return
	}

	pathToSave := os.Args[1]
	pathToRead := os.Args[2]
	rows, _ := models.ReadFile(pathToRead)
	for _, row := range rows {
		models.ExportRowToFile(pathToSave, row)
	}
}

func help() string {
	return "Help \nUsage: `go run main.go /path/to/save  file/to/read`  or `./{{appName}} /path/to/save file/to/read`"
}
