package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type RefreshTipsResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *RefreshTipsResult
}

type RefreshTipsResult struct {
	Count int
	Tip string
}