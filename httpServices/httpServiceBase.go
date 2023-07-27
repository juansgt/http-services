package httpServices

import (
	"errors"
	"net/http"
	"strings"
)

type HttpServiceBase struct {
}

func NewHttpServiceBase() *HttpServiceBase {
	return &HttpServiceBase{}
}

func (hsg *HttpServiceBase) GetAccessToken(request *http.Request) (string, error) {
	authHeader := request.Header.Get("Authorization")

	// Check that the Authorization header is not empty
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	// Check that the Authorization header uses the Bearer authentication scheme
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		return "", errors.New("invalid Authorization header")
	}

	return authHeaderParts[1], nil
}
