package user

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// GetUser 获取指定的用户信息
	GetUser(ctx kratosx.Context, id uint32) (*User, error)

	// ListUser 获取用户信息列表
	ListUser(ctx kratosx.Context, req *ListUserRequest) ([]*User, uint32, error)

	// CreateUser 创建用户信息
	CreateUser(ctx kratosx.Context, req *User) (uint32, error)

	// ImportUser 导入用户信息
	ImportUser(ctx kratosx.Context, req []*User) (uint32, error)

	// ExportUser 导出用户信息
	ExportUser(ctx kratosx.Context, req *ExportUserRequest) (string, error)

	// UpdateUser 更新用户信息
	UpdateUser(ctx kratosx.Context, req *User) error

	// UpdateUserStatus 更新用户信息状态
	UpdateUserStatus(ctx kratosx.Context, req *UpdateUserStatusRequest) error

	// DeleteUser 删除用户信息
	DeleteUser(ctx kratosx.Context, ids []uint32) (uint32, error)

	// GetTrashUser 获取指定的回收站用户信息
	GetTrashUser(ctx kratosx.Context, id uint32) (*User, error)

	// ListTrashUser 获取回收站用户信息列表
	ListTrashUser(ctx kratosx.Context, req *ListTrashUserRequest) ([]*User, uint32, error)

	// DeleteTrashUser 彻底删除用户信息
	DeleteTrashUser(ctx kratosx.Context, ids []uint32) (uint32, error)

	// RevertTrashUser 还原用户信息
	RevertTrashUser(ctx kratosx.Context, id uint32) error

	// GetUserByPhone 获取指定的用户信息
	GetUserByPhone(ctx kratosx.Context, phone string) (*User, error)

	// GetUserByEmail 获取指定的用户信息
	GetUserByEmail(ctx kratosx.Context, email string) (*User, error)

	// GetUserByUsername 获取指定的用户信息
	GetUserByUsername(ctx kratosx.Context, username string) (*User, error)
}
