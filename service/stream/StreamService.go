package stream_s

import "github.com/volcengine/volc-sdk-golang/base"

type StreamService struct {
	client *base.Client
}

func NewStreamService() (s *StreamService) {
	ServiceInfo := NewServiceInfo()
	ApiInfoList := NewApiInfoList()
	s.client = base.NewClient(ServiceInfo, *ApiInfoList)
	return
}
