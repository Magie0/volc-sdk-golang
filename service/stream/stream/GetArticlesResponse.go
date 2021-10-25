package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type GetArticlesResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *GetArticlesResult
}

type GetArticlesResult struct {
	ArticleInfos []*GetArticlesResult
}