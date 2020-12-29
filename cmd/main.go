package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"github.com/tgunsch/httpod/internal/docs"
	"github.com/tgunsch/httpod/internal/handler"
	"github.com/tgunsch/httpod/internal/util"
	"os"
	"strings"
)

func main() {
	const (
		SWAGGER_PATH = "/swagger"
		API_PATH     = "/api"
	)
	server := echo.New()
	basePath := os.Getenv("BASE_PATH")
	if basePath != "" {
		basePath = "/" + basePath
	}
	endpoints := server.Group(basePath)

	endpoints.GET(SWAGGER_PATH+"/*", echoSwagger.WrapHandler, swaggerMiddleware(basePath+API_PATH))

	api := endpoints.Group(API_PATH)
	api.DELETE("/delete", handler.HttpDeleteHandler)
	api.GET("/get", handler.HttpGetHandler)
	api.PATCH("/patch", handler.HttpPatchHandler)
	api.POST("/post", handler.HttpPostHandler)
	api.PUT("/put", handler.HttpPutHandler)

	api.DELETE("/status/:code", handler.StatusDeleteHandler)
	api.GET("/status/:code", handler.StatusGetHandler)
	api.PATCH("/status/:code", handler.StatusPatchHandler)
	api.POST("/status/:code", handler.StatusPostHandler)
	api.PUT("/status/:code", handler.StatusPutHandler)

	server.Logger.Fatal(server.Start(":8080"))
}

func swaggerMiddleware(path string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasSuffix(c.Request().URL.Path, "index.html") {

				docs.SwaggerInfo.Title = "httPod"
				docs.SwaggerInfo.Description = "A simple HTTP Request & Response Service, shamelessly stolen from httpbin."
				docs.SwaggerInfo.Version = "0.0.1"

				scheme, host := util.GetSchemeHost(c.Request())
				docs.SwaggerInfo.Host = host
				docs.SwaggerInfo.Schemes = []string{scheme}

				docs.SwaggerInfo.BasePath = util.GetPrefix(path, c.Request())
			}
			return next(c)
		}
	}
}
