package status

import (
	"github.com/labstack/echo/v4"
	http2 "github.com/tgunsch/httpod/internal/http"
	"net/http"
	"strconv"
)

// @Summary Do a GET request.
// @Tags Status codes
// @Description Requests using GET should only retrieve data.
// @Accept  json
// @Produce  json
// @Param code path string false "return status code"
// @response 100 {string} string	"Informational responses"
// @response 200 {string} string	"Success"
// @response 300 {string} string	"Redirection"
// @response 400 {string} string	"Client Errors"
// @response 418 {string} string	"I'm a teapot"
// @response 500 {string} string	"Server Errors"
// @Router /status/{code} [get]
func GetHandler(context echo.Context) error {
	return statusFromRequest(context)
}

// @Summary Do a DELETE request.
// @Tags Status codes
// @Description The DELETE method deletes the specified resource.
// @Accept  json
// @Produce  json
// @Param code path string false "return status code"
// @response 100 {string} string	"Informational responses"
// @response 200 {string} string	"Success"
// @response 300 {string} string	"Redirection"
// @response 400 {string} string	"Client Errors"
// @response 418 {string} string	"I'm a teapot"
// @response 500 {string} string	"Server Errors"
// @Router /status/{code} [delete]
func DeleteHandler(context echo.Context) error {
	return statusFromRequest(context)
}

// @Summary Do a POST request.
// @Tags Status codes
// @Description The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server.
// @Accept  json
// @Produce  json
// @Param code path string false "return status code"
// @response 100 {string} string	"Informational responses"
// @response 200 {string} string	"Success"
// @response 300 {string} string	"Redirection"
// @response 400 {string} string	"Client Errors"
// @response 418 {string} string	"I'm a teapot"
// @response 500 {string} string	"Server Errors"
// @Router /status/{code} [post]
func PostHandler(context echo.Context) error {
	return statusFromRequest(context)
}

// @Summary Do PUT request.
// @Tags Status codes
// @Description The PUT method replaces all current representations of the target resource with the request payload.
// @Accept  json
// @Produce  json
// @Param code path string false "return status code"
// @response 100 {string} string	"Informational responses"
// @response 200 {string} string	"Success"
// @response 300 {string} string	"Redirection"
// @response 400 {string} string	"Client Errors"
// @response 418 {string} string	"I'm a teapot"
// @response 500 {string} string	"Server Errors"
// @Router /status/{code} [put]
func PutHandler(context echo.Context) error {
	return statusFromRequest(context)
}

// @Summary Do a PATCH request.
// @Tags Status codes
// @Description The PATCH method is used to apply partial modifications to a resource.
// @Accept  json
// @Produce  json
// @Param code path string false "return status code"
// @response 100 {string} string	"Informational responses"
// @response 200 {string} string	"Success"
// @response 300 {string} string	"Redirection"
// @response 400 {string} string	"Client Errors"
// @response 418 {string} string	"I'm a teapot"
// @response 500 {string} string	"Server Errors"
// @Router /status/{code} [patch]
func PatchHandler(context echo.Context) error {
	return statusFromRequest(context)
}

func statusFromRequest(context echo.Context) error {
	param := context.Param("code")
	code, err := strconv.Atoi(param)
	if err != nil {
		return context.String(http.StatusBadRequest, "Illegal code "+param)
	}
	return context.String(code, http2.ResponseFromContext(context))
}
