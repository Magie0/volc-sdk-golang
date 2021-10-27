package log_s

import (
	"net/url"

	"github.com/volcengine/volc-sdk-golang/base"
)

type apiInfolist struct {
	api map[string]*base.ApiInfo
}

func NewApiInfoList() map[string]*base.ApiInfo {
	var a apiInfolist
	a.api = make(map[string]*base.ApiInfo)
	a.putClickLog()
	a.putDislikeLog()
	a.putFavouriteLog()
	a.putFollowCardLog()
	a.putFollowLog()
	a.putMultiShowLog()
	a.putShareLog()
	a.putSingleShowLog()
	a.putStayLog()
	a.putUnfollowLog()
	a.putVerifyLog()
	a.putVideoOverLog()
	a.putVideoPlayLog()
	return a.api
}

func (a apiInfolist) putClickLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.ClickLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.ClickLog] = &apiInfo
}

func (a apiInfolist) putVideoPlayLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.VideoPlayLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.VideoPlayLog] = &apiInfo
}

func (a apiInfolist) putVideoOverLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.VideoOverLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.VideoOverLog] = &apiInfo
}

func (a apiInfolist) putStayLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.StayLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.StayLog] = &apiInfo
}

func (a apiInfolist) putSingleShowLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.SingleShowLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.SingleShowLog] = &apiInfo
}

func (a apiInfolist) putMultiShowLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.MultiShowLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.MultiShowLog] = &apiInfo
}

func (a apiInfolist) putShareLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.ShareLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.ShareLog] = &apiInfo
}

func (a apiInfolist) putFavouriteLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.FavouriteLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.FavouriteLog] = &apiInfo
}

func (a apiInfolist) putVerifyLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.VerifyLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.VerifyLog] = &apiInfo
}

func (a apiInfolist) putFollowLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.FollowLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.FollowLog] = &apiInfo
}

func (a apiInfolist) putUnfollowLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.UnfollowLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.UnfollowLog] = &apiInfo
}

func (a apiInfolist) putFollowCardLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.FollowCardLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.FollowCardLog] = &apiInfo
}

func (a apiInfolist) putDislikeLog() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.DislikeLog)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.DislikeLog] = &apiInfo
}