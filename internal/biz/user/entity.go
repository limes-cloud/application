package user

import (
	ktypes "github.com/limes-cloud/kratosx/types"

	appbiz "github.com/limes-cloud/user-center/internal/biz/app"
	"github.com/limes-cloud/user-center/internal/biz/channel"
)

type User struct {
	ktypes.BaseModel
	Phone       *string      `json:"phone" gorm:"unique;type:char(15);comment:手机"`
	Email       *string      `json:"email" gorm:"unique;binary;size:64;comment:邮箱"`
	Username    *string      `json:"username" gorm:"unique;binary;type:char(32);comment:账号"`
	Password    string       `json:"password" gorm:"size:256;comment:密码"`
	NickName    string       `json:"nick_name" gorm:"size:32;comment:昵称"`
	RealName    *string      `json:"real_name" gorm:"size:32;comment:真实姓名"`
	Avatar      string       `json:"avatar" gorm:"size:128;comment:头像"`
	Gender      string       `json:"gender" gorm:"default:U;type:enum('F','M','U');comment:昵称"`
	Status      *bool        `json:"status" gorm:"comment:状态"`
	DisableDesc *string      `json:"disable_desc" gorm:"size:128;comment:禁用原因"`
	From        string       `json:"from" gorm:"not null;size:128;comment:用户来源标识"`
	FromDesc    string       `json:"from_desc" gorm:"not null;size:128;comment:用户来源"`
	UserApps    []*UserApp   `json:"-" gorm:"constraint:onDelete:cascade"`
	Auths       []*Auth      `json:"-" gorm:"constraint:onDelete:cascade"`
	UserExtras  []*UserExtra `json:"-" gorm:"constraint:onDelete:cascade"`
}

func (t User) TableName() string {
	return "user"
}

type UserExtra struct {
	ktypes.CreateModel
	UserID  uint32 `json:"user_id" gorm:"uniqueIndex:uk;not null;comment:用户id"`
	Keyword string `json:"keyword" gorm:"uniqueIndex:uk;binary;not null;size:32;comment:关键字"`
	Value   string `json:"value" gorm:"not null;size:1024;comment:扩展值"`
}

func (t UserExtra) TableName() string {
	return "user_extra"
}

type UserApp struct {
	ktypes.CreateModel
	UserID  uint32      `json:"user_id" gorm:"uniqueIndex:ua;not null;comment:用户id"`
	AppID   uint32      `json:"app_id" gorm:"uniqueIndex:ua;not null;comment:应用id"`
	LoginAt int64       `json:"login_at" gorm:"comment:最近登录"`
	User    *User       `json:"user" gorm:"constraint:onDelete:cascade"`
	App     *appbiz.App `json:"app"` // 不允许直接删除app
}

func (t UserApp) TableName() string {
	return "user_app"
}

type Auth struct {
	ktypes.CreateModel
	UserID          uint32           `json:"user_id" gorm:"uniqueIndex:uc;not null;comment:用户id"`
	ChannelID       uint32           `json:"channel_id"  gorm:"uniqueIndex:uc;uniqueIndex:ua;not null;comment:渠道id"`
	AuthID          *string          `json:"auth_id" gorm:"uniqueIndex:ua;binary;size:64;comment:渠道授权ID"`
	UnionID         *string          `json:"union_id" gorm:"binary;size:64;comment:渠道联合ID"`
	ChannelToken    *string          `json:"channel_token" gorm:"size:64;comment:渠道token"`
	ChannelExpireAt int64            `json:"channel_expire_at" gorm:"comment:渠道token过期时间"`
	JwtToken        string           `json:"jwt_token" gorm:"size:1024;comment:平台Token"`
	JwtExpireAt     int64            `json:"jwt_expire_at" gorm:"comment:过期时间"`
	LoginAt         int64            `json:"login_at" gorm:"comment:最近登录时间"`
	User            *User            `json:"user" gorm:"constraint:onDelete:cascade"`
	Channel         *channel.Channel `json:"channel"` // 不允许直接删除channel
}

func (t Auth) TableName() string {
	return "auth"
}
