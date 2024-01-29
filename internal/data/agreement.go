package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/internal/biz/types"
	"gorm.io/gorm"
)

type agreementRepo struct {
}

func NewAgreementRepo() biz.AgreementRepo {
	return &agreementRepo{}
}

func (u *agreementRepo) Get(ctx kratosx.Context, id uint32) (*biz.Agreement, error) {
	var agreement biz.Agreement
	return &agreement, ctx.DB().First(&agreement, "id=?", id).Error
}

func (u *agreementRepo) Page(ctx kratosx.Context, req *types.PageAgreementRequest) ([]*biz.Agreement, uint32, error) {
	var list []*biz.Agreement
	var total int64
	db := ctx.DB().Select("id,created_at,updated_at,name,status,description").Model(biz.Agreement{})
	if req.Name != nil {
		db.Where("name like ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *agreementRepo) Create(ctx kratosx.Context, agreement *biz.Agreement) (uint32, error) {
	return agreement.ID, ctx.DB().Create(agreement).Error
}

func (u *agreementRepo) Update(ctx kratosx.Context, agreement *biz.Agreement) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		return tx.Updates(agreement).Error
	})
}

func (u *agreementRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Agreement{}, "id=?", id).Error
}
