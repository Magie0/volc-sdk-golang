package stream_s

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/stream/stream"
)

func (s *StreamService) UnDigg(diggRequest stream.DiggRequest) (resp stream.DiggResponse,err error) {
	respBody, statusCode, err := s.Client.Query(base.UnDigg,base.ToUrlValues(&diggRequest))
	if err != nil || statusCode != 200 { 
		return resp, err
	}

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, err
	}
	fmt.Print("-----"+string(respBody))
	return resp, nil
}