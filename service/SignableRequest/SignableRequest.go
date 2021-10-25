package signablerequest

import (
	"net/http"
)

type SignableRequest struct {
	Method string
	Request http.Request
}