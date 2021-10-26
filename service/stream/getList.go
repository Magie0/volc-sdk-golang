package stream_s

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/stream/stream"
)

func (s *StreamService) GetList(getListRequest stream.GetListRequest) (resp stream.GetListResponse,err error) {
	respBody, statusCode, err := s.Client.Query(base.ContentStream,base.ToUrlValues(&getListRequest))
	if err != nil || statusCode != 200 { 
		return resp, err
	}

	if err := json.Unmarshal(respBody, &resp); err != nil {
		fmt.Println(err)
		return resp, err
	}
	fmt.Print("-----"+string(respBody))
	return resp, nil
}