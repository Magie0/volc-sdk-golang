package stream

type FeedBackRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Ouid string
	GroupId string
	ContentType string
	ReportFrom string
	Source string
	ReportType string
}