package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/user-center/internal/biz"
)

type authRepo struct {
	userRepo    biz.UserRepo
	channelRepo biz.ChannelRepo
	appRepo     biz.AppRepo
}

func NewAuthRepo() biz.AuthRepo {
	return &authRepo{
		userRepo:    NewUserRepo(),
		channelRepo: NewChannelRepo(),
		appRepo:     NewAppRepo(),
	}
}

func (u *authRepo) Register(ctx kratosx.Context, user *biz.User) (uint32, error) {
	return u.userRepo.Create(ctx, user)
}

func (u *authRepo) GetChannel(ctx kratosx.Context, platform string) (*biz.Channel, error) {
	return u.channelRepo.GetByPlatform(ctx, platform)
}

func (u *authRepo) GetApp(ctx kratosx.Context, keyword string) (*biz.App, error) {
	return u.appRepo.GetByKeyword(ctx, keyword)
}

func (u *authRepo) GetUserByPhone(ctx kratosx.Context, phone string) (*biz.User, error) {
	return u.userRepo.GetByPhone(ctx, phone)
}

func (u *authRepo) GetUserByEmail(ctx kratosx.Context, email string) (*biz.User, error) {
	return u.userRepo.GetByEmail(ctx, email)
}

func (u *authRepo) HasUserEmail(ctx kratosx.Context, email string) bool {
	var count int64
	ctx.DB().Model(biz.User{}).Where("email=?", email).Count(&count)
	return count != 0
}

func (u *authRepo) HasUserPhone(ctx kratosx.Context, phone string) bool {
	var count int64
	ctx.DB().Model(biz.User{}).Where("phone=?", phone).Count(&count)
	return count != 0
}

func (u *authRepo) HasUsername(ctx kratosx.Context, username string) bool {
	var count int64
	ctx.DB().Model(biz.User{}).Where("username=?", username).Count(&count)
	return count != 0
}

func (u *authRepo) GetUserByUsername(ctx kratosx.Context, username string) (*biz.User, error) {
	return u.userRepo.GetByUsername(ctx, username)
}

func (u *authRepo) Upsert(ctx kratosx.Context, auth *biz.Auth) error {
	if err := ctx.DB().Where("user_id=? and app_id=? and channel_id=?",
		auth.UserID, auth.AppID, auth.ChannelID).
		First(&biz.Auth{}).Error; err != nil {
		return ctx.DB().Create(auth).Error
	}
	return ctx.DB().Where("user_id=? and app_id=? and channel_id=?",
		auth.UserID, auth.AppID, auth.ChannelID).Updates(auth).Error
}

func (u *authRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Auth{}, "id=?", id).Error
}

func (u *authRepo) UpsertUserApp(ctx kratosx.Context, user *biz.UserApp) error {
	return u.userRepo.UpdateUserApp(ctx, user)
}

func (u *authRepo) UpsertUserChannel(ctx kratosx.Context, ua *biz.UserChannel) error {
	if err := ctx.DB().Where("user_id=? and channel_id=?", ua.UserID, ua.ChannelID).
		First(&biz.UserChannel{}).Error; err != nil {
		return ctx.DB().Create(ua).Error
	}
	return ctx.DB().Where("user_id=? and channel_id=?", ua.UserID, ua.ChannelID).Updates(ua).Error
}
