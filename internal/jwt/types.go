package jwt

import (
	"bytes"
	"context"
	"fmt"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
	"strings"
)

type Response struct {
	//Method    SigningMethod          // The signing method used or to be used
	Header    map[string]interface{} `json:"header"`
	Payload   jwt.Token                 `json:"payload"`
	Signature string                 `json:"signature"`
	Valid     bool                   `json:"valid"`
}


func NewResponse(rawToken string) (*Response,error) {


	msg, err := jws.ParseReader(strings.NewReader(rawToken))
	if err != nil {
		return nil, fmt.Errorf("failed to parse token data: %v",err)
	}

	token, err := jwt.ParseReader(bytes.NewReader(msg.Payload()))
	if err != nil {
		return nil,fmt.Errorf("failed to parse payload: %s\n", err)
	}

	var header = make(map[string]interface{})
	signatures := msg.Signatures()
	for _, signature := range signatures {
		header, _ = signature.ProtectedHeaders().AsMap(context.TODO())
	}

	return &Response{
		Payload: token,
		Header: header,
	},nil

}