package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/internal/biz/types"
	"gorm.io/gorm"
)

type sceneRepo struct {
}

func NewSceneRepo() biz.SceneRepo {
	return &sceneRepo{}
}

func (u *sceneRepo) GetByKeyword(ctx kratosx.Context, keyword string) (*biz.Scene, error) {
	var scene biz.Scene
	return &scene, ctx.DB().Model(biz.Scene{}).Preload("Agreements", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,created_at,name").Where("status=true")
	}).Find(&scene, "keyword=?", keyword).Error
}

func (u *sceneRepo) Page(ctx kratosx.Context, req *types.PageSceneRequest) ([]*biz.Scene, uint32, error) {
	var list []*biz.Scene
	var total int64
	db := ctx.DB().Model(biz.Scene{}).Preload("Agreements", "status=true")
	if req.Name != nil {
		db.Where("name like ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *sceneRepo) Create(ctx kratosx.Context, scene *biz.Scene) (uint32, error) {
	return scene.ID, ctx.DB().Create(scene).Error
}

func (u *sceneRepo) Update(ctx kratosx.Context, scene *biz.Scene) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(biz.AgreementScene{}, "scene_id=?", scene.ID).Error; err != nil {
			return err
		}
		return tx.Updates(scene).Error
	})
}

func (u *sceneRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Scene{}, "id=?", id).Error
}
