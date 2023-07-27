package httpServices

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type HttpServiceGin struct {
}

func NewHttpServiceGin() *HttpServiceGin {
	return &HttpServiceGin{}
}

func (hsg *HttpServiceGin) GetAccessToken(context *gin.Context) (string, error) {
	authHeader := context.Request.Header.Get("Authorization")

	// Check that the Authorization header is not empty
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	// Check that the Authorization header uses the Bearer authentication scheme
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		context.JSON(401, gin.H{"error": "Invalid Authorization header"})
	}

	return authHeaderParts[1], nil
}
