package app

import (
	ktypes "github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/user-center/internal/biz/channel"
	"github.com/limes-cloud/user-center/internal/biz/field"
)

type App struct {
	ktypes.BaseModel
	Keyword       string             `json:"keyword" gorm:"unique;not null;size:32;comment:应用标识"`
	Logo          string             `json:"logo" gorm:"not null;size:128;comment:应用logo"`
	Name          string             `json:"name" gorm:"not null;size:32;comment:应用名称"`
	Status        *bool              `json:"status" gorm:"not null;comment:应用状态"`
	Version       string             `json:"version" gorm:"size:32;comment:应用版本"`
	Copyright     string             `json:"copyright" gorm:"size:128;comment:应用版权"`
	AllowRegistry *bool              `json:"allow_registry" gorm:"not null;comment:是否允许注册"`
	Description   string             `json:"description" gorm:"size:128;comment:应用描述"`
	Fields        []*field.Field     `json:"fields" gorm:"many2many:app_field;constraint:onDelete:cascade"`
	Channels      []*channel.Channel `json:"channels" gorm:"many2many:app_channel;constraint:onDelete:cascade"`
	AppChannels   []*AppChannel      `json:"app_channels"`
	AppFields     []*AppField        `json:"app_fields"`
}

func (t App) TableName() string {
	return "app"
}

type AppChannel struct {
	AppID     uint32 `json:"app_id"`
	ChannelID uint32 `json:"channel_id"`
}

func (t AppChannel) TableName() string {
	return "app_channel"
}

type AppField struct {
	AppID   uint32 `json:"app_id"`
	FieldID uint32 `json:"field_id"`
}

func (t AppField) TableName() string {
	return "app_field"
}
