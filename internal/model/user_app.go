package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type UserApp struct {
	types.BaseModel
	UserID    uint32 `json:"user_id" gorm:"not null;comment:用户id"`
	App       string `json:"app" gorm:"not null;comment:应用标识"`
	Status    *bool  `json:"status"  gorm:"default:true;comment:使用状态"`
	LastLogin int64  `json:"last_login,omitempty" gorm:"comment:最后登录应用时间"`
	Token     string `json:"token,omitempty" gorm:"size:256;comment:登录应用Token"`
	ExpireAt  int64  `json:"expire_at,omitempty" gorm:"comment:过期时间"`
}

func (u *UserApp) Create(ctx kratosx.Context) error {
	// read redis
	return ctx.DB().Create(u).Error
}

func (u *UserApp) FindByUserId(ctx kratosx.Context, uid uint32) ([]*UserApp, error) {
	var list []*UserApp
	return list, ctx.DB().Find(&u, "user_id=?", uid).Error
}

func (u *UserApp) Delete(ctx kratosx.Context) error {
	return ctx.DB().Delete(u).Error
}
