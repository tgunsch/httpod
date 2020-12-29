package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Response struct {
	Args    map[string]string `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

func FromContext(c echo.Context) string {
	response := Response{
		Headers: getHeaders(c),
		Args:    getParams(c),
		Origin:  c.Request().Host,
		Url:     fmt.Sprintf("%s%s", c.Request().Host, c.Request().URL.Path),
	}
	prettyJSON, err := json.MarshalIndent(response, "", "   ")
	if err != nil {
		return fmt.Sprintf("Oops: %v", err.Error())
	}
	return string(prettyJSON)

}

func getHeaders(c echo.Context) map[string]string {
	headers := make(map[string]string)
	for k, v := range c.Request().Header {
		headers[k] = strings.Join(v, "; ")
	}
	return headers
}

func getParams(c echo.Context) map[string]string {
	parameters := make(map[string]string)
	for _, name := range c.ParamNames() {
		parameters[name] = c.Param(name)
	}
	return parameters
}

// getHost tries its best to return the request host.
func getHost(r *http.Request) string {
	if r.URL.IsAbs() {
		host := r.Host
		// Slice off any port information.
		if i := strings.Index(host, ":"); i != -1 {
			host = host[:i]
		}
		return host
	}
	return r.URL.Host
}
