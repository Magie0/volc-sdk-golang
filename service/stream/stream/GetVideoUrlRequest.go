package stream

type GetVideoUrlRequest struct {
	Partner string
	AccessToken string
	Timestamp int64
	Nonce string
	CodecType string
	Scene string
	VideoId string
}