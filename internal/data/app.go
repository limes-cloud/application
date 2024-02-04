package data

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/internal/biz/types"
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
	// return app.ID, ctx.DB().Transaction(func(tx *gorm.DB) error {
	//	if err := tx.Create(app).Error; err != nil {
	//		return err
	//	}
	//	ai := biz.AppInterface{AppID: app.ID, Type: "G", Title: app.Name}
	//	return tx.Create(ai).Error
	// })
	return app.ID, ctx.DB().Create(app).Error
}

func (u *appRepo) Update(ctx kratosx.Context, app *biz.App) error {
	old := biz.App{}
	if err := ctx.DB().First(&old, "id=?", app.ID).Error; err != nil {
		return err
	}

	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(biz.AppChannel{}, "app_id=?", app.ID).Error; err != nil {
			return err
		}
		// 先获取全部权限
		if old.Keyword != app.Keyword {
			enforce := ctx.Authentication().Enforce()
			policies := enforce.GetFilteredPolicy(0, old.Keyword)
			var newPolicies [][]string
			for _, item := range policies {
				newPolicies = append(newPolicies, []string{
					app.Keyword, item[1], item[2],
				})
			}
			if _, err := enforce.UpdatePolicies(policies, newPolicies); err != nil {
				return err
			}
		}

		return tx.Updates(app).Error
	})
}

func (u *appRepo) Delete(ctx kratosx.Context, id uint32) error {
	app := biz.App{}
	if err := ctx.DB().First(&app, "id=?", id).Error; err != nil {
		return err
	}
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if _, err := ctx.Authentication().Enforce().RemoveFilteredPolicy(0, app.Keyword); err != nil {
			return err
		}
		return ctx.DB().Delete(biz.App{}, "id=?", id).Error
	})
}
