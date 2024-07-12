package data

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/usercenter/internal/biz/app"
	"github.com/limes-cloud/usercenter/internal/data/model"
	"github.com/limes-cloud/usercenter/internal/pkg/resource"
)

type appRepo struct {
}

func NewAppRepo() biz.Repo {
	return &appRepo{}
}

// ToAppEntity model转entity
func (r appRepo) ToAppEntity(m *model.App) *biz.App {
	e := &biz.App{}
	_ = valx.Transform(m, e)
	return e
}

// ToAppModel entity转model
func (r appRepo) ToAppModel(e *biz.App) *model.App {
	m := &model.App{}
	_ = valx.Transform(e, m)
	return m
}

// GetAppByKeyword 获取指定数据
func (r appRepo) GetAppByKeyword(ctx kratosx.Context, keyword string) (*biz.App, error) {
	var (
		m  = model.App{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	db = db.Preload("Channels", "status=true").Preload("Fields", "status=true")
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	b := r.ToAppEntity(&m)
	if b.Logo != "" {
		b.LogoUrl = resource.GetURLBySha(ctx, b.Logo)
	}
	for ind, item := range b.Channels {
		b.Channels[ind].Logo = resource.GetURLBySha(ctx, item.Logo)
	}
	return b, nil
}

// GetApp 获取指定的数据
func (r appRepo) GetApp(ctx kratosx.Context, id uint32) (*biz.App, error) {
	var (
		m  = model.App{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	db = db.Preload("Channels", "status=true").Preload("Fields", "status=true")
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	b := r.ToAppEntity(&m)
	if b.Logo != "" {
		b.LogoUrl = resource.GetURLBySha(ctx, b.Logo)
	}
	for ind, item := range b.Channels {
		b.Channels[ind].Logo = resource.GetURLBySha(ctx, item.Logo)
	}
	return b, nil
}

// ListApp 获取列表
func (r appRepo) ListApp(ctx kratosx.Context, req *biz.ListAppRequest) ([]*biz.App, uint32, error) {
	var (
		bs    []*biz.App
		ms    []*model.App
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.App{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
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

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		b := r.ToAppEntity(m)
		if b.Logo != "" {
			b.LogoUrl = resource.GetURLBySha(ctx, b.Logo)
		}
		bs = append(bs, b)
	}
	return bs, uint32(total), nil
}

// CreateApp 创建数据
func (r appRepo) CreateApp(ctx kratosx.Context, req *biz.App) (uint32, error) {
	m := r.ToAppModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// UpdateApp 更新数据
func (r appRepo) UpdateApp(ctx kratosx.Context, req *biz.App) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if len(req.AppChannels) != 0 {
			if err := tx.Where("app_id=?", req.Id).Delete(model.AppChannel{}).Error; err != nil {
				return err
			}
		}

		if err := tx.Where("app_id=?", req.Id).Delete(model.AppField{}).Error; err != nil {
			return err
		}

		return tx.Updates(r.ToAppModel(req)).Error
	})
}

// UpdateAppStatus 更新数据状态
func (r appRepo) UpdateAppStatus(ctx kratosx.Context, req *biz.UpdateAppStatusRequest) error {
	return ctx.DB().Model(model.App{}).
		Where("id=?", req.Id).
		Updates(map[string]any{
			"status":       req.Status,
			"disable_desc": req.DisableDesc,
		}).Error
}

// DeleteApp 删除数据
func (r appRepo) DeleteApp(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&model.App{}).Error
}
