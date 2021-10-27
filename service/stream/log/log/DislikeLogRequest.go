package log

type DislikeLogRequest struct {
	Timestamp int64
	Partner string
	Ouid string
	AccessToken string
	Body []*DislikeLogRequestBody
}

type DislikeLogRequestBody struct {
	Category string
	GroupId string
	Timestamp int64
	FilterWords []string
}