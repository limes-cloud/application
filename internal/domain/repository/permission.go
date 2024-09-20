package repository

import "github.com/limes-cloud/kratosx"

type Permission interface {
	// GetPermission 获取当前用户，指定key的权限
	GetPermission(ctx kratosx.Context, keyword string) (bool, []uint32, error)

	// GetApp 获取当前用户对于app的权限
	GetApp(ctx kratosx.Context) (bool, []uint32, error)

	// HasApp 获取当前用户是否具有指定app的权限
	HasApp(ctx kratosx.Context, id uint32) bool
}
