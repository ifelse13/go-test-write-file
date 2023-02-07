package models

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type ExcelRow struct {
	PhoneNumber string `json:"phone_number"`
	CallerID    string `json:"caller_id"`
	Extension   string `json:"extension"`
	MaxRetries  string `json:"max_retries"`
	RetryTime   string `json:"retry_time"`
	WaitTime    string `json:"wait_time"`
	Context     string `json:"context"`
}

func (r *ExcelRow) ParseToFile(pathToSave string) File {

	return File{FileName: r.getFileName(), Lines: []string{r.channel(), r.callerID(), r.maxRetries(), r.retryTime(), r.waitTime(), r.context(), r.extension()}, PathToSave: pathToSave}
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
	return fmt.Sprintf("MaxRetries: %s\n", r.MaxRetries)
}
func (r *ExcelRow) retryTime() string {
	return fmt.Sprintf("RetryTime: %s\n", r.RetryTime)
}
func (r *ExcelRow) waitTime() string {
	return fmt.Sprintf("WaitTime: %s\n", r.WaitTime)
}
func (r *ExcelRow) context() string {
	return fmt.Sprintf("Context: %s\n", r.Context)
}
func (r *ExcelRow) extension() string {
	return fmt.Sprintf("Extension: %s\n", r.Extension)
}
func ReadFile(path string) ([]ExcelRow, error) {
	var exelRows []ExcelRow
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
		exelRows = append(exelRows, ExcelRow{PhoneNumber: columns[0], CallerID: columns[1], Extension: columns[2], MaxRetries: columns[3], RetryTime: columns[4], WaitTime: columns[5], Context: columns[6]})

	}
	return exelRows, nil
}
