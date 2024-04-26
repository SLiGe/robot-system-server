set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux
go build -ldflags="-s -w" -o ./bin/server ./cmd/server