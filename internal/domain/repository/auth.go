package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/application/internal/domain/entity"
	"github.com/limes-cloud/application/internal/types"
)

type Auth interface {
	// ListAuth 获取应用授权信息列表
	ListAuth(ctx kratosx.Context, req *types.ListAuthRequest) ([]*entity.Auth, uint32, error)

	// CreateAuth 创建应用授权信息
	CreateAuth(ctx kratosx.Context, req *entity.Auth) (uint32, error)

	// UpsertAuth 更新应用授权信息
	UpsertAuth(ctx kratosx.Context, req *entity.Auth) error

	// UpdateAuthStatus 更新应用授权信息状态
	UpdateAuthStatus(ctx kratosx.Context, req *types.UpdateAuthStatusRequest) error

	// DeleteAuth 删除应用授权信息
	DeleteAuth(ctx kratosx.Context, userId uint32, appId uint32) error

	// GetAuth 获取指定的应用授权信息
	GetAuth(ctx kratosx.Context, id uint32) (*entity.Auth, error)

	// GetAuthByUA 获取指定的应用授权信息
	GetAuthByUA(ctx kratosx.Context, uid uint32, aid uint32) (*entity.Auth, error)

	// IsBindOAuth 判断渠道是否绑定用户
	IsBindOAuth(ctx kratosx.Context, cid uint32, aid string) bool

	// GetOAuthByCA 获取指定的授权信息
	GetOAuthByCA(ctx kratosx.Context, cid uint32, aid string) (*entity.OAuth, error)

	// ListOAuth 获取应用授权信息列表
	ListOAuth(ctx kratosx.Context, req *types.ListOAuthRequest) ([]*entity.OAuth, uint32, error)

	// CreateOAuth 创建应用授权信息
	CreateOAuth(ctx kratosx.Context, req *entity.OAuth) (string, error)

	// GetOAuthByUid 获取指定的三方授权应用
	GetOAuthByUid(ctx kratosx.Context, uid string) (*entity.OAuth, error)

	// UpdateOAuth 更新应用授权信息
	UpdateOAuth(ctx kratosx.Context, req *entity.OAuth) error

	// DeleteOAuth 删除应用授权信息
	DeleteOAuth(ctx kratosx.Context, userId uint32, appId uint32) error

	// BindOAuthByUid 通过授权渠道号绑定
	BindOAuthByUid(ctx kratosx.Context, uid uint32, aid string) error
}
