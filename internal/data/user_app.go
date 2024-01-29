package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/user-center/internal/biz"
)

type userAppRepo struct {
}

func NewUserAppRepo() biz.UserAppRepo {
	return &userAppRepo{}
}

func (u *userAppRepo) Create(ctx kratosx.Context, uid, aid uint32) (uint32, error) {
	ua := &biz.UserApp{AppID: aid, UserID: uid}
	return ua.ID, ctx.DB().Create(ua).Error
}

func (u *userAppRepo) Delete(ctx kratosx.Context, uid, aid uint32) error {
	return ctx.DB().Delete(biz.UserApp{}, "user_id=? and app_id=?", uid, aid).Error
}
