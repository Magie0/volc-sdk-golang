package stream

type RecommendCategoryUserRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Ouid string
	ChannelId int
	Offset int
	Count int
}