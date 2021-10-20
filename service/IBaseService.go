package service

import (
	"net/http"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
)

type IBaseService interface {
	/**
     * Sets client no reuse.
     */
	 setClientNoReuse()

	 /**
	  * Gets access key.
	  *
	  * @return the access key
	  */
	getAccessKey() string
 
	 /**
	  * Sets access key.
	  *
	  * @param accessKey the access key
	  */
	setAccessKey(accessKey string)
 
	 /**
	  * Gets secret key.
	  *
	  * @return the secret key
	  */
	getSecretKey() string
 
	 /**
	  * Sets secret key.
	  *
	  * @param secretKey the secret key
	  */
	setSecretKey(secretKey string)
 
	 /**
	  * Sets region.
	  *
	  * @param region the region
	  */
	setRegion(region string)
 
	 /**
	  * Gets region.
	  *
	  * @return the region
	  */
	getRegion() string
 
	 /**
	  * Sets host.
	  *
	  * @param host the host
	  */
	setHost(host string)
 
	 /**
	  * Sets scheme.
	  *
	  * @param scheme the scheme
	  */
	setScheme(scheme string)
 
	 /**
	  * Sets http client.
	  *
	  * @param httpClient the http client
	  */
	setHttpClient(httpClient http.Client)
 
	 /**
	  * Sets service info.
	  *
	  * @param serviceInfo the service info
	  */
	setServiceInfo(serviceInfo base.ServiceInfo)
 
	 /**
	  * Sets socket timeout.
	  *
	  * @param socketTimeout the socket timeout
	  */
	setSocketTimeout(socketTimeout int)
 
	 /**
	  * Sets connection timeout.
	  *
	  * @param connectionTimeout the connection timeout
	  */
	setConnectionTimeout(connectionTimeout int)
 
	 /**
	  * Query raw response.
	  *
	  * @param api    the api
	  * @param params the params
	  * @return the raw response
	  * @throws Exception the exception
	  */
	query(api string,params url.Values) base.RawResponse
	
	 /**
	  * Gets sign url.
	  *
	  * @param api    the api
	  * @param params the params
	  * @return the sign url
	  * @throws Exception the exception
	  */
	getSignUrl(api string, params url.Values) string
 
	 /**
	  * Json raw response.
	  *
	  * @param api    the api
	  * @param params the params
	  * @param body   the body
	  * @return the raw response
	  * @throws Exception the exception
	  */
	json(api string, params url.Values, body string) base.RawResponse
 
	 /**
	  * Post raw response.
	  *
	  * @param api    the api
	  * @param params the params
	  * @param form   the form
	  * @return the raw response
	  * @throws Exception the exception
	  */
	post(api string, params url.Values, form url.Values) base.RawResponse
 
	 /**
	  * Put boolean.
	  *
	  * @param url      the url
	  * @param filePath the file path
	  * @param headers  the headers
	  * @return the boolean
	  * @throws Exception the exception
	  */
	put(url string, filePath string, headers map[string]string) bool
 
	 /**
	  * Put binary data.
	  *
	  * @param url target url
	  * @param data binary data
	  * @param headers http headers
	  * @return put status
	  * @throws Exception exception
	  */
	putData(url string, data []byte, headers map[string]string) bool
 
 
	 /**
	  * Sign by sts2.
	  * @param inlinePolicy  the Policy
	  * @param expire        expire time
	  * @return  the sts2
	  * @throws Exception    the exception
	  */
	signSts2(inlinePolicy base.Policy, expire int64) base.SecurityToken2

}