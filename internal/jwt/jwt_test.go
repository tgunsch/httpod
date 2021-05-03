package jwt_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("JWT", func() {
	It("validate jwt", func() {
		header := `{"alg":"HS256","typ":"JWT"}`
		payload := `{"sub":"1234567890","name":"John Doe","iat":1516239022}`
		unsignedToken := Base64Encode(header) + "." + Base64Encode(payload)
		Expect(unsignedToken).To(Equal("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ"))

		key := `your-256-bit-secret`
		mac := hmac.New(sha256.New, []byte(key))
		mac.Write([]byte(unsignedToken))
		signature := mac.Sum(nil)

		token := unsignedToken + "." + Base64Encode(string(signature))
		Expect(token).To(Equal("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"))

	})
})

func Base64Encode(src string) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString([]byte(src)), "=")
}
