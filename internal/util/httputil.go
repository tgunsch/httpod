package util

import (
	"net"
	"net/http"
	"strings"
)

func GetSchemeHost(request *http.Request) (string, string) {

	host := getHost(request)
	scheme := getScheme(request, host)
	return scheme, host
}

func GetPrefix(path string, request *http.Request) string {
	forwardedPrefix := getForwardedHeader(request, "Prefix")
	if forwardedPrefix != "" {
		path = join(forwardedPrefix, path)
	}
	return path
}

func join(first string, second string) string {
	result := first
	if !strings.HasSuffix(result, "/") {
		result = result + "/"
	}
	return result + strings.TrimPrefix(second, "/")
}

func getScheme(request *http.Request, host string) string {
	port := ""
	if _, p, err := net.SplitHostPort(host); err == nil {
		port = p
	}

	forwardedProto := getForwardedHeader(request, "Proto")
	scheme := ""
	switch {
	case forwardedProto != "":
		scheme = forwardedProto
	case request.TLS != nil:
		scheme = "https"
	case request.URL.Scheme != "":
		scheme = request.URL.Scheme
	case port == "443":
		scheme = "https"
	default:
		scheme = "http"
	}
	return scheme
}

func getHost(request *http.Request) string {
	host := ""
	forwardedHost := getForwardedHeader(request, "Host")
	switch {
	case forwardedHost != "":
		host = forwardedHost
	case request.Host != "":
		host = request.Host
	case request.URL.Host != "":
		host = request.URL.Host

	}
	return host
}

func getForwardedHeader(req *http.Request, prefix string) string {
	headerList := req.Header.Get("X-Forwarded-" + prefix)
	headerValue := strings.SplitN(headerList, ",", 2)[0]
	return strings.TrimSpace(headerValue)
}
