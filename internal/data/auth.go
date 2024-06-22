package data

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	biz "github.com/limes-cloud/usercenter/internal/biz/auth"
	"github.com/limes-cloud/usercenter/internal/data/model"
	"github.com/limes-cloud/usercenter/internal/pkg/resource"
)

type authRepo struct {
}

func NewAuthRepo() biz.Repo {
	return &authRepo{}
}

// ToAuthEntity model转entity
func (r authRepo) ToAuthEntity(m *model.Auth) *biz.Auth {
	e := &biz.Auth{}
	_ = valx.Transform(m, e)
	return e
}

// ToAuthModel entity转model
func (r authRepo) ToAuthModel(e *biz.Auth) *model.Auth {
	m := &model.Auth{}
	_ = valx.Transform(e, m)
	return m
}

// ToOAuthEntity model转entity
func (r authRepo) ToOAuthEntity(m *model.OAuth) *biz.OAuth {
	e := &biz.OAuth{}
	_ = valx.Transform(m, e)
	return e
}

// ToOAuthModel entity转model
func (r authRepo) ToOAuthModel(e *biz.OAuth) *model.OAuth {
	m := &model.OAuth{}
	_ = valx.Transform(e, m)
	return m
}

// ListAuth 获取列表
func (r authRepo) ListAuth(ctx kratosx.Context, req *biz.ListAuthRequest) ([]*biz.Auth, uint32, error) {
	var (
		bs    []*biz.Auth
		ms    []*model.Auth
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(model.Auth{}).Select(fs).Preload("App", "status=true").Where("user_id = ?", req.UserId)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}
	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		b := r.ToAuthEntity(m)
		if b.App.Logo != "" {
			b.App.Logo = resource.GetURLBySha(ctx, b.App.Logo)
		}
		bs = append(bs, b)
	}
	return bs, uint32(total), nil
}

