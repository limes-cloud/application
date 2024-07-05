package model

import (
	"github.com/limes-cloud/kratosx/types"
)

type FeedbackCategory struct {
	Name string `json:"name" gorm:"column:name"`
	types.CreateModel
}
