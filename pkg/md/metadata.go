package md

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/user-center/pkg/util"
)

type Data struct {
	UserID     uint32 `json:"user_id"`
	AppID      uint32 `json:"app_id"`
	AppKeyword string `json:"app_keyword"`
	ChannelID  uint32 `json:"channel_id"`
}

func New(uid, aid, cid uint32, appKey string) map[string]any {
	data := Data{
		UserID:     uid,
		AppID:      aid,
		ChannelID:  cid,
		AppKeyword: appKey,
	}

	m := map[string]any{}
	_ = util.Transform(data, &m)

	return m
}

func Get(ctx kratosx.Context) (*Data, error) {
	data := Data{}
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err = util.Transform(m, &data); err != nil {
		return nil, err
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

func AppKeyword(ctx kratosx.Context) string {
	m, err := Get(ctx)
	if err != nil {
		return ""
	}
	return m.AppKeyword
}
