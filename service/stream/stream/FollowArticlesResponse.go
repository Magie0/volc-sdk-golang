package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type FollowArticlesResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *FollowArticlesResult
}

type FollowArticlesResult struct {
	HasMore bool
	Articles []ArticleInfos
}

type ArticleInfos struct {
	ArticleUrl string
	CommentCount int64
	CommentUrl string
	CreatTime int64
	DiggCount int64
	FilterWords  []FilterWord
	GroupId string
	GroupSource int
	HasVideo bool
	ImageList []Image
	ItemId string
	LargeImageList []Image
	MiddleImageList []Image
	PublishTime int64
	PushId int64
	ReportUrl string
	ShareCount int64
	ShareUrl string
	Source string
	Summary string
	Title string
	UserId int64
	UserInfo *UserInfo
}

type FilterWord struct {
	Id string
	Name string
	IsSelect bool
}