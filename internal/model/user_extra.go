package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type UserExtra struct {
	types.CreateModel
	UserID  uint32 `json:"user_id" gorm:"uniqueIndex:uk;not null;comment:用户id"`
	Keyword string `json:"keyword" gorm:"uniqueIndex:uk;not null;size:32;comment:关键字"`
	Value   string `json:"value" gorm:"not null;size:256;comment:扩展值"`
}

func (u *UserExtra) Create(ctx kratosx.Context) error {
	// read redis
	return ctx.DB().Create(u).Error
}

func (u *UserExtra) Delete(ctx kratosx.Context) error {
	return ctx.DB().Delete(u).Error
}
