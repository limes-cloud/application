package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/user-center/internal/biz"
)

type userChannelRepo struct {
}

func NewUserChannelRepo() biz.UserChannelRepo {
	return &userChannelRepo{}
}

func (u *userChannelRepo) Create(ctx kratosx.Context, uid, aid uint32) (uint32, error) {
	ua := &biz.UserChannel{ChannelID: aid, UserID: uid}
	return ua.ID, ctx.DB().Create(ua).Error
}

func (u *userChannelRepo) Delete(ctx kratosx.Context, uid, aid uint32) error {
	return ctx.DB().Delete(biz.UserChannel{}, "user_id=? and channel_id=?", uid, aid).Error
}
