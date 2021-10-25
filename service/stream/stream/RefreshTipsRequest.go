package stream

type RefreshTipsRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Ouid string
	Count string
	Cursor int
}