package cookies

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tgunsch/httpod/internal/util"
	"net/http"
	"strings"
	"time"
)

const TIME_FORMAT = "2006-01-02T15:04:05Z07:00"

type JSONTime struct {
	time.Time
}

type GetCookies struct {
	Name  string `json:"name"`
	Value string `json:"value"`

	Path       string    `json:"path,omitempty"`
	Domain     string    `json:"domain,omitempty"`
	Expires    *JSONTime `json:"expires,omitempty"`
	RawExpires string    `json:"rawExpires,omitempty"`
	MaxAge     int       `json:"maxAge,omitempty"`
	Secure     bool      `json:"secure,omitempty"`
	HttpOnly   bool      `json:"httpOnly,omitempty"`
	SameSite   string    `json:"sameSite,omitempty"`
}

type SetCookie struct {
	Value          string `json:"value" example:"Test"`
	Path           string `json:"path,omitempty" example:"/"`
	ExpiresSeconds int    `json:"expiresSeconds,omitempty" example:"3600"`
	MaxAge         int    `json:"maxAge,omitempty" example:"0"`
	Secure         bool   `json:"secure,omitempty" example:"true"`
	HttpOnly       bool   `json:"httpOnly,omitempty" example:"true"`
	SameSite       string `json:"sameSite,omitempty"  example:"Strict"`
}

func toJsonCookie(cookie *http.Cookie) GetCookies {
	var expires *JSONTime = nil
	if cookie.Expires.After(time.Time{}) {
		expires = &JSONTime{cookie.Expires}
	}
	return GetCookies{
		Name:       cookie.Name,
		Value:      cookie.Value,
		Path:       cookie.Path,
		Domain:     cookie.Domain,
		Expires:    expires,
		RawExpires: cookie.RawExpires,
		MaxAge:     cookie.MaxAge,
		Secure:     cookie.Secure,
		HttpOnly:   cookie.HttpOnly,
		SameSite:   sameSiteString(cookie.SameSite),
	}
}

func toHttpCookie(context echo.Context) (*http.Cookie, error) {
	cookie := new(SetCookie)
	if err := context.Bind(cookie); err != nil {
		return nil, err
	}
	name := context.Param("cookieName")

	host := util.GetHost(context.Request())

	expires := time.Time{}
	maxAge := cookie.MaxAge
	if maxAge <= 0 {
		maxAge = 0
		if cookie.ExpiresSeconds > 0 {
			expires = time.Now().Local().Add(time.Second * time.Duration(cookie.ExpiresSeconds))
		}
	}
	return &http.Cookie{
		Name:     name,
		Value:    cookie.Value,
		Path:     cookie.Path,
		Domain:   host,
		Expires:  expires,
		MaxAge:   maxAge,
		Secure:   cookie.Secure,
		HttpOnly: cookie.HttpOnly,
		SameSite: sameSite(cookie.SameSite),
	}, nil
}

func sameSiteString(s http.SameSite) string {
	switch s {
	case http.SameSiteDefaultMode:
		return ""
	case http.SameSiteNoneMode:
		return "None"
	case http.SameSiteLaxMode:
		return "Lax"
	case http.SameSiteStrictMode:
		return "Strict"
	}
	return ""
}

func sameSite(s string) http.SameSite {
	lowerVal := strings.ToLower(s)
	switch lowerVal {
	case "lax":
		return http.SameSiteLaxMode
	case "strict":
		return http.SameSiteStrictMode
	case "none":
		return http.SameSiteNoneMode
	default:
		return http.SameSiteDefaultMode
	}
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", t.Format(TIME_FORMAT))
	return []byte(stamp), nil
}
