package data

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/pkg/tree"
)

type appInterfaceRepo struct {
}

func NewAppInterfaceRepo() biz.AppInterfaceRepo {
	return &appInterfaceRepo{}
}

func (u *appInterfaceRepo) All(ctx kratosx.Context, appId uint32) ([]*biz.AppInterface, error) {
	var list []*biz.AppInterface
	return list, ctx.DB().Model(biz.AppInterface{}).Find(&list, "app_id=?", appId).Error
}

func (u *appInterfaceRepo) Create(ctx kratosx.Context, ai *biz.AppInterface) (uint32, error) {
	return ai.ID(), ctx.DB().Transaction(func(tx *gorm.DB) error {
		app := biz.App{}
		if err := tx.First(&app, "id=?", ai.AppID).Error; err != nil {
			return err
		}
		if ai.Type == "A" {
			if _, err := ctx.Authentication().Enforce().AddPolicy(app.Keyword, *ai.Path, *ai.Method); err != nil {
				return err
			}
		}
		return tx.Create(ai).Error
	})
}

func (u *appInterfaceRepo) Update(ctx kratosx.Context, ai *biz.AppInterface) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		old := biz.AppInterface{}
		if err := tx.First(&old, "id=?", ai.ID).Error; err != nil {
			return err
		}

		app := biz.App{}
		if err := tx.First(&app, "id=?", old.AppID).Error; err != nil {
			return err
		}

		if old.Method != ai.Method || old.Path != ai.Method {
			if _, err := ctx.Authentication().Enforce().UpdatePolicy(
				[]string{app.Keyword, *old.Path, *old.Method},
				[]string{app.Keyword, *ai.Path, *ai.Method},
			); err != nil {
				return err
			}
		}

		return tx.Updates(ai).Error
	})
}

func (u *appInterfaceRepo) Delete(ctx kratosx.Context, id uint32) error {
	old := biz.AppInterface{}
	if err := ctx.DB().First(&old, "id=?", id).Error; err != nil {
		return err
	}

	app := biz.App{}
	if err := ctx.DB().First(&app, "id=?", old.AppID).Error; err != nil {
		return err
	}

	var list []*biz.AppInterface
	if err := ctx.DB().Model(biz.AppInterface{}).Find(&list, "app_id=?", old.AppID).Error; err != nil {
		return err
	}

	var (
		tl       []tree.Tree
		policies [][]string
		apiSet   = map[uint32]*biz.AppInterface{}
	)
	for _, item := range list {
		tl = append(tl, item)
		if item.Type == "A" {
			apiSet[item.ID()] = item
		}
	}
	// 获取删除菜单的下级菜单id
	t := tree.BuildTreeByID(tl, id)
	ids := tree.GetTreeID(t)

	// 筛选出其中是Api的路由
	for _, id := range ids {
		if apiSet[id] != nil {
			policies = append(policies, []string{app.Keyword, *apiSet[id].Path, *apiSet[id].Method})
		}
	}

	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		// 删除权限表
		if _, err := ctx.Authentication().Enforce().RemovePolicies(policies); err != nil {
			return err
		}
		return tx.Delete(biz.AppInterface{}, "id=?", id).Error
	})
}
