package models

import (
	"log"
	"os"
)

type File struct {
	Lines      []string `json:"lines"`
	FileName   string   `json:"file_name"`
	PathToSave string   `json:"path_to_save"`
}

func ExportRowToFile(pathToSave string, row ExcelRow) error {
	file := row.ParseToFile(pathToSave)
	return exportToFile(file)

}

func exportToFile(file File) error {
	f, err := os.Create(file.FileName)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer f.Close()

	for _, line := range file.Lines {
		_, err := f.WriteString(line)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}
