package app

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/user-center/internal/biz/app"
	fieldbiz "github.com/limes-cloud/user-center/internal/biz/field"
	fielddata "github.com/limes-cloud/user-center/internal/data/field"
)

type repo struct {
	field fieldbiz.Repo
}

func NewRepo() biz.Repo {
	return &repo{
		field: fielddata.NewRepo(),
	}
}

func (u *repo) GetByID(ctx kratosx.Context, id uint32) (*biz.App, error) {
	var app biz.App
	return &app, ctx.DB().Preload("Channels", "status=true").
		Preload("Fields").First(&app, "id=?", id).Error
}

func (u *repo) GetByKeyword(ctx kratosx.Context, keyword string) (*biz.App, error) {
	var app biz.App
	return &app, ctx.DB().Preload("Channels", "status=true").
		Preload("Fields").First(&app, "keyword=?", keyword).Error
}

func (u *repo) Page(ctx kratosx.Context, req *biz.PageAppRequest) ([]*biz.App, uint32, error) {
	var list []*biz.App
	var total int64
	db := ctx.DB().Model(biz.App{}).Preload("Fields").Preload("Channels", "status=true")
	if req.Keyword != nil {
		db.Where("keyword=?", *req.Keyword)
	}
	if req.Name != nil {
		db.Where("name like ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *repo) Create(ctx kratosx.Context, app *biz.App) (uint32, error) {
	return app.ID, ctx.DB().Create(app).Error
}

func (u *repo) Update(ctx kratosx.Context, app *biz.App) error {
	old := biz.App{}
	if err := ctx.DB().First(&old, "id=?", app.ID).Error; err != nil {
		return err
	}
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(biz.AppChannel{}, "app_id=?", app.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(biz.AppField{}, "app_id=?", app.ID).Error; err != nil {
			return err
		}
		return tx.Updates(app).Error
	})
}

func (u *repo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.App{}, "id=?", id).Error
}
