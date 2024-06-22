package auth

import (
	"github.com/limes-cloud/kratosx"
)

type Auth struct {
	UserId     uint32 `json:"userId"`
	AppKeyword string `json:"appKeyword"`
}

func Get(ctx kratosx.Context) (*Auth, error) {
	data := Auth{}
	return &data, ctx.Authentication().ParseAuth(ctx, &data)
}
