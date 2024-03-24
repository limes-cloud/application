package app

import (
	ktypes "github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/user-center/internal/biz/channel"
	"github.com/limes-cloud/user-center/internal/biz/field"
)

type App struct {
	ktypes.BaseModel
	Keyword       string             `json:"keyword"`
	Logo          string             `json:"logo"`
	Name          string             `json:"name"`
	Status        *bool              `json:"status"`
	Version       string             `json:"version"`
	Copyright     string             `json:"copyright"`
	AllowRegistry *bool              `json:"allow_registry"`
	Description   string             `json:"description"`
	Fields        []*field.Field     `json:"fields" gorm:"many2many:app_field"`
	Channels      []*channel.Channel `json:"channels" gorm:"many2many:app_channel"`
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
