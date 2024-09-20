package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type OAuth struct {
	UserId    *uint32  `json:"userId" gorm:"column:user_id"`
	ChannelId uint32   `json:"channelId" gorm:"column:channel_id"`
	AuthId    string   `json:"authId" gorm:"column:auth_id"`
	UnionId   *string  `json:"unionId" gorm:"column:union_id"`
	Token     string   `json:"token" gorm:"column:token"`
	LoggedAt  int64    `json:"loggedAt" gorm:"column:logged_at"`
	ExpiredAt int64    `json:"expiredAt" gorm:"column:expired_at"`
	Channel   *Channel `json:"channel"`
	types.CreateModel
}

func (OAuth) TableName() string {
	return "oauth"
}
