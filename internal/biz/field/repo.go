package field

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// ListField 获取用户字段列表
	ListField(ctx kratosx.Context, req *ListFieldRequest) ([]*Field, uint32, error)

	// CreateField 创建用户字段
	CreateField(ctx kratosx.Context, req *Field) (uint32, error)

	// UpdateField 更新用户字段
	UpdateField(ctx kratosx.Context, req *Field) error

	// UpdateFieldStatus 更新用户字段状态
	UpdateFieldStatus(ctx kratosx.Context, id uint32, status bool) error

	// DeleteField 删除用户字段
	DeleteField(ctx kratosx.Context, ids []uint32) (uint32, error)
}
