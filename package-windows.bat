set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=windows
go env -w GOOS=windows
go build -ldflags="-s -w" -o ./bin/server.exe ./cmd/server