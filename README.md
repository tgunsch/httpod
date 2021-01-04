# httpod

![Build](https://github.com/tgunsch/httpod/workflows/Go/badge.svg)
![Docker Image](https://github.com/tgunsch/httpod/workflows/Docker%20Image%20CI/badge.svg)
A simple HTTP Request & Response Service written in go, shamelessly stolen from https://httpbin.org.

## Devel

```shell
# Download swag once:
go get -u github.com/swaggo/swag/cmd/swag
# Create swagger info
swag init -g cmd/main.go -o internal/docs 

# Run
go run cmd/main.go

# Open browser on http://localhost:8080/swagger/index.html
```

## Build docker image

```shell
docker build -t thgunsch/httpod
```
