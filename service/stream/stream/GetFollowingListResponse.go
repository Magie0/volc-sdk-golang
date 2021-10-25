package stream

import "github.com/volcengine/volc-sdk-golang/service/stream/response"


type GetFollowingListResponse struct {
	ResponseMetadata response.ResponseMetadata
	Result *Result
}

type Result struct {
	HasMore bool
	User []GetFollowingListUserInfo
	Cursor int
	Total int
}

type GetFollowingListUserInfo struct {
	FollowTime int64
	Info *Info
	Relation *Relation
	RelationCount *RelationCount
	UpdateArticle int64
	RecentVisitTime int64
}

type Info struct {
	AvatarUrl string
	BanStatus bool
	Desc string
	Name string
	Schema string
	UserAuthInfo string
	UserId string
	UserVerified bool
	VerifiedContent string
}

type Relation struct {
	IsFollowed int
	IsFollowing int
	IsFriend int
}

type RelationCount struct {
	FollowersCount int
	FollowingsCount int
}