// CreateAuth 创建数据
func (r authRepo) CreateAuth(ctx kratosx.Context, req *biz.Auth) (uint32, error) {
	m := r.ToAuthModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// UpdateAuth 更新数据
func (r authRepo) UpdateAuth(ctx kratosx.Context, req *biz.Auth) error {
	return ctx.DB().Where("user_id=?", req.UserId).Where("app_id=?", req.AppId).Updates(r.ToAuthModel(req)).Error
}

// UpdateAuthStatus 更新数据状态
func (r authRepo) UpdateAuthStatus(ctx kratosx.Context, req *biz.UpdateAuthStatusRequest) error {
	return ctx.DB().Model(model.Auth{}).
		Where("id= ?", req.Id).
		Updates(map[string]any{
			"status":       req.Status,
			"disable_desc": req.DisableDesc,
		}).Error
}

// DeleteAuth 删除数据
func (r authRepo) DeleteAuth(ctx kratosx.Context, userId uint32, appId uint32) error {
	db := ctx.DB().Where("user_id = ? and app_id = ?", userId, appId).Delete(&model.Auth{})
	return db.Error
}

// GetAuthByUserIdAndAppId 获取指定数据
func (r authRepo) GetAuthByUserIdAndAppId(ctx kratosx.Context, userId uint32, appId uint32) (*biz.Auth, error) {
	var (
		m  = model.Auth{}
		fs = []string{"id", "status", "disable_desc", "token", "logged_at", "expired_at", "created_at"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("userId = ?", userId).Where("appId = ?", appId).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToAuthEntity(&m), nil
}

func (r authRepo) HasUserByEmail(ctx kratosx.Context, email string) bool {
	var count int64 = 0
	_ = ctx.DB().Model(model.User{}).Where("email=?", email).Count(&count)
	return count > 0
}

func (r authRepo) HasUserByUsername(ctx kratosx.Context, username string) bool {
	var count int64 = 0
	_ = ctx.DB().Model(model.User{}).Where("username=?", username).Count(&count)
	return count > 0
}

// ListOAuth 获取列表
func (r authRepo) ListOAuth(ctx kratosx.Context, req *biz.ListOAuthRequest) ([]*biz.OAuth, uint32, error) {
	var (
		bs    []*biz.OAuth
		ms    []*model.OAuth
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(model.OAuth{}).Select(fs).Preload("Channel", "status=true").Where("user_id = ?", req.UserId)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}
	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		b := r.ToOAuthEntity(m)
		if b.Channel.Logo != "" {
			b.Channel.Logo = resource.GetURLBySha(ctx, b.Channel.Logo)
		}
		bs = append(bs, b)
	}
	return bs, uint32(total), nil
}

// CreateOAuth 创建数据
func (r authRepo) CreateOAuth(ctx kratosx.Context, req *biz.OAuth) (string, error) {
	m := r.ToOAuthModel(req)
	if err := ctx.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(m).Error; err != nil {
		return "", err
	}
	uid := crypto.MD5([]byte(uuid.NewString()))
	if err := ctx.Redis().Set(ctx, uid, m.Id, 10*time.Minute).Err(); err != nil {
		return "", err
	}
	return uid, nil
}

// GetOAuthByUid 通过授id获取三方授权数据
func (r authRepo) GetOAuthByUid(ctx kratosx.Context, uid string) (*biz.OAuth, error) {
	var id uint32
	if err := ctx.Redis().Get(ctx, uid).Scan(&id); err != nil {
		return nil, err
	}
	if id == 0 {
		return nil, errors.New("授权已失效")
	}

	var oauth model.OAuth
	if err := ctx.DB().Where("id = ?", id).First(&oauth).Error; err != nil {
		return nil, err
	}
	return r.ToOAuthEntity(&oauth), nil
}

// UpdateOAuth 更新数据
func (r authRepo) UpdateOAuth(ctx kratosx.Context, req *biz.OAuth) error {
	return ctx.DB().Where("user_id=?", req.UserId).Where("channel_id=?", req.ChannelId).Updates(r.ToOAuthModel(req)).Error
}

// DeleteOAuth 删除数据
func (r authRepo) DeleteOAuth(ctx kratosx.Context, userId uint32, channelId uint32) error {
	db := ctx.DB().Where("user_id = ? and channel_id = ?", userId, channelId).Delete(&model.OAuth{})
	return db.Error
}

func (r authRepo) IsBindUser(ctx kratosx.Context, cid uint32, aid string) bool {
	var count int64
	ctx.DB().Model(model.OAuth{}).
		Where("`user_id` IS NOT NULL").
		Where("channel_id=?", cid).
		Where("auth_id=?", aid).
		Count(&count)
	return count != 0
}

func (r authRepo) HasAppScope(ctx kratosx.Context, aid uint32, uid uint32) error {
	var (
		auth model.Auth
		app  model.App
	)

	if err := ctx.DB().Where("id = ?", aid).First(&app).Error; err != nil {
		return err
	}

	if err := ctx.DB().Where("app_id = ?", aid).Where("user_id = ?", uid).First(&auth).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if app.AllowRegistry == nil || !*app.AllowRegistry {
			return errors.New("无应用权限")
		}
	}

	if auth.Status == nil || !*auth.Status {
		desc := "用户已被禁用登陆此应用"
		if auth.DisableDesc != nil {
			desc = *auth.DisableDesc
		}
		return errors.New(desc)
	}

	return nil
}

func (r authRepo) GetUserByUsername(ctx kratosx.Context, username string) (*biz.User, error) {
	var user model.User

	if err := ctx.DB().Where("username=?", username).First(&user).Error; err != nil {
		return nil, err
	}

	if user.Status == nil || !*user.Status {
		desc := "用户已被禁用"
		if user.DisableDesc != nil {
			desc = *user.DisableDesc
		}
		return nil, errors.New(desc)
	}
	return &biz.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		From:     user.From,
		FromDesc: user.FromDesc,
		Password: user.Password,
	}, nil
}

