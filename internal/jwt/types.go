package jwt

import (
	"context"
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Response struct {
	Raw               string                 `json:"raw"`
	Header            map[string]interface{} `json:"header"`
	Payload           map[string]interface{} `json:"payload"`
	Valid             *bool                  `json:"valid,omitempty"`
	VerifiedSignature *bool                  `json:"verifiedSignature,omitempty"`
}

func NewResponse(rawToken string, keys jwk.Set) (*Response, error) {
	var (
		payload           map[string]interface{}
		header            map[string]interface{}
		err               error
		msg               *jws.Message
		token             jwt.Token
		valid             *bool
		verifiedSignature *bool
	)

	// payload
	if token, err = jwt.ParseString(rawToken, jwt.WithVerify(false), jwt.WithValidate(false)); err != nil {
		return nil, fmt.Errorf("failed to parse payload: %s\n", err)
	}
	if payload, err = token.AsMap(context.TODO()); err != nil {
		return nil, err
	}

	// header
	if msg, err = jws.ParseString(rawToken); err != nil {
		return nil, fmt.Errorf("failed to parse token data: %v", err)
	}
	if header, err = msg.Signatures()[0].ProtectedHeaders().AsMap(context.TODO()); err != nil {
		return nil, fmt.Errorf(`failed to parse token data: %v`, err)
	}

	if err = jwt.Validate(token); err != nil {
		valid = newOptionalBool(false)
	} else {
		valid = newOptionalBool(true)
	}

	// verify
	if keys != nil {
		if _, err = jws.Verify([]byte(rawToken), jws.WithKeySet(keys)); err != nil {
			verifiedSignature = newOptionalBool(false)
		} else {
			verifiedSignature = newOptionalBool(true)
		}
	}

	return &Response{
		Raw:               rawToken,
		Payload:           payload,
		Header:            header,
		Valid:             valid,
		VerifiedSignature: verifiedSignature,
	}, nil
}

func newOptionalBool(b bool) *bool {
	return &b
}
