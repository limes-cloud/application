package md

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"

	"github.com/limes-cloud/user-center/api/auth"
	"github.com/limes-cloud/user-center/api/errors"
)

func New(uid, aid, cid uint32, appKey string) map[string]any {
	data := auth.Auth{
		UserID:     uid,
		AppID:      aid,
		ChannelID:  cid,
		AppKeyword: appKey,
	}

	m := map[string]any{}
	_ = util.Transform(data, &m)

	return m
}

func Get(ctx kratosx.Context) (*auth.Auth, error) {
	data := auth.Auth{}
	if err := ctx.JWT().Parse(ctx.Ctx(), &data); err != nil {
		return nil, err
	}
	if data.UserID == 0 {
		return nil, errors.Forbidden()
	}
	return &data, nil
}

func UserID(ctx kratosx.Context) uint32 {
	m, err := Get(ctx)
	if err != nil {
		return 0
	}
	return m.UserID
}

func AppID(ctx kratosx.Context) uint32 {
	m, err := Get(ctx)
	if err != nil {
		return 0
	}
	return m.AppID
}
