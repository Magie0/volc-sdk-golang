package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type GetListResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result []*GetListResult
}

type GetListResult struct {
	GroupId string
	VidoeId string
	Tag string
	Title string
	Source string
	ArticleUrl string
	ArticleClass string
	PublishTime int64
	Abstracts string
	ShareUrl string
	ShareCount int64
	HasVideo bool
	VideoWatchCount int64
	VideoDuration int64
	Label string
	DiggCount int64
	BuryCount int64
	CommentCount int64
	CommentUrl string
	CoverMode int64
	CoverImageList []Image
	FilterWords []FilterWords
	IsStick bool
	ProfileTags []string
	UserInfo *UserInfo
	MiddleImage *Image
	LargeImageList []Image
	ArticleType int
	CellType int
	ImageList []Image
}

type FilterWords struct {
	Name string
	Id string
}