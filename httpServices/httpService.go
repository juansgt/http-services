package httpServices

import (
	"golang.org/x/net/context"
)

type IHttpService[T context.Context] interface {
	GetAccessToken(value T) (string, error)
}
