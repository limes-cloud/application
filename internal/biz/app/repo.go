package app

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// GetApp 获取指定的应用信息
	GetApp(ctx kratosx.Context, id uint32) (*App, error)

	// ListApp 获取应用信息列表
	ListApp(ctx kratosx.Context, req *ListAppRequest) ([]*App, uint32, error)

	// CreateApp 创建应用信息
	CreateApp(ctx kratosx.Context, req *App) (uint32, error)

	// UpdateApp 更新应用信息
	UpdateApp(ctx kratosx.Context, req *App) error

	// UpdateAppStatus 更新应用信息状态
	UpdateAppStatus(ctx kratosx.Context, req *UpdateAppStatusRequest) error

	// DeleteApp 删除应用信息
	DeleteApp(ctx kratosx.Context, ids []uint32) (uint32, error)

	// GetAppByKeyword 获取指定的应用信息
	GetAppByKeyword(ctx kratosx.Context, keyword string) (*App, error)
}
