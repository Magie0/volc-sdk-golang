package stream

type FollowRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Ouid string
	UserId string
	Source int
	Reason int
	FromGid string
}