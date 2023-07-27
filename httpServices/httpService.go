package httpServices

import "net/http"

type IHttpService interface {
	GetAccessToken(request *http.Request) (string, error)
}
