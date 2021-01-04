package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary Do a GET request.
// @Tags HTTP Methods
// @Description Requests using GET should only retrieve data.
// @Accept  json
// @Produce  json
// @Success 200 {object} http.Response
// @Router /get [get]
func GetHandler(context echo.Context) error {
	return context.String(http.StatusOK, ResponseFromContext(context))
}

// @Summary Do a DELETE request.
// @Tags HTTP Methods
// @Description The DELETE method deletes the specified resource.
// @Accept  json
// @Produce  json
// @Success 200 {object} http.Response
// @Router /delete [delete]
func DeleteHandler(context echo.Context) error {
	return context.String(http.StatusOK, ResponseFromContext(context))
}

// @Summary Do a POST request.
// @Tags HTTP Methods
// @Description The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server.
// @Accept  json
// @Produce  json
/// @Success 200 {object} http.Response
// @Router /post [post]
func PostHandler(context echo.Context) error {
	return context.String(http.StatusOK, ResponseFromContext(context))
}

// @Summary Do PUT request.
// @Tags HTTP Methods
// @Description The PUT method replaces all current representations of the target resource with the request payload.
// @Accept  json
// @Produce  json
// @Success 200 {object} http.Response
// @Router /put [put]
func PutHandler(context echo.Context) error {
	return context.String(http.StatusOK, ResponseFromContext(context))
}

// @Summary Do a PATCH request.
// @Tags HTTP Methods
// @Description The PATCH method is used to apply partial modifications to a resource.
// @Accept  json
// @Produce  json
// @Success 200 {object} http.Response
// @Router /patch [patch]
func PatchHandler(context echo.Context) error {
	return context.String(http.StatusOK, ResponseFromContext(context))
}
