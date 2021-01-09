package cookies_test

import (
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tgunsch/httpod/internal/cookies"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("GetHandler", func() {
	It("return a empty list if no cookies exists", func() {

		ctx, _, responseRecorder := mockGetContext()

		err := cookies.GetHandler(ctx)
		Expect(err).Should(BeNil())

		// return 200
		Expect(responseRecorder.Code).Should(Equal(200))

		// response body contains json cookie
		Expect(responseRecorder.Body.String()).To(MatchJSON(`[]`))
	})
	It("return one cookie", func() {

		ctx, req, responseRecorder := mockGetContext()
		req.AddCookie(&http.Cookie{
			Name:  "testCookie",
			Value: "testValue",
		})

		err := cookies.GetHandler(ctx)
		Expect(err).Should(BeNil())

		// return 200
		Expect(responseRecorder.Code).Should(Equal(200))

		// response body contains json cookie
		Expect(responseRecorder.Body.String()).To(MatchJSON(`[{
				"name": "testCookie", 
				"value": "testValue"
			}]`))
	})
	It("return multiple cookies", func() {

		ctx, req, responseRecorder := mockGetContext()
		req.AddCookie(&http.Cookie{
			Name:  "testCookie1",
			Value: "testValue1",
		})
		req.AddCookie(&http.Cookie{
			Name:  "testCookie2",
			Value: "testValue2",
		})

		err := cookies.GetHandler(ctx)
		Expect(err).Should(BeNil())

		// return 200
		Expect(responseRecorder.Code).Should(Equal(200))

		// response body contains json cookie
		Expect(responseRecorder.Body.String()).To(MatchJSON(`[
			{
				"name": "testCookie1", 
				"value": "testValue1"
			},
			{
				"name": "testCookie2", 
				"value": "testValue2"
			}]`))
	})

})

func mockGetContext() (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://myapp.com/api/cookies", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	return c, req, res
}
