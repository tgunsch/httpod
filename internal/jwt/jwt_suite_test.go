package jwt_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJWT(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JWT Suite")
}
