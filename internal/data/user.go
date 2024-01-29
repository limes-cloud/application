package data

import (
	"time"

	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/internal/biz/types"
	"github.com/limes-cloud/user-center/pkg/util"
	"gorm.io/gorm"
)

type userRepo struct {
}

func NewUserRepo() biz.UserRepo {
	return &userRepo{}
}

func (u *userRepo) Create(ctx kratosx.Context, user *biz.User) (uint32, error) {
	if user.Password != "" {
		user.Password = util.ParsePwd(user.Password)
	}
	if user.From == "" {
		user.From = "system"
		user.FromDesc = "系统"
	}
	return user.ID, ctx.DB().Create(user).Error
}

func (u *userRepo) Import(ctx kratosx.Context, list []*biz.User) error {
	return ctx.DB().Model(biz.User{}).Create(&list).Error
}

func (u *userRepo) Get(ctx kratosx.Context, id uint32) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("Apps", "status=true").
		Preload("UserApps").
		Preload("Channels", "status=true").
		Preload("UserChannels").
		Preload("UserExtras").
		First(&user, "id=?", id).Error
}

func (u *userRepo) GetByPhone(ctx kratosx.Context, phone string) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("Apps", "status=true").
		Preload("UserApps").
		Preload("Channels", "status=true").
		Preload("UserChannels").
		Preload("UserExtras").
		First(&user, "phone=?", phone).Error
}

func (u *userRepo) GetByEmail(ctx kratosx.Context, email string) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("UserApps").
		Preload("Apps", "status=true").
		Preload("UserChannels").
		Preload("Channels", "status=true").
		Preload("UserExtras").
		First(&user, "email=?", email).Error
}

func (u *userRepo) GetByUsername(ctx kratosx.Context, un string) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("Apps", "status=true").
		Preload("UserApps").
		Preload("Channels", "status=true").
		Preload("UserChannels").
		Preload("UserExtras").
		First(&user, "username=?", un).Error
}

func (u *userRepo) GetByIdCard(ctx kratosx.Context, idCard string) (*biz.User, error) {
	var user biz.User
	return &user, ctx.DB().
		Preload("Apps", "status=true").
		Preload("UserApps").
		Preload("Channels", "status=true").
		Preload("UserChannels").
		Preload("UserExtras").
		First(&user, "id_card=?", idCard).Error
}

func (u *userRepo) Page(ctx kratosx.Context, options *ktypes.PageOptions) ([]*biz.User, uint32, error) {
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
func (u *userRepo) PageUser(ctx kratosx.Context, req *types.PageUserRequest) ([]*biz.User, uint32, error) {
	list, total, err := u.Page(ctx, &ktypes.PageOptions{
		Page:     req.Page,
		PageSize: req.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if req.App != nil {
				db = db.InnerJoins("UserApps", ctx.DB().Where("UserApps.app=?", req.App))
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
			if req.IdCard != nil {
				db = db.Where("id_card=?", *req.IdCard)
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
func (u *userRepo) Update(ctx kratosx.Context, user *biz.User) error {
	if user.Password != "" {
		user.Password = util.ParsePwd(user.Password)
	}
	return ctx.DB().Updates(user).Error
}

// UpdateUserApp 更新用户应用信息
func (u *userRepo) UpdateUserApp(ctx kratosx.Context, ua *biz.UserApp) error {
	if err := ctx.DB().First(&biz.UserApp{}, "user_id=? and app_id=?", ua.UserID, ua.AppID).Error; err != nil {
		return ctx.DB().Create(ua).Error
	}
	return ctx.DB().Where("user_id=? and app_id=?", ua.UserID, ua.AppID).Updates(ua).Error
}

func (u *userRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.User{}, "id=?", id).Error
}

func (u *userRepo) GetAllToken(ctx kratosx.Context, id uint32) []string {
	var token []string
	ctx.DB().Model(biz.Auth{}).Where("user_id=? and expire_at<", id, time.Now().Unix()).Scan(&token)
	return token
}
