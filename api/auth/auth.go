package auth

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/user-center/types"
)

func Get(ctx kratosx.Context) (*types.Auth, error) {
	data := types.Auth{}
	return &data, ctx.Authentication().ParseAuth(ctx, &data)
}
