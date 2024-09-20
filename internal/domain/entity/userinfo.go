package entity

import (
	"github.com/limes-cloud/kratosx/types"
)

type Userinfo struct {
	UserId  uint32 `json:"userId" gorm:"column:user_id"`
	FieldId uint32 `json:"field_id" gorm:"column:field_id"`
	Keyword string `json:"keyword" gorm:"column:keyword"`
	Value   string `json:"value" gorm:"column:value"`
	Field   *Field `json:"field"`
	types.BaseModel
}
