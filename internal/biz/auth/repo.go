package auth

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// ListAuth 获取应用授权信息列表
	ListAuth(ctx kratosx.Context, req *ListAuthRequest) ([]*Auth, uint32, error)

	// CreateAuth 创建应用授权信息
	CreateAuth(ctx kratosx.Context, req *Auth) (uint32, error)

	// UpdateAuth 更新应用授权信息
	UpdateAuth(ctx kratosx.Context, req *Auth) error

	// UpdateAuthStatus 更新应用授权信息状态
	UpdateAuthStatus(ctx kratosx.Context, req *UpdateAuthStatusRequest) error

	// DeleteAuth 删除应用授权信息
	DeleteAuth(ctx kratosx.Context, userId uint32, appId uint32) error

	// GetAuthByUserIdAndAppId 获取指定的应用授权信息
	GetAuthByUserIdAndAppId(ctx kratosx.Context, userId uint32, appId uint32) (*Auth, error)

	// HasUserByEmail 是否存在指定的用户邮箱信息
	HasUserByEmail(ctx kratosx.Context, email string) bool

	// HasUserByUsername 是否存在指定的用户账户信息
	HasUserByUsername(ctx kratosx.Context, un string) bool

	// IsBindUser 判断渠道是否绑定用户
	IsBindUser(ctx kratosx.Context, cid uint32, aid string) bool

	// HasAppScope 获取指定用户信息
	HasAppScope(ctx kratosx.Context, aid uint32, uid uint32) error

	// GetUserByCA 获取指定用户信息
	GetUserByCA(ctx kratosx.Context, cid uint32, aid string) (*User, error)

	// GetUserByEmail 获取指定用户信息
	GetUserByEmail(ctx kratosx.Context, email string) (*User, error)

	// GetUserByUsername 获取指定用户信息
	GetUserByUsername(ctx kratosx.Context, un string) (*User, error)

	// GetOAuthByCA 获取指定的授权信息
	GetOAuthByCA(ctx kratosx.Context, cid uint32, aid string) (*OAuth, error)

	// ListOAuth 获取应用授权信息列表
	ListOAuth(ctx kratosx.Context, req *ListOAuthRequest) ([]*OAuth, uint32, error)

	// CreateOAuth 创建应用授权信息
	CreateOAuth(ctx kratosx.Context, req *OAuth) (string, error)

	// GetOAuthByUid 获取指定的三方授权应用
	GetOAuthByUid(ctx kratosx.Context, uid string) (*OAuth, error)

	// UpdateOAuth 更新应用授权信息
	UpdateOAuth(ctx kratosx.Context, req *OAuth) error

	// DeleteOAuth 删除应用授权信息
	DeleteOAuth(ctx kratosx.Context, userId uint32, appId uint32) error

	// GetApp 获取指定应用的信息
	GetApp(ctx kratosx.Context, app string) (*App, error)

	// GetAppChannel 获取指定应用的授权渠道
	GetAppChannel(ctx kratosx.Context, app, channel string) (*Channel, error)

	// Register 注册用户
	Register(ctx kratosx.Context, user *User) error

	// BindOAuthByUid 通过授权渠道号绑定
	BindOAuthByUid(ctx kratosx.Context, uid uint32, aid string) error
}
