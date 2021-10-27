package log

type MultiShowLogRequest struct {
	Timestamp int64
	Partner string
	AccessToken string
	Body *MultiShowLogRequestBody
}

type MultiShowLogRequestBody struct {
	GroupId string
	Category string
	EventTime int64
	FromGid string
	Dt string
	Os string
	OsVersion string
	ClientVersion string
	DeviceBrand string
}