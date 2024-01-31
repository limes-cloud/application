package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/internal/biz/types"
	"gorm.io/gorm"
)

type appRepo struct {
}

func NewAppRepo() biz.AppRepo {
	return &appRepo{}
}

func (u *appRepo) GetByID(ctx kratosx.Context, id uint32) (*biz.App, error) {
	var app biz.App
	return &app, ctx.DB().Model(biz.App{}).Preload("Channels").Find(&app, "id=?", id).Error
}

func (u *appRepo) GetByKeyword(ctx kratosx.Context, keyword string) (*biz.App, error) {
	var app biz.App
	return &app, ctx.DB().Model(biz.App{}).Preload("Channels").Find(&app, "keyword=?", keyword).Error
}

func (u *appRepo) Page(ctx kratosx.Context, req *types.PageAppRequest) ([]*biz.App, uint32, error) {
	var list []*biz.App
	var total int64
	db := ctx.DB().Model(biz.App{}).Preload("Channels", "status=true")
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

func (u *appRepo) Create(ctx kratosx.Context, app *biz.App) (uint32, error) {
	return app.ID, ctx.DB().Create(app).Error
}

func (u *appRepo) Update(ctx kratosx.Context, app *biz.App) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(biz.AppChannel{}, "app_id=?", app.ID).Error; err != nil {
			return err
		}
		return tx.Updates(app).Error
	})
}

func (u *appRepo) Delete(ctx kratosx.Context, id uint32) error {
	//if err := ctx.DB().First(biz.UserApp{}, "app_id=?", id); err == nil {
	//	return errors.New("应用正在使用中，无法删除")
	//}
	return ctx.DB().Delete(biz.App{}, "id=?", id).Error
}
