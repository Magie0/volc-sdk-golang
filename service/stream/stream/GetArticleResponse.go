package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type GetArticleResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *GetArticleResult
}

type GetArticleResult struct {
	GroupId string
	VidoeId string
	Tag string
	Title string
	ArticleUrl string
	ArticleClass string
	PublishTime int64
	Abstracts string
	ShareUrl string
	ShareCount int64
	UserInfo *UserInfo
	HasVideo bool
	WatchCount int64
	Duration int64
	Label string
	BuryCount int64
	CommentCount int64
	CommentUrl string
	CoverMode int64
	CoverImageList []Image
	BanComment int64
	DetailSource string
	DiggCount int64
	DiggStatus bool
	ImageList []Image
	ItemIdStr string
	MiddleImageList []Image
	NewVersionVideoPage bool
	Author string
	HomePage string
}