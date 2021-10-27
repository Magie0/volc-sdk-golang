package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type RelatedArticleResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result []*RelatedArticleResult
}

type RelatedArticleResult struct {
	GroupId string
	VideoId string
	Tag string
	Title string
	ArticleUrl string
	PublishTime int64
	Abstracts string
	ArticleClass string
	ShareUrl string
	UserInfo *RelatedArticleUserInfo
	HasVideo bool
	BuryCount int64
	CommentUrl string
	BanComment int64
	DiggCount int64
	ImageList []Image
	MiddleImageList []Image
	BanAction bool
	BanBury bool
	FilterWords []string
	Source string
	VideoWatchCount int64
	VideoDuration int64
	Url string
}

type RelatedArticleUserInfo struct {
	AvatarUrl string
	Description string
	FollowerCount int64
	UserVerified bool
	VerifiedContent string
	HomePage string
	Name string
	UserId int
	//UserId string
}