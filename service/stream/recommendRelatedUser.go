package stream_s

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/stream/stream"
)

func (s *StreamService) RecommendRelatedUser(recommendRelatedUserRequest stream.RecommendRelatedUserRequest) (resp stream.RecommendRelatedUserResponse, err error) {
	respBody, statusCode, err := s.Client.Query(base.RecommendRelatedUser, base.ToUrlValues(&recommendRelatedUserRequest))
	if err != nil || statusCode != 200 {
		return resp, err
	}

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, err
	}
	fmt.Print("-----" + string(respBody))
	return resp, nil
}
