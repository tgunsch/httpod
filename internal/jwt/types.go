package jwt

import (
	"bytes"
	"context"
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

type Response struct {
	//Method    SigningMethod          // The signing method used or to be used
	Raw           string                 `json:"raw"`
	Header        map[string]interface{} `json:"header"`
	Payload       map[string]interface{} `json:"payload"`
	Valid         *bool                  `json:"valid,omitempty"`
	ValidateError string                 `json:"validateError,omitempty"`
}

func NewResponse(rawToken string, validate bool, keys jwk.Set) (*Response, error) {
	var (
		payload       map[string]interface{}
		header        map[string]interface{}
		err           error
		msg           *jws.Message
		token         jwt.Token
		valid         *bool
		validateError string
	)

	if msg, err = jws.ParseString(rawToken); err != nil {
		return nil, fmt.Errorf("failed to parse token data: %v", err)
	}

	if validate {
		if token, err, validateError = parseValidate(keys, msg); err != nil {
			return nil, err
		}
		if validateError != "" {
			valid = newOptionalBool(false)
		}
	} else {
		if token, err = jwt.ParseReader(bytes.NewReader(msg.Payload())); err != nil {
			// try to parse without validate
			return nil, fmt.Errorf("failed to parse payload: %s\n", err)
		}
	}

	if payload, err = token.AsMap(context.TODO()); err != nil {
		return nil, err
	}
	if header, err = getJWTHeader(msg); err != nil {
		return nil, err
	}

	return &Response{
		Raw:           rawToken,
		Payload:       payload,
		Header:        header,
		Valid:         valid,
		ValidateError: validateError,
	}, nil
}

func parseValidate(keys jwk.Set, msg *jws.Message) (jwt.Token, error, string) {
	var (
		err           error
		token         jwt.Token
		options       []jwt.ParseOption
		validateError string
	)

	options = []jwt.ParseOption{jwt.WithValidate(true)}
	if keys != nil {
		options = append(options, jwt.WithKeySet(keys))
	}

	if token, err = jwt.ParseReader(bytes.NewReader(msg.Payload()), options...); err != nil {
		validateError = err.Error()
		// try to parse without validate
		if token, err = jwt.ParseReader(bytes.NewReader(msg.Payload())); err != nil {
			// try to parse without validate
			return nil, fmt.Errorf("failed to parse payload: %s\n", err), validateError
		}
	}
	return token, nil, validateError
}

func getJWTHeader(msg *jws.Message) (map[string]interface{}, error) {
	hdr, err := msg.Signatures()[0].ProtectedHeaders().AsMap(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`failed to parse token data: %v`, err)
	}
	return hdr, nil
}

func newOptionalBool(b bool) *bool {
	return &b
}
