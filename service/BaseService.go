package service

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

type BaseService struct {
	VERSION string
	ServiceInfo base.ServiceInfo
	ApiInfoList map[string]base.ApiInfo
	HttpClient *base.Client
	//ISigner 
	SocketTimeout int
	ConnectionTimeout int
}

func NewBaseService1(info base.ServiceInfo,proxy string,apiInfoList map[string]base.ApiInfo) (b *BaseService) {
	b = new(BaseService)
	b.ServiceInfo = info
	b.ApiInfoList = apiInfoList
	//b.ISigner = new(SignerV4Impl)

	b.HttpClient = new(base.Client)
	b.HttpClient.ServiceInfo.Host = proxy

	b.init(info)
}

func NewBaseService2(info base.ServiceInfo,apiInfoList map[string]base.ApiInfo) (b *BaseService) {
	b = new(BaseService)
	b.ServiceInfo = info
	b.ApiInfoList = apiInfoList
	//b.ISigner = new(SignerV4Impl)
	
	b.HttpClient = new(base.Client)

	b.init(info)
}

func (b *BaseService) init(info base.ServiceInfo) {
	accessKey := os.Getenv(base.ACCESS_KEY)
	secretKey := os.Getenv(base.SECRET_KEY)

	if accessKey != "" && secretKey != "" {
		info.Credentials.AccessKeyID = accessKey
		info.Credentials.SecretAccessKey = secretKey
	} else {
		//file :=
		return
	}
}

func (b *BaseService) getSignUrl(api string,params url.Values) (err error) {
	apiInfo,ok := b.ApiInfoList[api]

	if !ok {
		err = fmt.Errorf("no apiInfo")
		return
	}

	mergedNV := mergeQuery(params,apiInfo.Query)
	request := *new(SignableRequest)
	Url := request.Url

	request.Method = strings.ToUpper(request.Method)
	Url.Scheme = b.ServiceInfo.Scheme
	Url.Host = b.ServiceInfo.Host
	Url.Path = apiInfo.Path
	value := 


	return
}

func mergeQuery(query1 url.Values, query2 url.Values) url.Values {
	res := *new(url.Values)
	for k,v := range query1 {
		res[k] = v
	}
	for k,v := range query2 {
		res[k] = v
	}
	return res
}