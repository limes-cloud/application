package authorizer

import (
	"errors"
	"time"

	"github.com/limes-cloud/kratosx"
)

type Authorizer interface {
	Name() string
	GetAccessToken(kratosx.Context, GetAccessTokenRequest) (*GetAccessTokenReply, error)
	GetAuthInfo(kratosx.Context, GetAuthInfoRequest) (*GetAuthInfoReply, error)
}

type GetAuthInfoRequest struct {
	Token string
	Code  string
}

type GetAuthInfoReply struct {
	AuthId  string
	UnionId *string
}

type GetAccessTokenRequest struct {
	Ak   string
	Sk   string
	Code string
}

type GetAccessTokenReply struct {
	Token  string
	Expire time.Duration
}

type Interface interface {
	GetAuthorizers() map[string]Authorizer
	GetToken(ctx kratosx.Context) (*GetAccessTokenReply, error)
	GetAuthInfo(ctx kratosx.Context, token string) (*GetAuthInfoReply, error)
}

type authorizer struct {
	config *Config
}

type Config struct {
	Ak       string
	Sk       string
	Code     string
	Redirect string
	Platform string
}

var ins = make(map[string]Authorizer)

func register(key string, at Authorizer) {
	ins[key] = at
}

func New(config *Config) Interface {
	return &authorizer{
		config: config,
	}
}

func (ath authorizer) GetAuthorizers() map[string]Authorizer {
	return ins
}

func (c *authorizer) GetAuthorizer(platform string) Authorizer {
	return ins[platform]
}

func (c *authorizer) GetToken(ctx kratosx.Context) (*GetAccessTokenReply, error) {
	atr := c.GetAuthorizer(c.config.Platform)
	if atr == nil {
		return nil, errors.New("not exist channel authorizer")
	}
	return atr.GetAccessToken(ctx, GetAccessTokenRequest{
		Ak:   c.config.Ak,
		Sk:   c.config.Sk,
		Code: c.config.Code,
	})
}

func (c *authorizer) GetAuthInfo(ctx kratosx.Context, token string) (*GetAuthInfoReply, error) {
	atr := c.GetAuthorizer(c.config.Platform)
	if atr == nil {
		return nil, errors.New("not exist channel authorizer")
	}
	return atr.GetAuthInfo(ctx, GetAuthInfoRequest{
		Token: token,
		Code:  c.config.Code,
	})
}
