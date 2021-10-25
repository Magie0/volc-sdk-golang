package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type UnfollowResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *UnfollowResult
}

type UnfollowResult struct {
	Create bool
	User *UserInfo
}