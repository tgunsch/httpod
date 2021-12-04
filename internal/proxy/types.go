package proxy

import (
	"net/http"
	"net/url"
)

type BackendRequest struct {
	Method            string
	Url               *url.URL
	AdditionalHeaders map[string]string
	Request           *http.Request
}

type BackendResponse struct {
	Code    int
	Url     string
	Headers http.Header
	Body    string
}
