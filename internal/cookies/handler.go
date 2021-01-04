package cookies

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary Get all cookies of the request.
// @Tags Cookies
// @Description Requests using GET should only retrieve data.
// @Accept  json
// @Produce  json
// @Success 200 {array} cookies.GetCookies
// @Router /cookies [get]
func GetHandler(context echo.Context) error {
	cookies := context.Cookies()
	getCookies := make([]GetCookies, len(cookies))
	for i, cookie := range cookies {
		getCookies[i] = toJsonCookie(cookie)
	}
	prettyJSON, err := json.MarshalIndent(getCookies, "", "   ")
	if err != nil {
		return context.String(http.StatusBadRequest, fmt.Sprintf("Error parsing cookies: %v", err.Error()))
	}
	return context.String(http.StatusOK, string(prettyJSON))
}

// @Summary Delete a cookie.
// @Tags Cookies
// @Description Delete a specific cookie.
// @Accept  json
// @Produce  json
// @Param cookieName path string false "The name of the cookie to delete"
// @Success 200 {object} cookies.GetCookies
// @Router /cookies/{cookieName} [delete]
func DeleteHandler(context echo.Context) error {
	name := context.Param("cookieName")
	cookie := &http.Cookie{
		Name:   name,
		MaxAge: -1,
		Path:   "/",
	}
	context.SetCookie(cookie)
	getCookie := toJsonCookie(cookie)
	prettyJSON, err := json.MarshalIndent(getCookie, "", "   ")
	if err != nil {
		return context.String(http.StatusBadRequest, fmt.Sprintf("Error parsing cookies: %v", err.Error()))
	}
	return context.String(http.StatusOK, string(prettyJSON))
}

// @Summary Create a new cookie.
// @Tags Cookies
// @Description
// @Accept  json
// @Produce  json
// @Param cookieName path string false "The name of the new cookie"
// @Param cookie body cookies.SetCookie true "The cookie"
// @Success 200 {object} cookies.GetCookies
// @Router /cookies/{cookieName} [post]
func PostHandler(context echo.Context) error {

	cookie, err := toHttpCookie(context)
	if err != nil {
		return context.String(http.StatusBadRequest, fmt.Sprintf("Oops: %v", err))
	}
	context.SetCookie(cookie)
	jsonCookie := toJsonCookie(cookie)
	prettyJSON, err := json.MarshalIndent(jsonCookie, "", "   ")
	if err != nil {
		return context.String(http.StatusBadRequest, fmt.Sprintf("Error parsing cookies: %v", err.Error()))
	}
	return context.String(http.StatusOK, string(prettyJSON))
}
