package dbs

import (
	"errors"
	"fmt"
	"sync"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/application/internal/domain/entity"
	"github.com/limes-cloud/application/internal/types"
)

type Userinfo struct {
}

var (
	userinfoIns  *Userinfo
	userinfoOnce sync.Once
)

func NewUserinfo() *Userinfo {
	userinfoOnce.Do(func() {
		userinfoIns = &Userinfo{}
	})
	return userinfoIns
}

// ListUserinfo 获取列表
func (r Userinfo) ListUserinfo(ctx kratosx.Context, req *types.ListUserinfoRequest) ([]*entity.Userinfo, uint32, error) {
	var (
		list  []*entity.Userinfo
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.Userinfo{}).Select(fs).Preload("Field", "status=true")
	db = db.Where("user_id = ?", req.UserId)

	if req.AppKeyword != nil {
		var appid uint32
		if err := ctx.DB().Model(entity.App{}).Select("id").Scan(&appid).Error; err != nil {
			return nil, 0, err
		}
		req.AppId = &appid
	}

	if req.AppId != nil {
		var ids []uint32
		if err := ctx.DB().Model(entity.AppField{}).
			Select("field_id").
			Where("app_id=?", *req.AppId).
			Scan(&ids).Error; err != nil {
			return nil, 0, err
		}
		db = db.Where("field_id in ?", ids)
	}
	if req.AppId == nil && req.AppIds != nil {
		var ids []uint32
		if err := ctx.DB().Model(entity.AppField{}).
			Select("field_id").
			Where("app_id in ?", *req.AppIds).
			Scan(&ids).Error; err != nil {
			return nil, 0, err
		}
		db = db.Where("field_id in ?", ids)
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

	return list, uint32(total), db.Find(&list).Error
}

// GetUserinfo 获取指定的数据
func (r Userinfo) GetUserinfo(ctx kratosx.Context, id uint32) (*entity.Userinfo, error) {
	var (
		userinfo = entity.Userinfo{}
		fs       = []string{"*"}
	)
	return &userinfo, ctx.DB().Select(fs).First(&userinfo, id).Error
}

// UpdateUserinfo 更新数据
func (r Userinfo) UpdateUserinfo(ctx kratosx.Context, userinfo *entity.Userinfo) error {
	return ctx.DB().Updates(userinfo).Error
}

func (r Userinfo) CheckKeywords(ctx kratosx.Context, appId uint32, keywords []string) error {
	var (
		keys []string
		ids  []uint32
	)
	if err := ctx.DB().Model(entity.AppField{}).
		Select("field_id").
		Where("app_id=?", appId).
		Scan(&ids).Error; err != nil {
		return err
	}

	if err := ctx.DB().Model(entity.Field{}).
		Select("keyword").
		Where("id in ?", ids).
		Scan(&keys).Error; err != nil {
		return err
	}

	for _, key := range keywords {
		if !valx.InList(keys, key) {
			return errors.New("not exist key:" + key)
		}
	}
	return nil
}

// CreateUserinfo 创建数据
func (r Userinfo) CreateUserinfo(ctx kratosx.Context, userinfo *entity.Userinfo) (uint32, error) {
	return userinfo.Id, ctx.DB().Create(userinfo).Error
}

// DeleteUserinfo 删除数据
func (r Userinfo) DeleteUserinfo(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Userinfo{}).Error
}
