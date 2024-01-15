package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/user-center/pkg/util"
)

// User 用户表
type User struct {
	types.BaseModel
	Phone       string       `json:"phone" gorm:"unique;type:char(15);comment:手机"`
	Email       string       `json:"email" gorm:"unique;size:64;comment:邮箱"`
	Username    string       `json:"username" gorm:"unique;type:char(32);comment:账号"`
	Password    string       `json:"password" gorm:"size:256;comment:密码"`
	NickName    string       `json:"nick_name" gorm:"size:32;comment:昵称"`
	RealName    string       `json:"real_name" gorm:"size:32;comment:真实姓名"`
	IdCard      string       `json:"id_card" gorm:"unique;size:32;comment:身份证号"`
	Avatar      string       `json:"avatar" gorm:"size:128;comment:头像"`
	Gender      string       `json:"gender" gorm:"default:U;type:enum('F','M','U');comment:昵称"`
	Status      *bool        `json:"status" gorm:"comment:状态"`
	From        string       `json:"from" gorm:"size:128;comment:用户来源"`
	DisableDesc *string      `json:"disable_desc" gorm:"size:128;comment:禁用原因"`
	UserApps    []*UserApp   `json:"user_apps,omitempty" gorm:"constraint:onDelete:cascade"`
	UserExtras  []*UserExtra `json:"user_extras,omitempty" gorm:"constraint:onDelete:cascade"`
}

func (u *User) Create(ctx kratosx.Context) error {
	if u.Password != "" {
		u.Password = util.ParsePwd(u.Password)
	}

	return ctx.DB().Create(u).Error
}

func (u *User) Import(ctx kratosx.Context, list []*User) error {
	return ctx.DB().Create(&list).Error
}

func (u *User) FindByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(u, "id=?", id).Error
}

func (u *User) FindByPhone(ctx kratosx.Context, phone string) error {
	return ctx.DB().First(u, "phone=?", phone).Error
}

func (u *User) FindByEmail(ctx kratosx.Context, email string) error {
	return ctx.DB().First(u, "email=?", email).Error
}

func (u *User) FindByUsername(ctx kratosx.Context, un string) error {
	return ctx.DB().First(u, "username=?", un).Error
}

// Page 查询分页数据
func (u *User) Page(ctx kratosx.Context, options *types.PageOptions) ([]*User, uint32, error) {
	var list []*User
	total := int64(0)

	db := ctx.DB().Model(u).Preload("UserApps").Preload("UserExtras")
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// Update 更新用户信息
func (u *User) Update(ctx kratosx.Context) error {
	if u.Password != "" {
		u.Password = util.ParsePwd(u.Password)
	}
	return ctx.DB().Updates(&u).Error
}

func (u *User) DeleteById(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(User{}, "id=?", id).Error
}