func (r authRepo) GetUserByEmail(ctx kratosx.Context, email string) (*biz.User, error) {
	var user model.User

	if err := ctx.DB().Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if user.Status == nil || !*user.Status {
		desc := "用户已被禁用"
		if user.DisableDesc != nil {
			desc = *user.DisableDesc
		}
		return nil, errors.New(desc)
	}
	return &biz.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		From:     user.From,
		FromDesc: user.FromDesc,
		Password: user.Password,
	}, nil
}

func (r authRepo) GetUserByCA(ctx kratosx.Context, cid uint32, aid string) (*biz.User, error) {
	var (
		oauth model.OAuth
		user  model.User
	)
	if err := ctx.DB().Where("channel_id=?", cid).Where("auth_id=?", aid).First(&oauth).Error; err != nil {
		return nil, err
	}

	if err := ctx.DB().Where("id=?", oauth.UserId).First(&user).Error; err != nil {
		return nil, err
	}

	if user.Status == nil || !*user.Status {
		desc := "用户已被禁用"
		if user.DisableDesc != nil {
			desc = *user.DisableDesc
		}
		return nil, errors.New(desc)
	}
	return &biz.User{Id: user.Id}, nil
}

func (r authRepo) GetOAuthByCA(ctx kratosx.Context, cid uint32, aid string) (*biz.OAuth, error) {
	var auth model.OAuth
	if err := ctx.DB().Where("channel_id=?", cid).Where("auth_id=?", aid).First(&auth).Error; err != nil {
		return nil, err
	}
	return r.ToOAuthEntity(&auth), nil
}

func (r authRepo) GetApp(ctx kratosx.Context, keyword string) (*biz.App, error) {
	var app model.App
	if err := ctx.DB().Where("keyword=?", keyword).First(&app).Error; err != nil {
		return nil, err
	}
	return &biz.App{
		Id:            app.Id,
		Logo:          app.Logo,
		Keyword:       app.Keyword,
		Name:          app.Name,
		AllowRegister: app.AllowRegistry,
	}, nil
}

func (r authRepo) GetAppChannel(ctx kratosx.Context, ak, ck string) (*biz.Channel, error) {
	app, err := r.GetApp(ctx, ak)
	if err != nil {
		return nil, err
	}

	var channel model.Channel
	if err := ctx.DB().Where("keyword=?", ck).First(&channel).Error; err != nil {
		return nil, err
	}
	if channel.Status == nil || !*channel.Status {
		return nil, errors.New("当前渠道已被禁用")
	}

	var count int64
	ctx.DB().Model(model.AppChannel{}).Where("app_id=? and channel_id=?", app.Id, channel.Id).Count(&count)
	if count == 0 {
		return nil, errors.New("应用暂未开通此登陆渠道")
	}

	ak = ""
	if channel.Ak != nil {
		ak = *channel.Ak
	}

	sk := ""
	if channel.Sk != nil {
		sk = *channel.Sk
	}

	extra := ""
	if channel.Extra != nil {
		extra = *channel.Extra
	}

	return &biz.Channel{
		Id:      channel.Id,
		Ak:      ak,
		Sk:      sk,
		Extra:   extra,
		Logo:    channel.Logo,
		Keyword: channel.Keyword,
		Name:    channel.Name,
	}, nil
}

func (r authRepo) Register(ctx kratosx.Context, user *biz.User) error {
	mu := model.User{
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		From:     user.From,
		FromDesc: user.FromDesc,
		Status:   user.Status,
		NickName: user.Nickname,
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if err := ctx.DB().Create(&mu).Error; err != nil {
			return err
		}
		var apps []*model.Auth
		for _, appId := range user.AppIds {
			apps = append(apps, &model.Auth{
				UserId: mu.Id,
				AppId:  appId,
				Status: proto.Bool(true),
			})
		}
		return ctx.DB().Create(&apps).Error
	})
}

func (r authRepo) BindOAuthByUid(ctx kratosx.Context, uid uint32, aid string) error {
	auth, err := r.GetOAuthByUid(ctx, aid)
	if err != nil {
		return err
	}
	return ctx.DB().Model(model.OAuth{}).Where("id=?", auth.Id).Update("user_id", uid).Error
}
