package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type RecommendRelatedUserResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result           *RecommendRelatedUserResult
}

type RecommendRelatedUserResult struct {
	ChannelId       int
	Description     string
	Intro           int
	RecommendReason string
	User            *UserInfo
	FilterWords     []*FilterWords
}
