package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"github.com/tgunsch/httpod/docs"
	"github.com/tgunsch/httpod/handler"
)

func main() {

	docs.SwaggerInfo.Title = "httPod"
	docs.SwaggerInfo.Description = "A simple HTTP Request & Response Service, shamelessly stolen from httpbin."
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	server := echo.New()

	server.GET("/swagger/*", echoSwagger.WrapHandler)

	api := server.Group("/api")
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
