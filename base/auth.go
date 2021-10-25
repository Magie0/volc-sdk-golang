package base

import (
	"time"

	signablerequest "github.com/volcengine/volc-sdk-golang/service/SignableRequest"
	"github.com/bits-and-blooms/bitset"
)

type ISignerV4 interface {
	Sign(request signablerequest.SignableRequest,credentials Credentials) (err error)

	SignUrl(request signablerequest.SignableRequest,credentials Credentials) (s string,err error)
}

type SignerV4Impl struct {
	tz 			time.Time
	H_INCLUDE 	map[string]bool
	URLENCODER	bitset.BitSet
	CONST_ENCODE string
}

func NewSignerV4Impl() (s *SignerV4Impl) {
	//s.tz.
	s.CONST_ENCODE = "0123456789ABCDEF"
	
	s.H_INCLUDE["Content-Type"] = true
	s.H_INCLUDE["Content-Md5"] = true
	s.H_INCLUDE["Host"] = true

	var i uint
	for i = 97;i <= 122;i=i+1 {
		s.URLENCODER.Set(i)
	}
	for i = 65;i <= 90;i=i+1 {
		s.URLENCODER.Set(i)
	}
	for i = 48;i <= 57;i=i+1 {
		s.URLENCODER.Set(i)
	}

	s.URLENCODER.Set('-')
	s.URLENCODER.Set('-')
	s.URLENCODER.Set('.')
	s.URLENCODER.Set('~')
	return
}
