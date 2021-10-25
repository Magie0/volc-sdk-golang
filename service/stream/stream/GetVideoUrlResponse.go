package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type GetVideoUrlResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *GetVideoUrlResult
}

type GetVideoUrlResult struct {
	Status int
	VideoDuration float64
	PosterUrl string
	VideoId string
	VideoList []VideoList
}

type VideoList struct {
	MainUrl string
	MainHttpUrl string
	BackupUrl1 string
	BackupHttpUrl string
	UrlExpire int64
	FileId string
	Bitrate int
	CodecType string
	LogoType string
	Size int
	Fps int
	Quality string
	Encrypt bool
	FileHash string
	VHeight int
	VWidth int
	VType string
	Definition string
	Redirect bool
}