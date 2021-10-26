package stream_s

import (
	"encoding/json"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/stream/stream"
)

func (s *StreamService) CheckRelation(checkRelationRequst stream.CheckRelationRequst) (resp stream.CheckRelationResponse, err error) {
	respBody, statusCode, err := s.Client.Query(base.CheckRelation, base.ToUrlValues(&checkRelationRequst))
	if err != nil || statusCode != 200 {
		return resp, err
	}

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, err
	}
	//fmt.Print("-----" + string(respBody))
	return resp, nil
}
