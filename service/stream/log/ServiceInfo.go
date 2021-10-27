package log_s

import (
	"net/http"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

func NewServiceInfo() (info base.ServiceInfo) {
	info.Timeout = 5*time.Second
	info.Host = "open.volcengineapi.com"
	header := make(http.Header)
	header.Add("Accept", "application/json")
	header.Add("x-tt-env", "ppe_zhangyuhang_debug")
	header.Add("x-use-ppe", "1")
	info.Header = header
	
	info.Credentials.Region = base.REGION_CN_NORTH_1
	info.Credentials.Service = "content"

	info.Scheme = "https"
	return
}