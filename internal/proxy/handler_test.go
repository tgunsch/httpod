package proxy_test

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tgunsch/httpod/internal/proxy"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T) {
	var (
		err                   error
		e                     *echo.Echo
		req                   *http.Request
		res                   *httptest.ResponseRecorder
		ctx                   echo.Context
		backendResponseHeader = make(http.Header)
		backendResponseBody   = ""
		backendResponseCode   = 0
	)

	type test struct {
		inputHeaders       http.Header
		expectResponseCode int
		httpClient         *http.Client
	}

	tests := []test{
		{
			inputHeaders:       map[string][]string{},
			expectResponseCode: 400,
		},
		{
			inputHeaders: map[string][]string{
				"url": {"://localhost"},
			},
			expectResponseCode: 400,
		},
		{
			inputHeaders: map[string][]string{
				"url": {"http://localhost"},
			},
			expectResponseCode: 200,
		},
		{
			inputHeaders: map[string][]string{
				"url": {"http://localhost:9876"},
			},
			expectResponseCode: 500,
			httpClient:         http.DefaultClient,
		},
	}

	for _, table := range tests {
		e = echo.New()
		req = httptest.NewRequest(http.MethodGet, "http://localhost/api/proxy", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		for key, value := range table.inputHeaders {
			req.Header.Set(key, value[0])
		}
		res = httptest.NewRecorder()
		ctx = e.NewContext(req, res)

		if table.httpClient == nil {
			proxy.Br = proxy.BackendRequest{
				HttpClient: newTestClient(backendResponseCode, backendResponseHeader, backendResponseBody),
			}
		}

		err = proxy.GetHandler(ctx)

		assert.Nil(t, err)
		assert.Equal(t, table.expectResponseCode, ctx.Response().Status)
		fmt.Printf("%+v \n", ctx.Response().Writer)

		proxy.Br = proxy.BackendRequest{}
		e = nil
	}

}

// Based on http://hassansin.github.io/Unit-Testing-http-client-in-Go
type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(statusCode int, header http.Header, response string) *http.Client {
	return &http.Client{
		Transport: roundTripFunc(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: statusCode,
				Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
				Header:     header,
			}
		}),
	}
}
