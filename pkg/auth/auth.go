package auth

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/user-center/pkg/md"
)

func Get(ctx kratosx.Context) (*md.Data, error) {
	data := md.Data{}
	return &data, ctx.Authentication().ParseAuth(ctx, &data)
}
