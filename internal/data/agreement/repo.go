package agreement

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/user-center/internal/biz/agreement"
)

type repo struct {
}

func NewRepo() biz.Repo {
	return &repo{}
}

func (u *repo) GetContent(ctx kratosx.Context, id uint32) (*biz.Content, error) {
	var content biz.Content
	return &content, ctx.DB().First(&content, "id=?", id).Error
}

func (u *repo) PageContent(ctx kratosx.Context, req *biz.PageContentRequest) ([]*biz.Content, uint32, error) {
	var list []*biz.Content
	var total int64
	db := ctx.DB().Select("id,created_at,updated_at,name,status,description").Model(biz.Content{})
	if req.Name != nil {
		db.Where("name like ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *repo) AddContent(ctx kratosx.Context, content *biz.Content) (uint32, error) {
	return content.ID, ctx.DB().Create(content).Error
}

func (u *repo) UpdateContent(ctx kratosx.Context, content *biz.Content) error {
	return ctx.DB().Updates(content).Error
}

func (u *repo) DeleteContent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Content{}, "id=?", id).Error
}

func (u *repo) GetSceneByKeyword(ctx kratosx.Context, keyword string) (*biz.Scene, error) {
	var scene biz.Scene
	return &scene, ctx.DB().Model(biz.Scene{}).Preload("Contents", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,created_at,name").Where("status=true")
	}).Find(&scene, "keyword=?", keyword).Error
}

func (u *repo) PageScene(ctx kratosx.Context, req *biz.PageSceneRequest) ([]*biz.Scene, uint32, error) {
	var list []*biz.Scene
	var total int64
	db := ctx.DB().Model(biz.Scene{}).Preload("Contents", "status=true")
	if req.Name != nil {
		db.Where("name like ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *repo) AddScene(ctx kratosx.Context, scene *biz.Scene) (uint32, error) {
	return scene.ID, ctx.DB().Create(scene).Error
}

func (u *repo) UpdateScene(ctx kratosx.Context, scene *biz.Scene) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(biz.SceneContent{}, "scene_id=?", scene.ID).Error; err != nil {
			return err
		}
		return tx.Updates(scene).Error
	})
}

func (u *repo) DeleteScene(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Scene{}, "id=?", id).Error
}
