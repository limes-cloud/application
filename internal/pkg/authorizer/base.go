package authorizer

import "time"

type Authorizer interface {
	Name() string
	GetAccessToken(req GetAccessTokenRequest) (*GetAccessTokenReply, error)
	GetAuthInfo() (*Info, error)
}

type Info struct {
	AuthId  string
	UnionId *string
}

type GetAccessTokenRequest struct {
	Ak       string
	Sk       string
	Code     string
	Redirect string
}

type GetAccessTokenReply struct {
	Token  string
	Expire time.Duration
}

type Interface interface {
	GetAuthorizer(key string) Authorizer
	GetAuthorizers() map[string]Authorizer
}

type authorizer struct {
}

type Config struct {
	Ak       string
	Sk       string
	Code     string
	Redirect string
}

var ins = make(map[string]Authorizer)

func register(key string, at Authorizer) {
	ins[key] = at
}

func New() Interface {
	return &authorizer{}
}

func (ath authorizer) GetAuthorizer(key string) Authorizer {
	return ins[key]
}

func (ath authorizer) GetAuthorizers() map[string]Authorizer {
	return ins
}
