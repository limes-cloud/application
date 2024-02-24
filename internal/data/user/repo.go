package user

import (
	"time"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	ktypes "github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"

	appbiz "github.com/limes-cloud/user-center/internal/biz/app"
	biz "github.com/limes-cloud/user-center/internal/biz/user"
	appdata "github.com/limes-cloud/user-center/internal/data/app"
)

type repo struct {
	app appbiz.Repo
}

func NewRepo() biz.Repo {
	return &repo{
		app: appdata.NewRepo(),
	}
}

func (u *repo) Add(ctx kratosx.Context, user *biz.User) (uint32, error) {
	if user.Password != "" {
		user.Password = util.ParsePwd(user.Password)
	}
	if user.From == "" {
		user.From = "system"
		user.FromDesc = "系统"
	}
	return user.ID, ctx.DB().Create(user).Error
}

func (u *repo) Import(ctx kratosx.Context, list []*biz.User) error {
	return ctx.DB().Model(biz.User{}).Create(&list).Error
}

func (u *repo) GetBase(ctx kratosx.Context, id uint32) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().First(&user, "id=?", id).Error
}

func (u *repo) Get(ctx kratosx.Context, id uint32) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("UserApps").
		Preload("UserApps.App", "status=true").
		Preload("UserApps.App.Fields").
		Preload("Auths").
		Preload("Auths.Channel", "status=true").
		Preload("UserExtras").
		First(&user, "id=?", id).Error
}

func (u *repo) GetByPhone(ctx kratosx.Context, phone string) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("UserApps").
		Preload("UserApps.App", "status=true").
		Preload("UserApps.App.Fields").
		Preload("Auths").
		Preload("Auths.Channel", "status=true").
		Preload("UserExtras").
		First(&user, "phone=?", phone).Error
}

func (u *repo) GetByEmail(ctx kratosx.Context, email string) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("UserApps").
		Preload("UserApps.App", "status=true").
		Preload("UserApps.App.Fields").
		Preload("Auths").
		Preload("Auths.Channel", "status=true").
		Preload("UserExtras").
		First(&user, "email=?", email).Error
}

func (u *repo) GetByUsername(ctx kratosx.Context, un string) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("UserApps").
		Preload("UserApps.App", "status=true").
		Preload("UserApps.App.Fields").
		Preload("Auths").
		Preload("Auths.Channel", "status=true").
		Preload("UserExtras").
		First(&user, "username=?", un).Error
}

func (u *repo) Page(ctx kratosx.Context, options *ktypes.PageOptions) ([]*biz.User, uint32, error) {
	var list []*biz.User
	var total int64

	db := ctx.DB().Model(biz.User{})
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// PageUser 获取用户
func (u *repo) PageUser(ctx kratosx.Context, req *biz.PageUserRequest) ([]*biz.User, uint32, error) {
	list, total, err := u.Page(ctx, &ktypes.PageOptions{
		Page:     req.Page,
		PageSize: req.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if req.App != nil {
				app, _ := u.app.GetByKeyword(ctx, *req.App)
				db = db.InnerJoins("UserApps", ctx.DB().Where("UserApps.app_id=?", app.ID))
			}
			if req.Username != nil {
				db = db.Where("username=?", *req.Username)
			}
			if req.Email != nil {
				db = db.Where("email=?", *req.Email)
			}
			if req.Phone != nil {
				db = db.Where("phone=?", *req.Phone)
			}
			if len(req.InIds) != 0 {
				db.Where("user.id in ?", req.InIds)
			}
			if len(req.NotInIds) != 0 {
				db.Where("user.id not in ?", req.NotInIds)
			}
			return db
		},
	})
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// Update 更新用户信息
func (u *repo) Update(ctx kratosx.Context, user *biz.User) error {
	if user.Password != "" {
		user.Password = util.ParsePwd(user.Password)
	}
	return ctx.DB().Updates(user).Error
}

func (u *repo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.User{}, "id=?", id).Error
}

func (u *repo) GetJwtTokens(ctx kratosx.Context, id uint32) []string {
	var token []string
	ctx.DB().Model(biz.Auth{}).Where("user_id=? and jwt_expire_at<", id, time.Now().Unix()).Scan(&token)
	return token
}

func (u *repo) AddUserApp(ctx kratosx.Context, uid, aid uint32) (uint32, error) {
	ua := &biz.UserApp{AppID: aid, UserID: uid}
	return ua.ID, ctx.DB().Create(ua).Error
}

func (u *repo) UpdateUserApp(ctx kratosx.Context, in *biz.UserApp) error {
	if err := ctx.DB().Where("user_id=? and app_id=?", in.UserID, in.AppID).
		First(&biz.UserApp{}).Error; err != nil {
		return ctx.DB().Create(in).Error
	}
	return ctx.DB().Where("user_id=? and app_id=?", in.UserID, in.AppID).Updates(in).Error
}

func (u *repo) DeleteUserApp(ctx kratosx.Context, uid, aid uint32) error {
	return ctx.DB().Delete(biz.UserApp{}, "user_id=? and app_id=?", uid, aid).Error
}

func (u *repo) GetAuthByCU(ctx kratosx.Context, cid, uid uint32) (*biz.Auth, error) {
	var res biz.Auth
	return &res, ctx.DB().Where("channel_id=? and user_id=?", cid, uid).First(&res).Error
}

func (u *repo) GetAuthByCA(ctx kratosx.Context, cid uint32, aid string) (*biz.Auth, error) {
	var res biz.Auth
	return &res, ctx.DB().Where("channel_id=? and auth_id=?", cid, aid).First(&res).Error
}

func (u *repo) AddAuth(ctx kratosx.Context, uc *biz.Auth) (uint32, error) {
	return uc.ID, ctx.DB().Create(uc).Error
}

func (u *repo) UpdateAuth(ctx kratosx.Context, uc *biz.Auth) (uint32, error) {
	return uc.ID, ctx.DB().Where("channel_id=? and user_id=?", uc.ChannelID, uc.UserID).Updates(uc).Error
}

func (u *repo) DeleteAuth(ctx kratosx.Context, uid, aid uint32) error {
	return ctx.DB().Delete(biz.Auth{}, "user_id=? and channel_id=?", uid, aid).Error
}

func (u *repo) HasUsername(ctx kratosx.Context, un string) bool {
	var count int64
	ctx.DB().Model(biz.User{}).Where("username=?", un).Count(&count)
	return count != 0
}

func (u *repo) HasUserEmail(ctx kratosx.Context, email string) bool {
	var count int64
	ctx.DB().Model(biz.User{}).Where("email=?", email).Count(&count)
	return count != 0
}

func (u *repo) HasUserPhone(ctx kratosx.Context, phone string) bool {
	var count int64
	ctx.DB().Model(biz.User{}).Where("phone=?", phone).Count(&count)
	return count != 0
}

func (u *repo) GetApp(ctx kratosx.Context, keyword string) (*appbiz.App, error) {
	return u.app.GetByKeyword(ctx, keyword)
}

func (u *repo) UpdateAuthByCU(ctx kratosx.Context, auth *biz.Auth) error {
	if err := ctx.DB().Where("user_id=? and channel_id=?", auth.UserID, auth.ChannelID).
		First(&biz.Auth{}).Error; err != nil {
		return ctx.DB().Create(auth).Error
	}
	return ctx.DB().Where("user_id=? and channel_id=?", auth.UserID, auth.ChannelID).Updates(auth).Error
}
