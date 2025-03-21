build-linux:
	env GOOS=linux GOARCH=amd64 go build -v -o simple-healthcheck *.go