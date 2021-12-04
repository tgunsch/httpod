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
	HttpClient        *http.Client
}

type BackendResponse struct {
	StatusCode int
	URI        string
	Headers    *http.Header
	Body       string
}
