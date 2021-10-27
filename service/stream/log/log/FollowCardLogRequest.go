package log

type FollowCardLogRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Ouid string
	Body *FollowCardLogRequestBody
}

type FollowCardLogRequestBody struct {
	ToUserId string
	FollowType string
	ProfileUserId string
	CategoryName string
	Source string
	ActionType string
	Order int
}