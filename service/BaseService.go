package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/sqs/goreturns/returns"
	"github.com/volcengine/volc-sdk-golang/base"
	signablerequest "github.com/volcengine/volc-sdk-golang/service/SignableRequest"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/response"
)

type BaseService struct {
	VERSION string
	ServiceInfo base.ServiceInfo
	ApiInfoList map[string]base.ApiInfo
	HttpClient *base.Client
	Credential base.Credentials
	Timeout time.Duration
}

func NewBaseService1(info base.ServiceInfo,proxy string,apiInfoList map[string]base.ApiInfo) (b *BaseService) {
	b = new(BaseService)
	b.ServiceInfo = info
	b.ApiInfoList = apiInfoList
	b.Credential = *new(base.Credentials)

	b.HttpClient = new(base.Client)
	b.HttpClient.ServiceInfo.Host = proxy

	b.Init(info)
	return
}

func NewBaseService2(info base.ServiceInfo,apiInfoList map[string]base.ApiInfo) (b *BaseService) {
	b = new(BaseService)
	b.ServiceInfo = info
	b.ApiInfoList = apiInfoList
	b.Credential = *new(base.Credentials)
	
	b.HttpClient = new(base.Client)

	b.Init(info)
	return
}

func (b *BaseService) Init(info base.ServiceInfo) {
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

func (b *BaseService) GetSignUrl(api string,params url.Values) (s string,err error) {
	apiInfo,ok := b.ApiInfoList[api]

	if !ok {
		err = fmt.Errorf("no apiInfo")
		return
	}

	mergedNV := MergeQuery(params,apiInfo.Query)
	var request signablerequest.SignableRequest
	Url := request.Request.URL

	request.Method = strings.ToUpper(request.Method)
	Url.Scheme = b.ServiceInfo.Scheme
	Url.Host = b.ServiceInfo.Host
	Url.Path = apiInfo.Path
	Url.RawQuery = mergedNV.Encode()

	return b.Credential.SignUrl(&request.Request),nil
}

func MergeQuery(query1 url.Values, query2 url.Values) url.Values {
	res := *new(url.Values)
	for k,v := range query1 {
		res[k] = v
	}
	for k,v := range query2 {
		res[k] = v
	}
	return res
}

func (b *BaseService) Put(url string, filePath string, headers map[string]string) (bol bool, err error) {

	payload,err := ioutil.ReadFile(filePath)
	body := strings.NewReader(string(payload))
	req, _ := http.NewRequest("PUT", url, body)
	if headers != nil && len(headers) > 0 {
		for k,v := range headers {
			req.Header.Add(k,v)
		}
	}

	client := *http.DefaultClient
	client.Timeout = b.Timeout
	response,err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("put error")
		return false,err
	}
	if response.StatusCode == 200 {
		return true,nil
	}
	return false,nil
}

func (b *BaseService)PutData(url string, data []byte, headers map[string]string) (bol bool, err error) {

	body := strings.NewReader(string(data))
	req, _ := http.NewRequest("PUT", url, body)
	if headers != nil && len(headers) > 0 {
		for k,v := range headers {
			req.Header.Add(k,v)
		}
	}

	client := *http.DefaultClient
	client.Timeout = b.Timeout
	response,err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("put error")
		return false,err
	}
	if response.StatusCode == 200 {
		return true,nil
	}
	return false,nil
}

func (b *BaseService) PutDataWithResponse(url string,data []byte,headers map[string]string) (response *http.Response) {
	body := strings.NewReader(string(data))
	httpPut, _ := http.NewRequest("PUT", url, body)
	if headers != nil && len(headers) > 0 {
		for k,v := range headers {
			httpPut.Header.Add(k,v)
		}
	}

	client := *http.DefaultClient
	response,_ = client.Do(httpPut)
	return
}

func (b *BaseService) Json(api string, params url.Values, body string) (raw base.RawResponse) {
	_,ok := b.ApiInfoList[api]
	if !ok {
		raw.Data = nil
		raw.Code = base.SdkError.GetNumber(base.ENOAPI)
		raw.Err  = fmt.Errorf(base.ENOAPI.String())
		return
	}

	request := b.PrepareRequest(api,params)
	request.Request.Header.Set("Content-Type", "application/json")
	raw,_ = b.MakeRequest(api,*request)
	return
}

func (b *BaseService) Post(api string, query url.Values, form url.Values) (raw base.RawResponse) {
	apiInfo,ok := b.ApiInfoList[api]
	if !ok {
		raw.Data = nil
		raw.Code = base.SdkError.GetNumber(base.ENOAPI)
		raw.Err  = fmt.Errorf(base.SdkError.String(base.ENOAPI))
		return
	}

	request := b.PrepareRequest(api,query)
	request.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	mergedForm := MergeQuery(form,apiInfo.Form)
	
	request.Request.Body.Read([]byte(mergedForm.Encode()))

	raw,_ = b.MakeRequest(api, *request)
	return
}

func (b *BaseService) query(api string, params url.Values) (raw base.RawResponse) {
	_,ok := b.ApiInfoList[api]
	if !ok {
		raw.Data = nil
		raw.Code = base.SdkError.GetNumber(base.ENOAPI)
		raw.Err = fmt.Errorf(base.SdkError.String(base.ENOAPI))
		return
	}

	request := b.PrepareRequest(api,params)
	raw,_ = b.MakeRequest(api,*request)
	return
}

