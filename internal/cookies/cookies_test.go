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

			ctx, _, responseRecorder := mockContext(`{"value":"testValue"}`)

			err := cookies.PostHandler(ctx)
			Expect(err).Should(BeNil())

			// return 200
			Expect(responseRecorder.Code).Should(Equal(200))

			// response has set-cookie
			setCookieHeader := responseRecorder.Result().Header["Set-Cookie"][0]
			Expect(setCookieHeader).Should(Equal("testCookie=testValue; Domain=myapp.com; SameSite"))

			// response body contains json cookie
			Expect(responseRecorder.Body.String()).To(MatchJSON(`{ "name": "testCookie", "value": "testValue", "domain": "myapp.com" }`))
		})
		It("creates cookie with specific values", func() {

			ctx, _, responseRecorder := mockContext(`{
				"value":"testValue",
				"path": "/blubb",
				"expiresSeconds":3600,
				"secure": true,
				"httpOnly": true,
				"sameSite": "strict"
				}`)

			expireTime := time.Now().Add(time.Second * time.Duration(3600))

			err := cookies.PostHandler(ctx)
			Expect(err).Should(BeNil())

			// return 200
			Expect(responseRecorder.Code).Should(Equal(200))

			// response has set-cookie
			setCookieHeader := responseRecorder.Result().Header["Set-Cookie"][0]
			expireString := expireTime.UTC().Format(http.TimeFormat)
			Expect(setCookieHeader).Should(Equal("testCookie=testValue; Path=/blubb; Domain=myapp.com; Expires=" + expireString + "; HttpOnly; Secure; SameSite=Strict"))
			// response body contains json cookie
			expireJson := expireTime.Format(cookies.TIME_FORMAT)
			Expect(responseRecorder.Body.String()).To(MatchJSON(`{
				"name": "testCookie", 
				"value": "testValue", 
				"path": "/blubb",
				"domain": "myapp.com", 
				"expires": "` + expireJson + `", 
				"secure": true,
          		"httpOnly": true,
          		"sameSite": "Strict"
			}`))
		})
		It("creates cookie max age higher priority than expires", func() {

			ctx, _, responseRecorder := mockContext(`{
				"value":"testValue",
				"expiresSeconds":3600,
				"maxAge": 1
				}`)

			err := cookies.PostHandler(ctx)
			Expect(err).Should(BeNil())

			// return 200
			Expect(responseRecorder.Code).Should(Equal(200))

			// response has set-cookie
			setCookieHeader := responseRecorder.Result().Header["Set-Cookie"][0]
			Expect(setCookieHeader).Should(Equal("testCookie=testValue; Domain=myapp.com; Max-Age=1; SameSite"))
			// response body contains json cookie
			Expect(responseRecorder.Body.String()).To(MatchJSON(`{
				"name": "testCookie", 
				"value": "testValue", 
				"domain": "myapp.com", 
				"maxAge": 1
			}`))
		})
		It("don't overwrite existing cookie", func() {

			ctx, req, responseRecorder := mockContext(`{"value":"testValue"}`)
			req.AddCookie(&http.Cookie{
				Name: "testCookie",
				Path: "/",
			})
			err := cookies.PostHandler(ctx)
			Expect(err).Should(BeNil())

			// return 200
			Expect(responseRecorder.Code).Should(Equal(400))
			Expect(responseRecorder.Body.String()).Should(Equal("Cookie testCookie already exists"))
		})

	})
})

func mockContext(testCookie string) (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "http://myapp.com/api/cookies", strings.NewReader(testCookie))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("cookieName")
	c.SetParamValues("testCookie")
	return c, req, res
}
