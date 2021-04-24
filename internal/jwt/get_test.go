package jwt_test

import (
	"fmt"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tgunsch/httpod/internal/jwt"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("GetHandler", func() {
	It("return a jwt", func() {

		ctx, _, responseRecorder := mockGetContext("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")

		err := jwt.GetHandler(ctx)
		Expect(err).Should(BeNil())

		// return 200
		Expect(responseRecorder.Code).Should(Equal(200))

		// response body contains json cookie
		Expect(responseRecorder.Body.String()).To(MatchJSON(`{
        "iat": 1516239022,
        "name": "John Doe",
        "sub": "1234567890"
      }`))
	})

})

func mockGetContext(token string) (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://myapp.com/api/jwt", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	return c, req, res
}
