# httpod

![Build](https://github.com/tgunsch/httpod/workflows/Go/badge.svg)
![Docker Image](https://github.com/tgunsch/httpod/workflows/Docker%20Image%20CI/badge.svg)

/ˌeɪtʃ tiː tiː ˈpɒd/: a simple HTTP Request & Response Service written in go, shamelessly stolen from [httpbin](https://httpbin.org).

![swagger-ui](docs/swagger-ui.png)

### Additional features
Like the famous `httpbin` this service provides several API's for testing HTTP requests and a 
corresponding swagger GUI. It adds some features helpful for testing in kubernetes or istio environments:
* With environment variable `BASE_PATH`, a "*context root*" can be configured. This is helpful if routing is done via path prefix.
* The service handles `X-Forwarded-Host` and `X-Forwarded-Prefix` headers. This is helpful, if ingress controller or virtual services rewrite the destination path.

### Currently implemented API's:
* **HTTP Methods**: Testing DELETE, GET, PATCH, POST and  PUT requests. TODO: add url parameters
* **Status codes**: Generates responses with given status code
* **Cookies**: Creates, reads and deletes cookies
* **JWT**: Show and validated JWT Bearer Token

## Devel

```shell
# Download swag once:
go install github.com/swaggo/swag/cmd/swag@latest
# Create swagger info
swag init -g cmd/main.go -o internal/docs

# Run
go run cmd/main.go

# Open browser on http://localhost:8080/swagger/index.html
```

## Run docker image

Two docker images are available:
* latest: based on `ubuntu`. This is mainly for debugging purpose (i.e. running curl in running container)
* latest-slim: scratch image only containing the binary 

```shell
docker pull ghcr.io/tgunsch/httpod:latest-slim

# Access http://localhost:8080/swagger/index.html
docker run --rm -p 8080:8080  ghcr.io/tgunsch/httpod:latest-slim

# Access http://localhost:8080/httpod/swagger/index.html
docker run --rm -p 8080:8080  -e BASE_PATH=httpod ghcr.io/tgunsch/httpod:latest-slim
```
