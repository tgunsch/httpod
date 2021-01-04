package cookies_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCookies(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cookies Suite")
}
