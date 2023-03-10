package jwt_test

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	jwt2 "github.com/lestrrat-go/jwx/jwt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tgunsch/httpod/internal/jwt"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe("GetHandler", func() {
	It("return a jwt", func() {

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
        },
        "valid": false
      }`))
	})
	It("return a jwt with verified signature", func() {

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		key := createSymmetricKey("your-256-bit-secret", "12345")

		token := createToken(key)

		mockJWKSEndpoint(key, "http://my.jwks.com/jwks")

		ctx, _, responseRecorder := mockGetContext("http://myapp.com/api/jwt?jwksUri=http%3A%2F%2Fmy.jwks.com%2Fjwks", token)

		err := jwt.GetHandler(ctx)
		Expect(err).Should(BeNil())

		// return 200
		Expect(responseRecorder.Code).Should(Equal(200))

		// response body contains json cookie
		Expect(responseRecorder.Body.String()).To(MatchJSON(`{
        "raw": "eyJhbGciOiJIUzI1NiIsImtpZCI6IjEyMzQ1IiwidHlwIjoiSldUIn0.eyJleHAiOjg3ODQ1NDg0MCwiaWF0Ijo0NTMyNzkxMjAsImlzcyI6InNreW5ldCIsInN1YiI6IlQtODAwIn0.GWYi_xOOQG3xzH6zhRFbomaIZ4xra6Accn0FhaZ_87A",
        "header": {
          "alg": "HS256",
          "kid": "12345",
          "typ": "JWT"
        },
        "payload": {
          "exp": "1997-11-02T07:14:00Z",
          "iat": "1984-05-13T06:52:00Z",
          "iss": "skynet",
          "sub": "T-800"
        },
        "valid": false,
        "verifiedSignature": true
      }`))
	})

})

func createSymmetricKey(value string, kid string) jwk.Key {
	raw := []byte(value)
	key, err := jwk.New(raw)
	Expect(err).Should(BeNil())
	_ = key.Set("kid", kid)
	_ = key.Set("alg", jwa.HS256)
	return key
}

func createToken(key jwk.Key) string {
	loc, _ := time.LoadLocation("EST")
	token := jwt2.New()
	_ = token.Set(jwt2.IssuerKey, "skynet")
	_ = token.Set(jwt2.ExpirationKey, time.Date(1997, 8, 94, 2, 14, 0, 0, loc).Unix())
	_ = token.Set(jwt2.IssuedAtKey, time.Date(1984, 05, 13, 1, 52, 0, 0, loc).Unix())
	_ = token.Set(jwt2.SubjectKey, "T-800")
	headers := jws.NewHeaders()
	_ = headers.Set(jws.KeyIDKey, key.KeyID())
	signedToken, _ := jwt2.Sign(token, jwa.HS256, key, jwt2.WithHeaders(headers))
	return string(signedToken)
}

func mockGetContext(uri string, token string) (echo.Context, *http.Request, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, uri, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	return c, req, res
}

func mockJWKSEndpoint(key jwk.Key, jwksUrl string) {

	set := jwk.NewSet()
	set.Add(key)

	responder := func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, set)
		return resp, err
	}
	httpmock.RegisterResponder("GET", jwksUrl, responder)

}
