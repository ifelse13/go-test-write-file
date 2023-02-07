package models

import (
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type ExcelRow struct {
	PhoneNumber string `json:"phone_number"`
	CallerID    string `json:"caller_id"`
	Extension   string `json:"extension"`
}

func (r *ExcelRow) ParseToFile() File {

	return File{FileName: r.getFileName(), Lines: []string{r.channel(), r.callerID(), r.maxRetries(), r.retryTime(), r.waitTime(), r.context(), r.extension()}}
}
func (r *ExcelRow) getFileName() string {
	return fmt.Sprintf("file-name-%s-%s.txt", r.PhoneNumber, r.CallerID)
}
func (r *ExcelRow) channel() string {
	return fmt.Sprintf("Channel: PJSIP/%s/@%s\n", r.PhoneNumber, r.CallerID)
}
func (r *ExcelRow) callerID() string {
	return fmt.Sprintf("Callerid: %s\n", r.CallerID)
}
func (r *ExcelRow) maxRetries() string {
	return fmt.Sprintf("MaxRetries: %s\n", os.Getenv("MAX_RETRIES"))
}
func (r *ExcelRow) retryTime() string {
	return fmt.Sprintf("RetryTime: %s\n", os.Getenv("RETRY_TIME"))
}
func (r *ExcelRow) waitTime() string {
	return fmt.Sprintf("WaitTime: %s\n", os.Getenv("WAIT_TIME"))
}
func (r *ExcelRow) context() string {
	return fmt.Sprintf("Context: %s\n", os.Getenv("CONTEXT"))
}
func (r *ExcelRow) extension() string {
	return fmt.Sprintf("Extension: %s\n", r.Extension)
}
func ReadFile(path string) ([]ExcelRow, error) {
	var exelRows []ExcelRow
	fmt.Println(exelRows)
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	firstSheet := f.WorkBook.Sheets.Sheet[0].Name
	rows, err := f.Rows(firstSheet)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		columns, _ := rows.Columns()
		fmt.Println(columns)
		exelRows = append(exelRows, ExcelRow{PhoneNumber: columns[0], CallerID: columns[1], Extension: columns[2]})

	}
	return exelRows, nil
}
