package channel

import (
	ktypes "github.com/limes-cloud/kratosx/types"
)

type Channel struct {
	ktypes.CreateModel
	Logo     string `json:"logo"`
	Platform string `json:"platform"`
	Name     string `json:"name"`
	Ak       string `json:"ak"`
	Sk       string `json:"sk"`
	Extra    string `json:"extra"`
	Status   *bool  `json:"status"`
}

func (t Channel) TableName() string {
	return "channel"
}

type Typer struct {
	Platform string `json:"platform"`
	Name     string
}
