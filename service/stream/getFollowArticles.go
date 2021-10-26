package stream_s

import (
	"encoding/json"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/stream/stream"
)

func (s *StreamService) GetFollowArticles(followArticlesRequest stream.FollowArticlesRequest) (resp stream.FollowArticlesResponse, err error) {
	respBody, statusCode, err := s.Client.Query(base.FollowArticles, base.ToUrlValues(&followArticlesRequest))
	if err != nil || statusCode != 200 {
		return resp, err
	}

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, err
	}
	//fmt.Print("-----" + string(respBody))
	return resp, nil
}
