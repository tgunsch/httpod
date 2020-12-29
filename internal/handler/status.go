package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// @Summary The request's query parameters.
// @Tags Status codes
// @Description get
// @Accept  json
// @Produce  json
// @Param code path string false "return status code"
// @response 100 {string} string	"Informational responses"
// @response 200 {string} string	"Success"
// @response 300 {string} string	"Redirection"
// @response 400 {string} string	"Client Errors"
// @response 500 {string} string	"Server Errors"
// @Router /status/{code} [get]
func StatusGetHandler(context echo.Context) error {
	return statusFromRequest(context)
}

func statusFromRequest(context echo.Context) error {
	param := context.Param("code")
	code, err := strconv.Atoi(param)
	if err != nil {
		return context.String(http.StatusBadRequest, "Illegal code "+param)
	}
	return context.String(code, FromContext(context))
}

// @Summary The request's query parameters.
// @Tags Status codes
// @Description delete
// @Accept  json
// @Produce  json
// @Param code path string false "return status code"
// @Success 200
// @Router /status/{code} [delete]
func StatusDeleteHandler(context echo.Context) error {
	return statusFromRequest(context)
}

// @Summary The request's query parameters.
// @Tags Status codes
// @Description post
// @Accept  json
// @Produce  json
// @Param code path string false "return status code"
// @Success 200
// @Router /status/{code} [post]
func StatusPostHandler(context echo.Context) error {
	return statusFromRequest(context)
}

// @Summary The request's query parameters.
// @Tags Status codes
// @Description get
// @Accept  json
// @Produce  json
// @Param code path string false "return status code"
// @Success 200
// @Router /status/{code} [put]
func StatusPutHandler(context echo.Context) error {
	return statusFromRequest(context)
}

// @Summary The request's query parameters.
// @Tags Status codes
// @Description get
// @Accept  json
// @Produce  json
// @Param code path string false "return status code"
// @Success 200
// @Router /status/{code} [patch]
func StatusPatchHandler(context echo.Context) error {
	return statusFromRequest(context)
}
