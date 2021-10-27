package log

import "github.com/volcengine/volc-sdk-golang/service/stream/response"

type LogResponse struct {
	ResponseMetadata *response.ResponseMetadata
	Result string
}