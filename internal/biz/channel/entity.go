package channel

import ktypes "github.com/limes-cloud/kratosx/types"

type Channel struct {
	ktypes.CreateModel
	Logo     string `json:"logo" gorm:"not null;size:128;comment:渠道logo"`
	Platform string `json:"platform" gorm:"unique;not null;binary;type:char(32);comment:渠道标识"`
	Name     string `json:"name" gorm:"not null;size:32;comment:渠道名称"`
	Ak       string `json:"ak" gorm:"size:32;comment:渠道ak"`
	Sk       string `json:"sk" gorm:"size:32;comment:渠道sk"`
	Extra    string `json:"extra" gorm:"size:256;comment:渠道状态"`
	Status   *bool  `json:"status" gorm:"not null;comment:渠道状态"`
}

type Typer struct {
	Platform string `json:"platform"`
	Name     string
}

func (t Channel) TableName() string {
	return "channel"
}
