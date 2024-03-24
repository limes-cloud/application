package user

import (
	ktypes "github.com/limes-cloud/kratosx/types"

	appbiz "github.com/limes-cloud/user-center/internal/biz/app"
	"github.com/limes-cloud/user-center/internal/biz/channel"
)

type User struct {
	ktypes.BaseModel
	Phone       *string      `json:"phone"`
	Email       *string      `json:"email"`
	Username    *string      `json:"username"`
	Password    string       `json:"password"`
	NickName    string       `json:"nick_name"`
	RealName    *string      `json:"real_name"`
	Avatar      string       `json:"avatar"`
	Gender      string       `json:"gender"`
	Status      *bool        `json:"status"`
	DisableDesc *string      `json:"disable_desc"`
	From        string       `json:"from"`
	FromDesc    string       `json:"from_desc"`
	UserApps    []*UserApp   `json:"-" gorm:"foreignKey:user_id;references:id"`
	Auths       []*Auth      `json:"-" gorm:"foreignKey:user_id;references:id"`
	UserExtras  []*UserExtra `json:"-" gorm:"foreignKey:user_id;references:id"`
}

func (t User) TableName() string {
	return "user"
}

type UserExtra struct {
	ktypes.CreateModel
	UserID  uint32 `json:"user_id"`
	Keyword string `json:"keyword"`
	Value   string `json:"value"`
}

func (t UserExtra) TableName() string {
	return "user_extra"
}

type UserApp struct {
	ktypes.CreateModel
	UserID  uint32      `json:"user_id"`
	AppID   uint32      `json:"app_id"`
	LoginAt int64       `json:"login_at"`
	App     *appbiz.App `json:"app" gorm:"foreignKey:app_id;references:id"` // 不允许直接删除app
}

func (t UserApp) TableName() string {
	return "user_app"
}

type Auth struct {
	ktypes.CreateModel
	UserID          uint32           `json:"user_id"`
	ChannelID       uint32           `json:"channel_id"`
	AuthID          *string          `json:"auth_id"`
	UnionID         *string          `json:"union_id"`
	ChannelToken    *string          `json:"channel_token"`
	ChannelExpireAt int64            `json:"channel_expire_at"`
	JwtToken        string           `json:"jwt_token"`
	JwtExpireAt     int64            `json:"jwt_expire_at"`
	LoginAt         int64            `json:"login_at"`
	User            *User            `json:"user" gorm:"foreignKey:user_id;references:id"`
	Channel         *channel.Channel `json:"channel" gorm:"foreignKey:channel_id;references:id"` // 不允许直接删除channel
}

func (t Auth) TableName() string {
	return "auth"
}
