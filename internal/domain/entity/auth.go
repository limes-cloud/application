package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type Auth struct {
	UserId      uint32  `json:"userId" gorm:"column:user_id"`
	AppId       uint32  `json:"appId" gorm:"column:app_id"`
	Status      *bool   `json:"status" gorm:"column:status"`
	DisableDesc *string `json:"disableDesc" gorm:"column:disable_desc"`
	Setting     *string `json:"setting" gorm:"column:setting"`
	Token       string  `json:"token" gorm:"column:token"`
	LoggedAt    int64   `json:"loggedAt" gorm:"column:logged_at"`
	ExpiredAt   int64   `json:"expiredAt" gorm:"column:expired_at"`
	App         *App    `json:"app"`
	types.CreateModel
}

func (a *Auth) GetDisableDescString() string {
	if a.DisableDesc == nil {
		return "用户异常"
	}
	return *a.DisableDesc
}
