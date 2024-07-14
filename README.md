# Initialize module file
go mod init module/go

# Downloading module

## Gin module
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
<!-- Package debug -->
go install github.com/go-delve/delve/cmd/dlv@latest

# Running application
go run main.go

# Reference
- Go for dev:
    https://go.dev/tour/welcome/1

