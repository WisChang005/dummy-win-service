set GOARCH=amd64
go build -o .\bin\x64\dummy_service.exe dummy_service.go

set GOARCH=386
go build -o .\bin\win32\dummy_service.exe dummy_service.go
