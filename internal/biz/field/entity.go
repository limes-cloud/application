package field

import ktypes "github.com/limes-cloud/kratosx/types"

type Field struct {
	ktypes.BaseModel
	Keyword     string  `json:"keyword" gorm:"unique;not null;binary;size:32;comment:字段标识"`
	Type        string  `json:"type" gorm:"not null;size:32;comment:字段类型"`
	Name        string  `json:"name" gorm:"not null;size:32;comment:字段名称"`
	Description *string `json:"description" gorm:"size:128;comment:字段描述"`
}
