package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/biz/types"
	"github.com/limes-cloud/user-center/pkg/field"
)

type ExtraField struct {
	ktypes.BaseModel
	Keyword     string  `json:"keyword" gorm:"unique;not null;binary;size:32;comment:字段标识"`
	Type        string  `json:"type" gorm:"not null;size:32;comment:字段类型"`
	Name        string  `json:"name" gorm:"not null;size:32;comment:字段名称"`
	Description *string `json:"description" gorm:"size:128;comment:字段描述"`
}

type ExtraFieldRepo interface {
	Page(ctx kratosx.Context, req *types.PageExtraFieldRequest) ([]*ExtraField, uint32, error)
	All(ctx kratosx.Context) ([]*ExtraField, error)
	Create(ctx kratosx.Context, c *ExtraField) (uint32, error)
	Update(ctx kratosx.Context, c *ExtraField) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}

type ExtraFieldUseCase struct {
	config *config.Config
	repo   ExtraFieldRepo
}

func NewExtraFieldUseCase(config *config.Config, repo ExtraFieldRepo) *ExtraFieldUseCase {
	return &ExtraFieldUseCase{config: config, repo: repo}
}

// Types 获取全部字段类型
func (u *ExtraFieldUseCase) Types() []*types.ExtraFieldType {
	var list []*types.ExtraFieldType
	tps := field.New().GetTypes()

	for key, tp := range tps {
		list = append(list, &types.ExtraFieldType{Type: key, Name: tp.Name()})
	}
	return list
}

func (u *ExtraFieldUseCase) FiledTypeSet(ctx kratosx.Context) map[string]*ExtraField {
	m := map[string]*ExtraField{}
	efs, err := u.repo.All(ctx)
	if err != nil {
		return m
	}
	for _, ef := range efs {
		m[ef.Keyword] = ef
	}
	return m
}

// Page 获取全部登录字段
func (u *ExtraFieldUseCase) Page(ctx kratosx.Context, req *types.PageExtraFieldRequest) ([]*ExtraField, uint32, error) {
	ef, total, err := u.repo.Page(ctx, req)
	if err != nil {
		return nil, 0, v1.NotRecordError()
	}
	return ef, total, nil
}

// Add 添加登录字段信息
func (u *ExtraFieldUseCase) Add(ctx kratosx.Context, ef *ExtraField) (uint32, error) {
	id, err := u.repo.Create(ctx, ef)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新登录字段信息
func (u *ExtraFieldUseCase) Update(ctx kratosx.Context, ef *ExtraField) error {
	if err := u.repo.Update(ctx, ef); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除登录字段信息
func (u *ExtraFieldUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
