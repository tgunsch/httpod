package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
	"strings"
)

// @Summary Get jwt of the request.
// @Tags JWT
// @Description Requests using GET should only retrieve data.
// @Accept  json
// @Produce  json
// @Success 200 {array} cookies.GetCookies
// @Router /jwt [get]
func GetHandler(context echo.Context) error {

	auth := context.Request().Header.Get(echo.HeaderAuthorization)
	l := len("Bearer")
	if auth[:l] == "Bearer" {
		rawToken := auth[l+1:]
		token, err := jwt.ParseReader(strings.NewReader(rawToken))
		if err != nil {
			return context.String(http.StatusBadRequest, fmt.Sprintf("failed to parse payload: %s\n", err))
		}
		buf, err := json.MarshalIndent(token, "", "  ")
		if err != nil {
			return context.String(http.StatusBadRequest, fmt.Sprintf("failed to generate JSON: %s\n", err))
		}
		return context.String(http.StatusOK, string(buf))
	}
	return context.String(http.StatusBadRequest, "No JWT in request header")

}
