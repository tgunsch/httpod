package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"github.com/tgunsch/httpod/internal/cookies"
	"github.com/tgunsch/httpod/internal/docs"
	"github.com/tgunsch/httpod/internal/http"
	"github.com/tgunsch/httpod/internal/jwt"
	"github.com/tgunsch/httpod/internal/proxy"
	"github.com/tgunsch/httpod/internal/status"
	"github.com/tgunsch/httpod/internal/util"
	"html"
	"os"
	"strconv"
	"strings"
)

// @title httPod
// @version 1.0.0
// @description A simple HTTP Request & HTTPResponse Service, shamelessly stolen from httpbin.org.
// @tag.name HTTP Methods
// @tag.description Testing different HTTP methods
// @tag.name Status codes
// @tag.description Generates responses with given status code
// @tag.name Cookies
// @tag.description Creates, reads and deletes Cookies

func main() {
	const (
		SwaggerPath = "/swagger"
		ApiPath     = "/api"
		BasePathEnv = "BASE_PATH"
		PORT        = "PORT"
	)
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true
	port := os.Getenv(PORT)
	if port == "" {
		port = "8080"
	}
	basePath := os.Getenv(BasePathEnv)
	if basePath != "" {
		basePath = "/" + basePath
	}
	endpoints := server.Group(basePath)

	// swagger ui will be available on /basePath/swagger/index.html
	// api will be available on /basePath/api
	// swagger info will use X-Forwarded headers if available;
	// e.g.: X-Forwarded-Host=my.domain.com X-Forwarded-Prefix=myPrefix swagger ui show api on url http://my.domain.com/myPrefix/basePath/api
	apiPath := basePath + ApiPath
	endpoints.GET(SwaggerPath+"/*", echoSwagger.WrapHandler, swaggerMiddleware(apiPath))

	api := endpoints.Group(ApiPath)
	api.GET("/get", http.GetHandler)
	api.DELETE("/delete", http.DeleteHandler)
	api.PATCH("/patch", http.PatchHandler)
	api.POST("/post", http.PostHandler)
	api.PUT("/put", http.PutHandler)

	api.DELETE("/status/:code", status.DeleteHandler)
	api.GET("/status/:code", status.GetHandler)
	api.PATCH("/status/:code", status.PatchHandler)
	api.POST("/status/:code", status.PostHandler)
	api.PUT("/status/:code", status.PutHandler)

	api.GET("/cookies", cookies.GetHandler)
	api.POST("/cookies/:cookieName", cookies.PostHandler)
	api.DELETE("/cookies/:cookieName", cookies.DeleteHandler)

	api.GET("/jwt", jwt.GetHandler)

	api.GET("/proxy", proxy.GetHandler)

	println(banner("http://localhost:" + port + basePath + SwaggerPath + "/index.html"))
	server.Logger.Fatal(server.Start(":" + port))
}

func banner(localURL string) string {
	const BANNER = `/ˌeɪtʃ tiː tiː ˈpɒd/ %s trapping on %s`
	honeyPod := html.UnescapeString("&#" + strconv.Itoa(0x1f36f) + ";")
	return fmt.Sprintf(BANNER, honeyPod, localURL)
}

func swaggerMiddleware(path string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasSuffix(c.Request().URL.Path, "index.html") {

				scheme, host := util.GetSchemeHost(c.Request())
				docs.SwaggerInfo.Host = host
				switch scheme {
				case "https", "HTTPS":
					docs.SwaggerInfo.Schemes = []string{"https", "http"}
				default:
					docs.SwaggerInfo.Schemes = []string{"http", "https"}
				}

				docs.SwaggerInfo.BasePath = util.GetPath(path, c.Request())
			}
			return next(c)
		}
	}
}
