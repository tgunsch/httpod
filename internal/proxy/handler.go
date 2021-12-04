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
// @Description Requests using proxy GET should query target.
// @Accept  json
// @Produce  json
// @Param method header string false "Method to use for the backend request. Defaults to 'GET'."
// @Param url header string false "URL to use for the backend request. Mandatory."
// @Param body header string false "Body to use for a POST request."
// @Param additionalHeaders header string false "JSON of headers to add to the backend request."
// @Success 200 {object} http.Response
// @Router /proxy [get]
func GetHandler(context echo.Context) error {
	var (
		jsonResp []byte
		resp     *BackendResponse
		err      error
		br       BackendRequest
	)
	if err = br.configureBackendRequest(context.Request()); err != nil {
		return context.String(http.StatusBadRequest, string(err.Error()))
	}

	if resp, err = br.requestBackend(); err != nil {
		return context.String(http.StatusInternalServerError, string(err.Error()))
	}

	if jsonResp, err = json.MarshalIndent(resp, "", "   "); err != nil {
		return context.String(http.StatusInternalServerError, string(err.Error()))
	}
	return context.String(http.StatusOK, string(jsonResp))

}

func (br *BackendRequest) configureBackendRequest(req *http.Request) error {
	var err error
	br.Method = strings.ToUpper(req.Header.Get("method"))
	if br.Method == "" {
		br.Method = "GET"
	}
	br.Url, err = url.Parse(req.Header.Get("url"))
	if err != nil {
		return err
	}
	if br.Url.Scheme == "" || br.Url.Host == "" {
		return errors.New("Invalid query input.")
	}
	additionalHeaders := []byte(req.Header.Get("additionalHeaders")) // TODO
	json.Unmarshal(additionalHeaders, &br.AdditionalHeaders)

	return nil
}

func (br *BackendRequest) requestBackend() (*BackendResponse, error) {
	var (
		backendTransport = &http.Transport{
			MaxIdleConns:       3,
			IdleConnTimeout:    5 * time.Second,
			DisableCompression: true,
		}
		backendClient            = &http.Client{Transport: backendTransport}
		backendHttpResponseBytes []byte
		backendResponse          = &BackendResponse{}
		res                      *http.Response
		err                      error
	)

	if br.Request, err = http.NewRequest(br.Method, br.Url.String(), nil); err != nil {
		return nil, errors.New("Error on proxied request: " + err.Error())
	}

	for _, header := range br.AdditionalHeaders {
		br.Request.Header.Add(header, br.AdditionalHeaders["header"])
	}

	if res, err = backendClient.Do(br.Request); err != nil {
		return nil, errors.New("Invalid response: " + err.Error())
	}
	defer res.Body.Close()

	if backendHttpResponseBytes, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, errors.New("Error parsing body: " + err.Error())
	}

	backendResponse.Code = res.StatusCode
	backendResponse.Url = res.Request.RequestURI
	backendResponse.Headers = res.Header
	backendResponse.Body = string(backendHttpResponseBytes)

	//fmt.Printf("Read backend response: %+v \n", backendResponse)

	return backendResponse, nil
}
