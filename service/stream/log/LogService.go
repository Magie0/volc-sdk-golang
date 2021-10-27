package log_s

import (
	"github.com/volcengine/volc-sdk-golang/base"
)

type LogService struct {
	Client *base.Client
}

func NewStreamService() (s LogService) {
	ServiceInfo := NewServiceInfo()
	ApiInfoList := NewApiInfoList()
	s.Client = new(base.Client)
	s.Client = base.NewClient(&ServiceInfo, ApiInfoList)
	return
}