package model

import (
	"github.com/limes-cloud/kratosx/types"
)

type Userinfo struct {
	UserId  uint32 `json:"userId" gorm:"column:user_id"`
	Keyword string `json:"keyword" gorm:"column:keyword"`
	Value   string `json:"value" gorm:"column:value"`
	Field   *Field `json:"field" gorm:"foreignKey:keyword;references:keyword"`
	types.BaseModel
}
