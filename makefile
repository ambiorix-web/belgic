default: darwin
	GOOS=linux GOARCH=amd64 go build -o bin/belgic *.go

darwin: windows
	GOOS=darwin GOARCH=amd64 go build -o bin/belgic-mac *.go

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/belgic.exe *.go
