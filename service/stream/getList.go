package stream_s

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/stream/stream"
)

func (s *StreamService) GetList(getListRequest stream.GetListRequest) (resp stream.GetListResponse,err error) {
	respBody, statusCode, err := s.client.Query(base.WapRegister,base.ToUrlValues(getListRequest))
	if err != nil || statusCode != base.SdkError.GetNumber(base.SUCCESS) { 
		return resp, err
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return resp, err
	}
	fmt.Print([]rune("-----"+string(respBody)))
	return resp, nil
}