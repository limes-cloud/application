package model

import (
	"github.com/limes-cloud/kratosx/types"
)

type Channel struct {
	Logo    string  `json:"logo" gorm:"column:logo"`
	Keyword string  `json:"keyword" gorm:"column:keyword"`
	Name    string  `json:"name" gorm:"column:name"`
	Status  *bool   `json:"status" gorm:"column:status"`
	Ak      *string `json:"ak" gorm:"column:ak"`
	Sk      *string `json:"sk" gorm:"column:sk"`
	Extra   *string `json:"extra" gorm:"column:extra"`
	types.BaseModel
}
