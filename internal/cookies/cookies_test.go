package cookies_test

import (
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/tgunsch/httpod/internal/cookies"
)

var _ = Describe("Cookies", func() {
	Context("PostHandler", func() {
		It("creates cookie with default values", func() {

			testCookie := `{"value":"testValue"}`
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "http://myapp.com/api/cookies", strings.NewReader(testCookie))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("cookieName")
			c.SetParamValues("testCookie")

			cookies.PostHandler(c)

			// return 200
			Expect(rec.Code).Should(Equal(200))

			// response has set-cookie
			setCookieHeader := rec.Result().Header["Set-Cookie"][0]
			Expect(setCookieHeader).Should(Equal("testCookie=testValue; Domain=myapp.com; SameSite"))

			// response body contains json cookie
			Expect(rec.Body.String()).To(MatchJSON(`{ "name": "testCookie", "value": "testValue", "domain": "myapp.com" }`))
		})
		It("creates cookie with specific values", func() {

			testCookie := `{"value":"testValue", "expiresSeconds":3600 }`
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "http://myapp.com/api/cookies", strings.NewReader(testCookie))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("cookieName")
			c.SetParamValues("testCookie")

			expireTime := time.Now().Add(time.Second * time.Duration(3600))

			cookies.PostHandler(c)

			// return 200
			Expect(rec.Code).Should(Equal(200))

			// response has set-cookie
			setCookieHeader := rec.Result().Header["Set-Cookie"][0]
			expireString := expireTime.UTC().Format(http.TimeFormat)
			Expect(setCookieHeader).Should(Equal("testCookie=testValue; Domain=myapp.com; Expires=" + expireString + "; SameSite"))

			// response body contains json cookie
			expireJson := expireTime.Format(cookies.TIME_FORMAT)
			Expect(rec.Body.String()).To(MatchJSON(`{ "name": "testCookie", "value": "testValue", "domain": "myapp.com", "expires": "` + expireJson + `" }`))
		})
	})
})
