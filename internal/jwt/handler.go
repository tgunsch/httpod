package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary Get jwt passed as authorization bearer token of the request.
// @Tags JWT
// @Description Requests using GET should only retrieve data.
// @Accept  json
// @Produce  json
// @Success 200 {array} jwt.Token
// @Router /jwt [get]
func GetHandler(context echo.Context) error {

	auth := context.Request().Header.Get(echo.HeaderAuthorization)
	l := len("Bearer")
	if auth[:l] == "Bearer" {
		rawToken := auth[l+1:]

		response, err := NewResponse(rawToken)
		if err != nil {
			return context.String(http.StatusBadRequest, fmt.Sprintf("failed to parse payload: %s\n", err))
		}


		prettyJSON, err := json.MarshalIndent(response, "", "   ")
		if err != nil {
			return context.String(http.StatusBadRequest, fmt.Sprintf("Error parsing cookies: %v", err.Error()))
		}
		return context.String(http.StatusOK, string(prettyJSON))

	}
	return context.String(http.StatusBadRequest, "No JWT in request header")

}
