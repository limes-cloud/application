package userinfo

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// ListUserinfo 获取用户扩展信息列表
	ListUserinfo(ctx kratosx.Context, req *ListUserinfoRequest) ([]*Userinfo, uint32, error)

	// UpdateUserinfo 更新用户扩展信息
	UpdateUserinfo(ctx kratosx.Context, req *Userinfo) error

	// CheckKeywords 检查keyword是否合法
	CheckKeywords(ctx kratosx.Context, appId uint32, keywords []string) error

	// CreateUserinfo 创建用户扩展信息
	CreateUserinfo(ctx kratosx.Context, req *Userinfo) (uint32, error)

	// DeleteUserinfo 删除用户扩展信息
	DeleteUserinfo(ctx kratosx.Context, ids []uint32) (uint32, error)
}
