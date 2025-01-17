package common

import (
	"net/http"
	"os"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	DefaultRegion  = ""
	ServiceVersion = "2021-07-05"
	ServiceName    = "gtm"
	Region         = ""
	Timeout        = 15
)

var (
	ServiceInfo = &base.ServiceInfo{
		Timeout: Timeout * time.Second,
		Host:    "",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Scheme: "http",
		Credentials: base.Credentials{
			Service:         ServiceName,
			Region:          DefaultRegion,
			AccessKeyID:     os.Getenv("VOLC_ACCESSKEY"),
			SecretAccessKey: os.Getenv("VOLC_SECRETKEY"),
		},
	}
)

// SDKClient .
type SDKClient struct {
	Client *base.Client
}

// GetServiceInfo interface
func (p *SDKClient) GetServiceInfo() *base.ServiceInfo {
	return ServiceInfo
}

func (p *SDKClient) SetRegion(region string) {
	p.Client.ServiceInfo.Credentials.Region = region
}

func (p *SDKClient) SetHost(host string) {
	p.Client.ServiceInfo.Host = host
}

func (p *SDKClient) SetSchema(schema string) {
	p.Client.ServiceInfo.Scheme = schema
}
