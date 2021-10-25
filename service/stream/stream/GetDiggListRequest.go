package stream

type GetDiggListRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Ouid string
	Offset int
	Limit int
}