package proxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"time"
)

// @Summary Do a GET request.
// @Tags Proxy Methods
// @Description Requests using proxy GET should query target.
// @Accept  json
// @Produce  json
// @Success 200 {object} http.Response
// @Router /proxy/{target} [get]
func GetHandler(context echo.Context) error {
	var (
		jsonResp []byte
		resp     *Response
		err      error
	)

	if resp, err = responseFromTarget(context); err != nil {
		return err
	}

	if jsonResp, err = json.MarshalIndent(resp, "", "   "); err != nil {
		return errors.New("JSON error: " + err.Error())
	}
	return context.String(http.StatusOK, string(jsonResp))

}

func responseFromTarget(context echo.Context) (*Response, error) {
	var (
		target         string
		tr             *http.Transport
		c              *http.Client
		targetReq      *http.Request
		targetResp     *http.Response
		resp           *Response
		targetRespBody []byte
		err            error
	)

	target = context.Param("target")
	target = fmt.Sprintf("http://%s", target) // TODO Support url / paths

	tr = &http.Transport{
		MaxIdleConns:       3,
		IdleConnTimeout:    5 * time.Second,
		DisableCompression: true,
	}

	c = &http.Client{Transport: tr}

	if targetReq, err = http.NewRequest(http.MethodGet, target, nil); err != nil {
		// TODO Return response code
		return nil, errors.New("Illegal request: " + target)
	}

	// TODO Forward headers
	// targetReq.Header.Add("", "")

	if targetResp, err = c.Do(targetReq); err != nil {
		return nil, errors.New("Invalid response: " + err.Error())
	}
	defer targetResp.Body.Close()

	if targetRespBody, err = ioutil.ReadAll(targetResp.Body); err != nil {
		return nil, errors.New("Error parsing body: " + err.Error())
	}

	resp = &Response{
		Code: targetResp.StatusCode,
		Url:  targetReq.URL.RequestURI(),
		//Headers: targetResp.Header, // TODO Iterate
		Body: string(targetRespBody),
	}

	return resp, nil
}
