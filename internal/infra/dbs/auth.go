package dbs

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"

	"github.com/limes-cloud/application/internal/domain/entity"
	"github.com/limes-cloud/application/internal/types"
)

type Auth struct {
}

var (
	authIns  *Auth
	authOnce sync.Once
)

func NewAuth() *Auth {
	authOnce.Do(func() {
		authIns = &Auth{}
	})
	return authIns
}

// ListAuth 获取列表
func (r Auth) ListAuth(ctx kratosx.Context, req *types.ListAuthRequest) ([]*entity.Auth, uint32, error) {
	var (
		list  []*entity.Auth
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Auth{}).Select(fs).
		Preload("App", "status=true").
		Where("user_id = ?", req.UserId)

	if req.AppIds != nil {
		db = db.Where("app_id in ?", req.AppIds)
	}

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
	if err := db.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateAuth 创建数据
func (r Auth) CreateAuth(ctx kratosx.Context, auth *entity.Auth) (uint32, error) {
	return auth.Id, ctx.DB().Create(auth).Error
}

// UpsertAuth 存在更新数据，不存在新增
func (r Auth) UpsertAuth(ctx kratosx.Context, auth *entity.Auth) error {
	var count int64
	if ctx.DB().
		Model(entity.Auth{}).
		Where("user_id=?", auth.UserId).
		Where("app_id=?", auth.AppId).
		Count(&count); count == 0 {
		auth.Status = proto.Bool(true)
		return ctx.DB().Create(auth).Error
	}
	return ctx.DB().Where("user_id=?", auth.UserId).
		Where("app_id=?", auth.AppId).
		Updates(auth).Error
}

// UpdateAuthStatus 更新数据状态
func (r Auth) UpdateAuthStatus(ctx kratosx.Context, req *types.UpdateAuthStatusRequest) error {
	return ctx.DB().Model(entity.Auth{}).
		Where("id= ?", req.Id).
		Updates(map[string]any{
			"status":       req.Status,
			"disable_desc": req.DisableDesc,
		}).Error
}

// DeleteAuth 删除数据
func (r Auth) DeleteAuth(ctx kratosx.Context, userId uint32, appId uint32) error {
	return ctx.DB().Where("user_id = ? and app_id = ?", userId, appId).Delete(&entity.Auth{}).Error
}

// GetAuthByUA 获取指定数据
func (r Auth) GetAuthByUA(ctx kratosx.Context, userId uint32, appId uint32) (*entity.Auth, error) {
	var (
		auth = entity.Auth{}
		fs   = []string{"id", "status", "disable_desc", "token", "logged_at", "expired_at", "created_at"}
	)
	return &auth, ctx.DB().Select(fs).
		Where("userId = ?", userId).
		Where("appId = ?", appId).First(&auth).Error
}

// ListOAuth 获取列表
func (r Auth) ListOAuth(ctx kratosx.Context, req *types.ListOAuthRequest) ([]*entity.OAuth, uint32, error) {
	var (
		list  []*entity.OAuth
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.OAuth{}).Select(fs).
		Preload("Channel", "status=true").
		Where("user_id = ?", req.UserId)

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
	if err := db.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, uint32(total), nil
}

// CreateOAuth 创建数据
func (r Auth) CreateOAuth(ctx kratosx.Context, oauth *entity.OAuth) (string, error) {
	if err := ctx.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(oauth).Error; err != nil {
		return "", err
	}
	uid := crypto.MD5([]byte(uuid.NewString()))
	if err := ctx.Redis().Set(ctx, uid, oauth.Id, 10*time.Minute).Err(); err != nil {
		return "", err
	}
	return uid, nil
}

// GetOAuthByUid 通过授id获取三方授权数据
func (r Auth) GetOAuthByUid(ctx kratosx.Context, uid string) (*entity.OAuth, error) {
	var id uint32
	if err := ctx.Redis().Get(ctx, uid).Scan(&id); err != nil {
		return nil, err
	}
	if id == 0 {
		return nil, errors.New("授权已失效")
	}

	var oauth entity.OAuth
	return &oauth, ctx.DB().Where("id = ?", id).First(&oauth).Error
}

// UpdateOAuth 更新数据
func (r Auth) UpdateOAuth(ctx kratosx.Context, oauth *entity.OAuth) error {
	var count int64
	if ctx.DB().
		Model(entity.OAuth{}).
		Where("user_id=?", oauth.UserId).
		Where("channel_id=?", oauth.ChannelId).
		Count(&count); count == 0 {
		return ctx.DB().Create(oauth).Error
	}
	return ctx.DB().
		Where("user_id=?", oauth.UserId).
		Where("channel_id=?", oauth.ChannelId).
		Updates(oauth).Error
}

// DeleteOAuth 删除数据
func (r Auth) DeleteOAuth(ctx kratosx.Context, userId uint32, channelId uint32) error {
	return ctx.DB().Where("user_id = ? and channel_id = ?", userId, channelId).Delete(&entity.OAuth{}).Error
}

func (r Auth) IsBindOAuth(ctx kratosx.Context, cid uint32, aid string) bool {
	var count int64
	ctx.DB().Model(entity.OAuth{}).
		Where("`user_id` IS NOT NULL").
		Where("channel_id=?", cid).
		Where("auth_id=?", aid).
		Count(&count)
	return count != 0
}

func (r Auth) GetOAuthByCA(ctx kratosx.Context, cid uint32, aid string) (*entity.OAuth, error) {
	var oauth entity.OAuth
	return &oauth, ctx.DB().Where("channel_id=?", cid).Where("auth_id=?", aid).First(&oauth).Error
}

func (r Auth) BindOAuthByUid(ctx kratosx.Context, uid uint32, aid string) error {
	auth, err := r.GetOAuthByUid(ctx, aid)
	if err != nil {
		return err
	}
	return ctx.DB().Model(entity.OAuth{}).Where("id=?", auth.Id).Update("user_id", uid).Error
}
