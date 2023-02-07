How to run code
- go run main.go ./ ./phone-number-test-data.xlsx

How to build application
- go build -o {{AppName}} main.go

How to run application
- MAX_RETRIES=2 RETRY_TIME=60 WAIT_TIME=60 CONTEXT=cos-all  ./{{AppName}} {{pathToSave}} {{pathToRead}}
- ex.  MAX_RETRIES=2 RETRY_TIME=60 WAIT_TIME=60 CONTEXT=cos-all ./go-test ./ ./phone-number-test-data.xlsx



