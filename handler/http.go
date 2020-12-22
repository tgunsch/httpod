package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary The request's query parameters.
// @Tags HTTP Methods
// @Description get
// @Accept  json
// @Produce  json
// @Success 200
// @Router /get [get]
func HttpGetHandler(context echo.Context) error {
	return context.String(http.StatusOK, FromContext(context))
}

// @Summary The request's query parameters.
// @Tags HTTP Methods
// @Description delete
// @Accept  json
// @Produce  json
// @Success 200
// @Router /delete [delete]
func HttpDeleteHandler(context echo.Context) error {
	return context.String(http.StatusOK, FromContext(context))
}

// @Summary The request's query parameters.
// @Tags HTTP Methods
// @Description post
// @Accept  json
// @Produce  json
// @Success 200
// @Router /post [post]
func HttpPostHandler(context echo.Context) error {
	return context.String(http.StatusOK, FromContext(context))
}

// @Summary The request's query parameters.
// @Tags HTTP Methods
// @Description get
// @Accept  json
// @Produce  json
// @Success 200
// @Router /put [put]
func HttpPutHandler(context echo.Context) error {
	return context.String(http.StatusOK, FromContext(context))
}

// @Summary The request's query parameters.
// @Tags HTTP Methods
// @Description get
// @Accept  json
// @Produce  json
// @Success 200
// @Router /patch [patch]
func HttpPatchHandler(context echo.Context) error {
	return context.String(http.StatusOK, FromContext(context))
}
