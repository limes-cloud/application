package md

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/user-center/internal/model"
)

const (
	uid = "user_id"
)

func New(info *model.User) map[string]any {
	return map[string]any{
		uid: info.ID,
	}
}

func UserId(ctx kratosx.Context) uint32 {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return 0
	}
	return uint32(m[uid].(float64))
}
