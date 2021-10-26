package stream_s

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type StreamService struct {
	Client *base.Client
}

func NewStreamService() (s StreamService) {
	ServiceInfo := NewServiceInfo()
	ApiInfoList := NewApiInfoList()
	s.Client = new(base.Client)
	s.Client = base.NewClient(&ServiceInfo, ApiInfoList)
	return
}
