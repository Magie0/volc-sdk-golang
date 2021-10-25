package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type RecommendCategoryUserResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *RecommendCategoryUserResult
}

type RecommendCategoryUserResult struct {
	HasMore bool
	Users []*RecommendRelatedUserResult
}