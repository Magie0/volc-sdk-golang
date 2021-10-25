package response

type ResponseMetadata struct {
	RequestId string
	Action string
	Version string
	Service string
	Region string
	Error *Error
}

type Error struct {
	CodeN int64
	Code string
	Message string
}