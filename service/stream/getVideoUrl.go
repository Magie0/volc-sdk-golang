package stream_s

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/stream/stream"
)
TODO
func (s *StreamService) GetVideoUrl(getVideoUrlRequest stream.GetVideoUrlRequest) (resp stream.GetVideoUrlResponse,err error) {
	respBody, statusCode, err := s.Client.Query(base.VideoUrl,base.ToUrlValues(&getVideoUrlRequest))
	if err != nil || statusCode != 200 { 
		return resp, err
	}

	if err := json.Unmarshal(respBody, &resp); err != nil {
		return resp, err
	}
	fmt.Print("-----"+string(respBody))
	return resp, nil
}