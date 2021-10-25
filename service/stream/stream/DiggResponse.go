package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type DiggResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *DiggResult
}

type DiggResult struct {
	BuryCount int64
	DiggCount int64
}