package md

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/usercenter/api/usercenter/auth"
	"github.com/limes-cloud/usercenter/api/usercenter/errors"
)

func New(in auth.Auth) map[string]any {
	m := map[string]any{}
	_ = valx.Transform(in, &m)
	return m
}

func Get(ctx kratosx.Context) (*auth.Auth, error) {
	data := auth.Auth{}
	if err := ctx.JWT().Parse(ctx.Ctx(), &data); err != nil {
		return nil, err
	}
	if data.UserId == 0 {
		return nil, errors.ForbiddenError()
	}
	return &data, nil
}

func UserID(ctx kratosx.Context) uint32 {
	m, err := Get(ctx)
	if err != nil {
		ctx.Logger().Warnw("msg", "get user_id error", "err", err.Error())
		return 0
	}
	return m.UserId
}

func AppKey(ctx kratosx.Context) string {
	m, err := Get(ctx)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app_keyword error", "err", err.Error())
		return ""
	}
	return m.AppKeyword
}
