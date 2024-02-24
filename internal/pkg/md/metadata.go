package md

import (
	"errors"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"

	"github.com/limes-cloud/user-center/types"
)

func New(uid, aid, cid uint32, appKey string) map[string]any {
	data := types.Auth{
		UserID:     uid,
		AppID:      aid,
		ChannelID:  cid,
		AppKeyword: appKey,
	}

	m := map[string]any{}
	_ = util.Transform(data, &m)

	return m
}

func Get(ctx kratosx.Context) (*types.Auth, error) {
	data := types.Auth{}
	m, err := ctx.JWT().ParseMapClaims(ctx.Ctx())
	if err = util.Transform(m, &data); err != nil {
		return nil, err
	}
	if data.UserID == 0 {
		return nil, errors.New("data error")
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
