package stream

type FollowArticlesRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Ouid string
	Count int
	Offset int
	Source string
	StartCursor string
	EndCursor string
}