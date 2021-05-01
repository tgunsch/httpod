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
	It("return a not validated jwt", func() {

		ctx, _, responseRecorder := mockGetContext("http://myapp.com/api/jwt", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImp0aSI6Ijg4MmY3MWEzLWRiM2EtNGE2Ny05NTllLTZmZDE3MmFhYWNhMCIsImlhdCI6MTYxOTM0NjE2MywiZXhwIjoxNjE5MzQ5NzYzfQ.3GRfe59wu2KuXJyZV0uGqxpX6WWdeQTEsARbwow_ZG4")

		err := jwt.GetHandler(ctx)
		Expect(err).Should(BeNil())

		// return 200
		Expect(responseRecorder.Code).Should(Equal(200))

		// response body contains json cookie
		Expect(responseRecorder.Body.String()).To(MatchJSON(`{
        "raw": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImp0aSI6Ijg4MmY3MWEzLWRiM2EtNGE2Ny05NTllLTZmZDE3MmFhYWNhMCIsImlhdCI6MTYxOTM0NjE2MywiZXhwIjoxNjE5MzQ5NzYzfQ.3GRfe59wu2KuXJyZV0uGqxpX6WWdeQTEsARbwow_ZG4",
        "header": {
          "alg": "HS256",
          "typ": "JWT"
        },
        "payload": {
          "admin": true,
          "exp": "2021-04-25T11:22:43Z",
          "iat": "2021-04-25T10:22:43Z",
          "jti": "882f71a3-db3a-4a67-959e-6fd172aaaca0",
          "name": "John Doe",
          "sub": "1234567890"
        }
      }`))
	})
	It("return a validated jwt", func() {

		ctx, _, responseRecorder := mockGetContext("http://myapp.com/api/jwt?validate=true", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImp0aSI6Ijg4MmY3MWEzLWRiM2EtNGE2Ny05NTllLTZmZDE3MmFhYWNhMCIsImlhdCI6MTYxOTM0NjE2MywiZXhwIjoxNjE5MzQ5NzYzfQ.3GRfe59wu2KuXJyZV0uGqxpX6WWdeQTEsARbwow_ZG4")

		err := jwt.GetHandler(ctx)
		Expect(err).Should(BeNil())

		// return 200
		Expect(responseRecorder.Code).Should(Equal(200))

		// response body contains json cookie
		Expect(responseRecorder.Body.String()).To(MatchJSON(`{
        "raw": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImp0aSI6Ijg4MmY3MWEzLWRiM2EtNGE2Ny05NTllLTZmZDE3MmFhYWNhMCIsImlhdCI6MTYxOTM0NjE2MywiZXhwIjoxNjE5MzQ5NzYzfQ.3GRfe59wu2KuXJyZV0uGqxpX6WWdeQTEsARbwow_ZG4",
        "header": {
          "alg": "HS256",
          "typ": "JWT"
        },
        "payload": {
          "admin": true,
          "exp": "2021-04-25T11:22:43Z",
          "iat": "2021-04-25T10:22:43Z",
          "jti": "882f71a3-db3a-4a67-959e-6fd172aaaca0",
          "name": "John Doe",
          "sub": "1234567890"
        },
        "valid": false,
        "validateError": "exp not satisfied"
      }`))
	})

})

func mockGetContext(uri string, token string) (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, uri, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	return c, req, res
}
