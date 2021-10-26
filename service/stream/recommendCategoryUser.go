package stream_s

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/stream/stream"
)

func (s *StreamService) RecommendCategoryUser(recommendCategoryUserRequest stream.RecommendCategoryUserRequest) (resp stream.RecommendCategoryUserResponse, err error) {
	respBody, statusCode, err := s.Client.Query(base.RecommendCategoryUser, base.ToUrlValues(&recommendCategoryUserRequest))
	if err != nil || statusCode != 200 {
		return resp, err
	}

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, err
	}
	fmt.Print("-----" + string(respBody))
	return resp, nil
}