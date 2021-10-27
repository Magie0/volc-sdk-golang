package log

type ClickLogRequest struct {
	Timestamp int64
	Partner string
	GroupId string
	AccessToken string
	Category string
	EventTime string
	FromGid string
	Dt string
	Os string
	OsVersion string
	ClientVersion string
	DeviceBrand string
}