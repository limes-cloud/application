package data

import (
	"strings"

	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/internal/biz/types"
	"github.com/limes-cloud/user-center/pkg/util"
)

type fieldRepo struct {
}

func NewExtraFieldRepo() biz.ExtraFieldRepo {
	return &fieldRepo{}
}

func (u *fieldRepo) All(ctx kratosx.Context) ([]*biz.ExtraField, error) {
	var list []*biz.ExtraField
	return list, ctx.DB().Find(&list).Error
}

func (u *fieldRepo) AllAppField(ctx kratosx.Context, appid uint32) ([]*biz.ExtraField, error) {
	var list []*biz.ExtraField
	if err := ctx.DB().Model(biz.ExtraField{}).Find(&list).Error; err != nil {
		return nil, err
	}

	app := biz.App{}
	if err := ctx.DB().First(&app, "id=?", appid).Error; err != nil {
		return nil, err
	}
	fields := strings.Split(app.UserFields, ",")

	var arr []*biz.ExtraField
	for _, item := range list {
		if !util.InList(fields, item.Keyword) {
			continue
		}
		arr = append(arr, item)
	}

	return arr, nil
}

func (u *fieldRepo) Page(ctx kratosx.Context, req *types.PageExtraFieldRequest) ([]*biz.ExtraField, uint32, error) {
	var list []*biz.ExtraField
	var total int64
	db := ctx.DB().Model(biz.ExtraField{})
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

func (u *fieldRepo) Create(ctx kratosx.Context, field *biz.ExtraField) (uint32, error) {
	return field.ID, ctx.DB().Create(field).Error
}

func (u *fieldRepo) Update(ctx kratosx.Context, field *biz.ExtraField) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		return tx.Updates(field).Error
	})
}

func (u *fieldRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.ExtraField{}, "id=?", id).Error
}
