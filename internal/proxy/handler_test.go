package proxy_test

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tgunsch/httpod/internal/proxy"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T) {
	var (
		err error
		e   *echo.Echo
		req *http.Request
		res *httptest.ResponseRecorder
		ctx echo.Context
	)

	type test struct {
		inputHeaders       http.Header
		expectResponseCode int
		expectResponseBody string
	}

	tests := []test{
		{
			inputHeaders:       map[string][]string{},
			expectResponseCode: 400,
			expectResponseBody: "Invalid query input.",
		},
		{
			inputHeaders: map[string][]string{
				"url": {"://example.org"},
			},
			expectResponseCode: 400,
		},
		{
			inputHeaders: map[string][]string{
				"url": {"http://example.org"},
			},
			expectResponseCode: 200,
		},
		{
			inputHeaders: map[string][]string{
				"url":               {"http://example.org"},
				"additionalHeaders": {"c3", "po"},
			},
			expectResponseCode: 200,
		},
		{
			inputHeaders: map[string][]string{
				"url": {"http://localhost:9876"},
			},
			expectResponseCode: 500,
		},
	}

	for _, table := range tests {
		e = echo.New()
		req = httptest.NewRequest(http.MethodGet, "http://example.org/api/proxy", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		for key, value := range table.inputHeaders {
			req.Header.Set(key, value[0])
		}
		res = httptest.NewRecorder()
		ctx = e.NewContext(req, res)

		err = proxy.GetHandler(ctx)

		assert.Nil(t, err)
		if table.expectResponseBody != "" {
			assert.Equal(t, table.expectResponseBody, res.Body.String())
		}
		assert.Equal(t, table.expectResponseCode, ctx.Response().Status)

		e = nil
	}

}
