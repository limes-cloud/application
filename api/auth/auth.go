package auth

import (
	"github.com/limes-cloud/kratosx"
)

type Auth struct {
	UserID     uint32 `json:"user_id"`
	AppID      uint32 `json:"app_id"`
	AppKeyword string `json:"app_keyword"`
	ChannelID  uint32 `json:"channel_id"`
}

func Get(ctx kratosx.Context) (*Auth, error) {
	data := Auth{}
	return &data, ctx.Authentication().ParseAuth(ctx, &data)
}
