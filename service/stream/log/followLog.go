package log_s

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/stream/log/log"
)

func (s *LogService) FollowLog(followLogRequest log.FollowLogRequest) (resp log.LogResponse, err error) {
	respBody, statusCode, err := s.Client.Query(base.FollowLog, base.ToUrlValues(&followLogRequest))
	if err != nil || statusCode != 200 {
		return resp, err
	}

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, err
	}
	fmt.Print("-----" + string(respBody))
	return resp, nil
}