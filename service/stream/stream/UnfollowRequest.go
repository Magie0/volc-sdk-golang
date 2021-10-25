package stream

type UnfollowRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Ouid string
	UserId string
	FromGid string
}