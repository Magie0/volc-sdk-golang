package service

import "net/url"

type SignableRequest struct {
	Method string
	Url url.URL
}