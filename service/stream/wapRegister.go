package stream_s

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/stream/stream"
)

func (s *StreamService) WapRegister(wapRegisterRequest stream.WapRegisterRequest) (resp stream.WapRegisterResponse,err error) {
	values := base.ToUrlValues(&wapRegisterRequest)
	respBody, statusCode, err := s.Client.Query(base.WapRegister,values)
	if err != nil || statusCode != 200 { 
		return resp, err
	}
	//println(string(respBody))
	if err := json.Unmarshal(respBody, &resp); err != nil {
		fmt.Print(err)
		return resp, err
	}
	//fmt.Println([]rune("-----"+string(respBody)))
	return resp, nil
}
