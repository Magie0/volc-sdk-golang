package stream

type RecommendRelatedUserRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Ouid string
	Source string
	FollowUserId string
	GroupId string
}