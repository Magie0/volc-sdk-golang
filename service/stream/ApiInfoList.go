package stream_s

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
	a.putBury()
	a.putCheckRelation()
	a.putContentStream()
	a.putDigg()
	a.putDiggList()
	a.putFeedback()
	a.putFollow()
	a.putFollowArticles()
	a.putFollowingList()
	a.putMultiArticleInfo()
	a.putRecommendCategoryInfo()
	a.putRecommendCategoryUser()
	a.putRecommendRelatedUser()
	a.putRefreshTips()
	a.putRelatedArticle()
	a.putSingleArticleInfo()
	a.putUnBury()
	a.putUnDigg()
	a.putUnfollow()
	a.putVideoUrl()
	a.putWapRegister()
	return a.api
}

func (a apiInfolist) putWapRegister() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.WapRegister)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.WapRegister] = &apiInfo
}

func (a apiInfolist) putContentStream() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.ContentStream)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.ContentStream] = &apiInfo
}

func (a apiInfolist) putDigg() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.Digg)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.Digg] = &apiInfo
}

func (a apiInfolist) putBury() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.Bury)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.Bury] = &apiInfo
}

func (a apiInfolist) putUnDigg() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.UnDigg)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.UnDigg] = &apiInfo
}

func (a apiInfolist) putUnBury() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.UnBury)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.UnBury] = &apiInfo
}

func (a apiInfolist) putSingleArticleInfo() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.SingleArticleInfo)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.SingleArticleInfo] = &apiInfo
}

func (a apiInfolist) putMultiArticleInfo() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.MultiArticleInfo)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.MultiArticleInfo] = &apiInfo
}

func (a apiInfolist) putFeedback() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.Feedback)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.Feedback] = &apiInfo
}

func (a apiInfolist) putFollow() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.Follow)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.Follow] = &apiInfo
}

func (a apiInfolist) putUnfollow() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.Unfollow)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.Unfollow] = &apiInfo
}

func (a apiInfolist) putCheckRelation() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.CheckRelation)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.CheckRelation] = &apiInfo
}

func (a apiInfolist) putFollowingList() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.FollowingList)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.FollowingList] = &apiInfo
}

func (a apiInfolist) putRefreshTips() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.RefreshTips)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.RefreshTips] = &apiInfo
}

func (a apiInfolist) putFollowArticles() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.FollowArticles)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.FollowArticles] = &apiInfo
}

func (a apiInfolist) putRecommendRelatedUser() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.RecommendRelatedUser)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.RecommendRelatedUser] = &apiInfo
}

func (a apiInfolist) putRecommendCategoryUser() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.RecommendCategoryUser)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.RecommendCategoryUser] = &apiInfo
}

func (a apiInfolist) putRecommendCategoryInfo() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.RecommendCategoryInfo)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.RecommendCategoryInfo] = &apiInfo
}

func (a apiInfolist) putRelatedArticle() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.RelatedArticle)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.RelatedArticle] = &apiInfo
}

func (a apiInfolist) putDiggList() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.DiggList)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.DiggList] = &apiInfo
}

func (a apiInfolist) putVideoUrl() {
	var apiInfo base.ApiInfo
	apiInfo.Query = make(url.Values)
	apiInfo.Method = "GET"
	apiInfo.Path = "/"
	apiInfo.Query.Add("Action", base.VideoUrl)
	apiInfo.Query.Add("Version", base.ContentVersion)
	a.api[base.VideoUrl] = &apiInfo
}