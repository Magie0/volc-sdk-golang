package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type FollowResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *FollowResult
}

type FollowResult struct {
	Create bool
	User *UserInfo
	IsFollowing bool
	IsFollowed bool
	UserVerified bool
	Name string
	UserId string
	ScreenName string
	AvatarUrl string
	Type int
	Description string
}