package proxy

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// @Summary Do a GET request.
// @Tags Proxy Methods
// @Description Query httpod as reverse proxy to uri.
// @Accept  json
// @Produce  json
// @Param uri header string false "Full URI to use for the backend request. Mandatory. e.g. https://example.org/path "
// @Param method header string false "Method to use for the backend request. Optional, defaults to 'GET'."
// TODO @Param body header string false "Body to use for a POST request. Optional."
// TODO @Param additionalHeaders header string false "JSON of headers to add to the backend request. Optional."
// @Success 200
// @Failure 400
// @Failure 500
// @Router /proxy [get]
func GetHandler(context echo.Context) error {
	if err = Br.configureBackendRequest(context.Request()); err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	if resp, err = Br.requestBackend(); err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}

	if jsonResp, err = json.MarshalIndent(resp, "", "   "); err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}
	return context.String(http.StatusOK, string(jsonResp))

}

var (
	jsonResp []byte
	resp     *BackendResponse
	err      error
	Br       BackendRequest
)

func (br *BackendRequest) configureBackendRequest(req *http.Request) error {
	var err error
	br.Method = strings.ToUpper(req.Header.Get("method"))
	br.URI, err = url.Parse(req.Header.Get("uri"))
	if err != nil {
		return err
	}
	if br.URI.Scheme == "" || br.URI.Host == "" {
		return errors.New("invalid query input")
	}
	return nil
}

func (br *BackendRequest) requestBackend() (*BackendResponse, error) {
	var (
		backendTransport = &http.Transport{
			MaxIdleConns:       3,
			IdleConnTimeout:    5 * time.Second,
			DisableCompression: true,
		}
		backendHttpResponseBytes []byte
		backendResponse          = &BackendResponse{}
		res                      *http.Response
		err                      error
	)
	if br.HttpClient == nil {
		br.HttpClient = &http.Client{Transport: backendTransport}
	}

	if br.Request, err = http.NewRequest(br.Method, br.URI.String(), nil); err != nil {
		return nil, errors.New("Error on proxied request: " + err.Error())
	}

	if res, err = br.HttpClient.Do(br.Request); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if backendHttpResponseBytes, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, errors.New("Error parsing body: " + err.Error())
	}

	if res.Request != nil {
		backendResponse.URI = res.Request.RequestURI
	}
	backendResponse.StatusCode = res.StatusCode
	backendResponse.Headers = res.Header
	backendResponse.Body = string(backendHttpResponseBytes)

	return backendResponse, nil
}
