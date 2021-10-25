package stream

type GetFollowingListRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Ouid string
	Count int
	Cursor int
	OrderField string
	OrderType string
}