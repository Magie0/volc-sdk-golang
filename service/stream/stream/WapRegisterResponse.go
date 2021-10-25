package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type WapRegisterResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *WapRegisterResult
}

type WapRegisterResult struct {
	AccessToken string
	ExpiresIn int
}