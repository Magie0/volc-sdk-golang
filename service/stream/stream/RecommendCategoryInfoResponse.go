package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type RecommendCategoryInfoResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result []RecommendCategoryInfoResult
}

type RecommendCategoryInfoResult struct {
	Channel string
	ChannelId int64
}