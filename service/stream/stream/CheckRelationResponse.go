package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"

type CheckRelationResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *CheckRelationResult
}

type CheckRelationResult struct {
	Followed bool
	Following bool
}