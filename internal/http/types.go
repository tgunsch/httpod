package http

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"strings"
)

type Response struct {
	Host       string            `json:"host"`
	RemoteAddr string            `json:"remote-address"`
	Url        string            `json:"url"`
	Args       map[string]string `json:"args"`
	Headers    map[string]string `json:"headers"`
}

func ResponseFromContext(c echo.Context) string {
	response := Response{
		Headers:    getHeaders(c),
		Args:       getParams(c),
		Host:       c.Request().Host,
		RemoteAddr: c.Request().RemoteAddr,
		Url:        fmt.Sprintf("%s%s", c.Request().Host, c.Request().URL.Path),
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
