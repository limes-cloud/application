package data

import (
	"fmt"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	biz "github.com/limes-cloud/usercenter/internal/biz/field"
	"github.com/limes-cloud/usercenter/internal/data/model"
	"google.golang.org/protobuf/proto"
)

type fieldRepo struct {
}

func NewFieldRepo() biz.Repo {
	return &fieldRepo{}
}

// ToFieldEntity model转entity
func (r fieldRepo) ToFieldEntity(m *model.Field) *biz.Field {
	e := &biz.Field{}
	_ = valx.Transform(m, e)
	return e
}

// ToFieldModel entity转model
func (r fieldRepo) ToFieldModel(e *biz.Field) *model.Field {
	m := &model.Field{}
	_ = valx.Transform(e, m)
	return m
}

// ListField 获取列表
func (r fieldRepo) ListField(ctx kratosx.Context, req *biz.ListFieldRequest) ([]*biz.Field, uint32, error) {
	var (
		bs    []*biz.Field
		ms    []*model.Field
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Field{}).Select(fs)

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
		bs = append(bs, r.ToFieldEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateField 创建数据
func (r fieldRepo) CreateField(ctx kratosx.Context, req *biz.Field) (uint32, error) {
	m := r.ToFieldModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// GetField 获取指定的数据
func (r fieldRepo) GetField(ctx kratosx.Context, id uint32) (*biz.Field, error) {
	var (
		m  = model.Field{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToFieldEntity(&m), nil
}

// UpdateField 更新数据
func (r fieldRepo) UpdateField(ctx kratosx.Context, req *biz.Field) error {
	return ctx.DB().Updates(r.ToFieldModel(req)).Error
}

// UpdateFieldStatus 更新数据状态
func (r fieldRepo) UpdateFieldStatus(ctx kratosx.Context, id uint32, status bool) error {
	return ctx.DB().Model(model.Field{}).Where("id=?", id).Update("status", status).Error
}

// DeleteField 删除数据
func (r fieldRepo) DeleteField(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.Field{})
	return uint32(db.RowsAffected), db.Error
}
