package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"net/http"
)

// @Summary Get jwt passed as header. By default from Authorization bearer header of the request.
// @Tags JWT
// @Description Requests using GET should only retrieve data.
// @Accept  json
// @Produce  json
// @Param headerName query string false "if set, JWT is read from this header name. Otherwise from Authorization header"
// @Param jwksUri query string false "if set, the jwt is verified with the key received from jwks endpoint"
// @Success 200 {array} jwt.Response
// @Router /jwt [get]
func GetHandler(ctx echo.Context) error {
	var (
		auth       string
		headerName string
		rawToken   string
		keys       jwk.Set
		err        error
		response   *Response
		prettyJSON []byte
	)
	headerName = ctx.QueryParam("headerName")
	if headerName == "" {
		headerName = echo.HeaderAuthorization
	}
	auth = ctx.Request().Header.Get(headerName)
	if auth == "" {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("No %s header in request", headerName))
	}

	l := len("Bearer")
	if auth[:l] == "Bearer" {
		rawToken = auth[l+1:]
	} else {
		rawToken = auth
	}
	jwksUri := ctx.QueryParam("jwksUri")
	if jwksUri != "" {
		if keys, err = jwk.Fetch(context.Background(), jwksUri); err != nil {
			return ctx.String(http.StatusBadRequest, fmt.Sprintf("failed to validate token: %s\n", err))
		}
	}
	if response, err = NewResponse(rawToken, keys); err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("failed to parse payload: %s\n", err))
	}

	if prettyJSON, err = json.MarshalIndent(response, "", "   "); err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Error parsing cookies: %v", err.Error()))
	}
	return ctx.String(http.StatusOK, string(prettyJSON))

}
