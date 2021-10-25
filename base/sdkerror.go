package base

type SdkError int

const (
	UNKNOWN		SdkError = -1
    SUCCESS 	SdkError = 0
    ENOAPI 	    SdkError = 10001
    ENOFILE 	SdkError = 10002
    ERESP       SdkError = 10003
    ESIGN	    SdkError = 10004
    EHTTP 	    SdkError = 10005
    EINTERNAL 	SdkError = 10006
    EENCODING 	SdkError = 10007
    EUPLOAD 	SdkError = 10008
)

func (s SdkError) String() string {
	switch(s) {
	case	UNKNOWN  : return "unkown error"
	case 	SUCCESS  : return "success"
	case 	ENOAPI   : return "no this api"
	case 	ENOFILE  : return "cant find file"
	case 	ERESP    : return "response error"
	case 	ESIGN    : return "sign error"
	case 	EHTTP    : return "http request error"
	case 	EENCODING: return "encoding error"
	case 	EINTERNAL: return "internal error"
	case 	EUPLOAD  : return "upload error"
	default			 : return "unknown error"	
	}
}

func (s SdkError) GetNumber() int {
	return int(s)
}