package md

import (
	"github.com/limes-cloud/kratosx"
)

const (
	_uid  = "user_id"
	_aid  = "app_id"
	_akey = "app_keyword"
	_cid  = "channel_id"
)

func New(uid, aid, cid uint32, appKey string) map[string]any {
	return map[string]any{
		_uid:  uid,
		_aid:  aid,
		_cid:  cid,
		_akey: appKey,
	}
}

func UserID(ctx kratosx.Context) uint32 {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return 0
	}
	return uint32(m[_uid].(float64))
}

func AppID(ctx kratosx.Context) uint32 {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return 0
	}
	return uint32(m[_aid].(float64))
}

func AppKeyword(ctx kratosx.Context) string {
	m, err := ctx.JWT().ParseMapClaims(ctx)
	if err != nil {
		return ""
	}
	return m[_akey].(string)
}
