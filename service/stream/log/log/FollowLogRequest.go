package log

type FollowLogRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Ouid string
	Body FollowLogRequestBody
}

type FollowLogRequestBody struct {
	ToUserId string
	GroupId string
	ProfileUserId string
	CategoryName string
	Source string
}