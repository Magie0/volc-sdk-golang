package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type GetDiggListResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *GetDiggListResult
}

type GetDiggListResult struct {
	Hasmore bool
	Data []*ArticleInfo
}

type ArticleInfo struct {
	GroupId string
	VideoId string
	Tag string
	Title string
	ArticleUrl string
	ArtcleClass string
	PublishTime int64
	Abstracts string
	ShareUrl string
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
	DetailSource string
	DiggCount int64
	ImageList []Image
	ItemIdStr string
	LargeImageList []Image
	MiddleImageList []Image
	Author string
	CellType int
	GroupSource int
	HomePage string
	MediaId int64
	Timestamp int64
	UserId int64
	Video string
}