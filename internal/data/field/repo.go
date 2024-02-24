package field

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/user-center/internal/biz/field"
)

type repo struct {
}

func NewRepo() biz.Repo {
	return &repo{}
}

func (u *repo) All(ctx kratosx.Context) ([]*biz.Field, error) {
	var list []*biz.Field
	return list, ctx.DB().Find(&list).Error
}

func (u *repo) Page(ctx kratosx.Context, req *biz.PageFieldRequest) ([]*biz.Field, uint32, error) {
	var list []*biz.Field
	var total int64
	db := ctx.DB().Model(biz.Field{})
	if req.Keyword != nil {
		db.Where("keyword=?", *req.Keyword)
	}
	if req.Name != nil {
		db.Where("name like ?", *req.Keyword+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *repo) Create(ctx kratosx.Context, field *biz.Field) (uint32, error) {
	return field.ID, ctx.DB().Create(field).Error
}

func (u *repo) Update(ctx kratosx.Context, field *biz.Field) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		return tx.Updates(field).Error
	})
}

func (u *repo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Field{}, "id=?", id).Error
}
