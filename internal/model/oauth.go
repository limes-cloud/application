package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

// OAuth 三方授权表
type OAuth struct {
	types.CreateModel
	UserID   uint32 `json:"user_id" gorm:"not null;comment:用户id"`
	Platform string `json:"platform" gorm:"not null;type:char(15);comment:平台标识"`
	AuthID   string `json:"auth_id" gorm:"not null;size:64;comment:授权ID"`
	UnionID  string `json:"union_id" gorm:"size:64;comment:平台联合ID"`
	Token    string `json:"token" gorm:"size:64;comment:平台Token"`
	ExpireAt int64  `json:"expire_at" gorm:"comment:过期时间"`
}

func (u *OAuth) Create(ctx kratosx.Context) error {
	// read redis
	return ctx.DB().Create(u).Error
}

func (u *OAuth) Delete(ctx kratosx.Context) error {
	return ctx.DB().Delete(u).Error
}
