package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type User struct {
	Phone       *string `json:"phone" gorm:"column:phone"`
	Email       *string `json:"email" gorm:"column:email"`
	Username    *string `json:"username" gorm:"column:username"`
	Password    *string `json:"password" gorm:"column:password"`
	NickName    string  `json:"nickName" gorm:"column:nick_name"`
	RealName    *string `json:"realName" gorm:"column:real_name"`
	Avatar      *string `json:"avatar" gorm:"column:avatar"`
	AvatarUrl   *string `json:"avatarUrl" gorm:"-"`
	Gender      *string `json:"gender" gorm:"column:gender"`
	Status      *bool   `json:"status" gorm:"column:status"`
	DisableDesc *string `json:"disableDesc" gorm:"column:disable_desc"`
	From        string  `json:"from" gorm:"column:from"`
	FromDesc    string  `json:"fromDesc" gorm:"column:from_desc"`
	Auths       []*Auth `json:"-" gorm:"foreignKey:user_id;references:id"`
	types.DeleteModel
}
