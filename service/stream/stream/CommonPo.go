package stream

/*type CommonPo struct {
	UserInfo *UserInfo
	Image *Image
	Url *Url
}*/

type UserInfo struct {
	AvatarUrl string
	Description string
	Follow bool
	FollowerCount int64
	UserVerified bool
	VerifiedContent string
	HomePage string
	Name string
	UserId int
	//UserId string
	CreateTime int64
	IsFollowed bool
	IsFollowing bool
	LastUpdate string
	MediaId int64
	ScreenName string
	Type int
	UserAuthInfo string
	Desc string
}

type Image struct {
	Url string
	Uri string
	Width int64
	Height int64
	UrlList []Url
	Mimetype string
}

type Url struct {
	Url string
}