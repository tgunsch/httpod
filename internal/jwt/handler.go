package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/tgunsch/httpod/internal/util"
	"net/http"
)

// @Summary Get jwt passed as authorization bearer token of the request.
// @Tags JWT
// @Description Requests using GET should only retrieve data.
// @Accept  json
// @Produce  json
// @Param validate query bool false "if true, the jwt is validated"
// @Param jwksUri query string false "if set, the jwt is verified with the key received from jwks endpoint"
// @Success 200 {array} jwt.Token
// @Router /jwt [get]
func GetHandler(ctx echo.Context) error {
	var (
		auth       string
		keys       jwk.Set
		err        error
		response   *Response
		prettyJSON []byte
		validate   bool
	)

	auth = ctx.Request().Header.Get(echo.HeaderAuthorization)
	l := len("Bearer")
	if auth[:l] == "Bearer" {
		rawToken := auth[l+1:]

		jwksUri := ctx.QueryParam("jwksUri")
		if jwksUri != "" {
			if keys, err = jwk.Fetch(context.Background(), jwksUri); err != nil {
				return ctx.String(http.StatusBadRequest, fmt.Sprintf("failed to validate token: %s\n", err))
			}
		}
		validate = util.GetBoolParam(ctx, "validate")
		if response, err = NewResponse(rawToken, validate, keys); err != nil {
			return ctx.String(http.StatusBadRequest, fmt.Sprintf("failed to parse payload: %s\n", err))
		}

		if prettyJSON, err = json.MarshalIndent(response, "", "   "); err != nil {
			return ctx.String(http.StatusBadRequest, fmt.Sprintf("Error parsing cookies: %v", err.Error()))
		}
		return ctx.String(http.StatusOK, string(prettyJSON))

	}
	return ctx.String(http.StatusBadRequest, "No JWT in request header")

}
