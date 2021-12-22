package proxy

import (
	"net/http"
	"net/url"
)

type BackendRequest struct {
	Method     string
	URI        *url.URL
	Request    *http.Request
	HttpClient *http.Client
}

type BackendResponse struct {
	StatusCode int
	URI        string
	Headers    http.Header
	Body       string
}
