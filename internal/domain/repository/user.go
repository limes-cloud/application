package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/application/internal/domain/entity"
	"github.com/limes-cloud/application/internal/types"
)

type User interface {
	// GetUser 获取指定的用户信息
	GetUser(ctx kratosx.Context, id uint32) (*entity.User, error)

	// ListUser 获取用户信息列表
	ListUser(ctx kratosx.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error)

	// CreateUser 创建用户信息
	CreateUser(ctx kratosx.Context, req *entity.User) (uint32, error)

	// ImportUser 导入用户信息
	ImportUser(ctx kratosx.Context, req []*entity.User) (uint32, error)

	// ExportUser 导出用户信息
	ExportUser(ctx kratosx.Context, req *types.ExportUserRequest) (string, error)

	// UpdateUser 更新用户信息
	UpdateUser(ctx kratosx.Context, req *entity.User) error

	// UpdateUserStatus 更新用户信息状态
	UpdateUserStatus(ctx kratosx.Context, req *types.UpdateUserStatusRequest) error

	// DeleteUser 删除用户信息
	DeleteUser(ctx kratosx.Context, ids []uint32) (uint32, error)

	// GetTrashUser 获取指定的回收站用户信息
	GetTrashUser(ctx kratosx.Context, id uint32) (*entity.User, error)

	// ListTrashUser 获取回收站用户信息列表
	ListTrashUser(ctx kratosx.Context, req *types.ListTrashUserRequest) ([]*entity.User, uint32, error)

	// DeleteTrashUser 彻底删除用户信息
	DeleteTrashUser(ctx kratosx.Context, ids []uint32) (uint32, error)

	// RevertTrashUser 还原用户信息
	RevertTrashUser(ctx kratosx.Context, id uint32) error

	// GetUserByPhone 获取指定的用户信息
	GetUserByPhone(ctx kratosx.Context, phone string) (*entity.User, error)

	// GetUserByEmail 获取指定的用户信息
	GetUserByEmail(ctx kratosx.Context, email string) (*entity.User, error)

	// GetUserByUsername 获取指定的用户信息
	GetUserByUsername(ctx kratosx.Context, username string) (*entity.User, error)

	// HasUserByEmail 是否存在指定的用户邮箱信息
	HasUserByEmail(ctx kratosx.Context, email string) bool

	// HasUserByUsername 是否存在指定的用户账户信息
	HasUserByUsername(ctx kratosx.Context, un string) bool
}