func (b *BaseService) MakeRequest(api string, request signablerequest.SignableRequest) (raw base.RawResponse,err error) {
	//TODO SSigner
	request.Request = *b.Credential.Sign(&request.Request)

	var client http.Client
	var response *http.Response

	client = *http.DefaultClient
	//response = http.Response{}
	
	response,err = client.Do(&request.Request)

	statusCode := response.StatusCode
	if statusCode >= 300 {
		msg := base.SdkError.String(base.EHTTP)
		bytes,_ := ioutil.ReadAll(response.Body)
		if bytes != nil && len(bytes) > 0 {
			msg = string(bytes)
		}
		raw.Data = nil
		raw.Code = base.SdkError.GetNumber(base.EHTTP)
		raw.Err = fmt.Errorf(msg)
		return
	}
	bytes,_ := ioutil.ReadAll(response.Body)
	raw.Data = bytes
	raw.Code = base.SdkError.GetNumber(base.SUCCESS)
	raw.Err = nil
	return

	//TODO ?
}

func (b *BaseService) PrepareRequest(api string, params url.Values) (request *signablerequest.SignableRequest) {
	apiInfo := b.ApiInfoList[api]

	request.Method = strings.ToUpper(apiInfo.Method)
	request.Request.Header = MergeHeader(b.ServiceInfo.Header,apiInfo.Header)
	request.Request.Header.Add("User-Agent", "volc-sdk-java/v" + b.VERSION)
	mergedNV := MergeQuery(params, apiInfo.Query)

	request.Request.URL.Scheme = b.ServiceInfo.Scheme
	request.Request.URL.Host = b.ServiceInfo.Host
	request.Request.URL.Path = apiInfo.Path
	request.Request.URL.RawQuery = mergedNV.Encode()

	timeout := b.GetTimeout(b.ServiceInfo.Timeout,apiInfo.Timeout)
	b.Timeout = timeout

	return
}

func (b *BaseService) GetTimeout(serviceTimeout time.Duration,apiTimeout time.Duration) time.Duration {
	var timeout time.Duration
	timeout = 5000
	if serviceTimeout != 0 {
		timeout = serviceTimeout
	}
	if apiTimeout != 0 {
		timeout = apiTimeout
	}
	if b.Timeout != 0 {
		timeout = b.Timeout
	}
	return timeout
}

func MergeHeader(header1 http.Header,header2 http.Header) (header http.Header) {
	header = header1
	for k,v := range header2 {
		header.Set(k,v[0])
	}
	return
}

func (b *BaseService) SetClientNoReuse() {
	b.HttpClient = nil
}

func (b *BaseService) GetAccessKey() string {
	return b.ServiceInfo.Credentials.AccessKeyID
}

func (b *BaseService) SetAccessKey(accessKey string) {
	b.ServiceInfo.Credentials.AccessKeyID = accessKey
}

func (b *BaseService) GetSecretKey() string {
	return b.ServiceInfo.Credentials.SecretAccessKey
}

func (b *BaseService) SetSecretKey(secretKey string) {
	b.ServiceInfo.Credentials.SecretAccessKey = secretKey
}

func (b *BaseService) GetRegion() string {
	return b.ServiceInfo.Credentials.Region
}

func (b *BaseService) SetRegion(region string) {
	b.ServiceInfo.Credentials.Region = region
}

func (b *BaseService) SetHost(host string) {
	b.ServiceInfo.Host = host
}

func (b *BaseService) SetScheme(scheme string) {
	b.ServiceInfo.Scheme = scheme
}

func (b *BaseService) GetHttpClient() base.Client {
	return *b.HttpClient
}

func (b *BaseService) SetHttpClient(httpClient base.Client) {
	b.HttpClient = &httpClient
}

func (b *BaseService) GetServiceInfo() base.ServiceInfo {
	return b.ServiceInfo
}

func (b *BaseService) SetServiceInfo(serviceInfo base.ServiceInfo) {
	b.ServiceInfo = serviceInfo
}

func (b *BaseService) GetApiInfoList() map[string]base.ApiInfo {
	return b.ApiInfoList
}

func (b *BaseService) GetCredential() base.Credentials {
	return b.Credential
}

func (b *BaseService) SetTimeout(timeout int) {
	b.Timeout = time.Duration(timeout)
}

func (b *BaseService) SignSts2(inlinePolicy base.Policy,expire int64) (sts2 base.SecurityToken2) {
	sts2.AccessKeyID = base.GenerateAccessKeyId("AKTP")
	sts2.SecretAccessKey = base.GenerateSecretKey()

	sts2.CurrentTime = time.Now().Format("yyyy-MM-dd'T'HH:mm:ssXXX")
	sts2.ExpiredTime = time.Now().Add(time.Duration(expire)).Format("yyyy-MM-dd'T'HH:mm:ssXXX")
	
	innerToken := base.CreateInnerToken(b.ServiceInfo.Credentials,&sts2,&inlinePolicy,time.Now().Add(time.Duration(expire)).Unix()/1000)


}