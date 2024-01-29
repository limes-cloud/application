package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/biz/types"
	"google.golang.org/protobuf/proto"
)

type User struct {
	ktypes.BaseModel
	Phone        *string        `json:"phone" gorm:"unique;type:char(15);comment:手机"`
	Email        *string        `json:"email" gorm:"unique;binary;size:64;comment:邮箱"`
	Username     *string        `json:"username" gorm:"unique;binary;type:char(32);comment:账号"`
	Password     string         `json:"password" gorm:"size:256;comment:密码"`
	NickName     string         `json:"nick_name" gorm:"size:32;comment:昵称"`
	RealName     string         `json:"real_name" gorm:"size:32;comment:真实姓名"`
	Avatar       string         `json:"avatar" gorm:"size:128;comment:头像"`
	Gender       string         `json:"gender" gorm:"default:U;type:enum('F','M','U');comment:昵称"`
	Status       *bool          `json:"status" gorm:"comment:状态"`
	DisableDesc  *string        `json:"disable_desc" gorm:"size:128;comment:禁用原因"`
	From         string         `json:"from" gorm:"not null;size:128;comment:用户来源标识"`
	FromDesc     string         `json:"from_desc" gorm:"not null;size:128;comment:用户来源"`
	Apps         []*App         `json:"-" gorm:"many2many:user_app;-:migration"`
	UserApps     []*UserApp     `json:"-" gorm:"-:migration"`
	Channels     []*Channel     `json:"-" gorm:"many2many:user_channel;-:migration"`
	UserChannels []*UserChannel `json:"-" gorm:"-:migration"`
	UserExtras   []*UserExtra   `json:"-" gorm:"constraint:onDelete:cascade"`
}

type UserExtra struct {
	ktypes.CreateModel
	UserID  uint32 `json:"user_id" gorm:"uniqueIndex:uk;not null;comment:用户id"`
	Keyword string `json:"keyword" gorm:"uniqueIndex:uk;binary;not null;size:32;comment:关键字"`
	Value   string `json:"value" gorm:"not null;size:1024;comment:扩展值"`
}

type UserRepo interface {
	Create(ctx kratosx.Context, user *User) (uint32, error)
	Import(ctx kratosx.Context, list []*User) error
	Get(ctx kratosx.Context, id uint32) (*User, error)
	GetByPhone(ctx kratosx.Context, phone string) (*User, error)
	GetByEmail(ctx kratosx.Context, email string) (*User, error)
	GetByUsername(ctx kratosx.Context, un string) (*User, error)
	GetByIdCard(ctx kratosx.Context, idCard string) (*User, error)
	PageUser(ctx kratosx.Context, req *types.PageUserRequest) ([]*User, uint32, error)
	Update(ctx kratosx.Context, user *User) error
	UpdateUserApp(ctx kratosx.Context, user *UserApp) error
	Delete(ctx kratosx.Context, id uint32) error
	GetAllToken(ctx kratosx.Context, id uint32) []string
}

type UserUseCase struct {
	config *config.Config
	repo   UserRepo
}

func NewUserUseCase(config *config.Config, repo UserRepo) *UserUseCase {
	return &UserUseCase{config: config, repo: repo}
}

// Get 获取用户信息
func (u *UserUseCase) Get(ctx kratosx.Context, id uint32) (*User, error) {
	user, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return user, nil
}

// GetByEmail 获取用户信息
func (u *UserUseCase) GetByEmail(ctx kratosx.Context, email string) (*User, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return user, nil
}

// GetByPhone 获取用户信息
func (u *UserUseCase) GetByPhone(ctx kratosx.Context, phone string) (*User, error) {
	user, err := u.repo.GetByPhone(ctx, phone)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return user, nil
}

// GetByUsername 获取用户信息
func (u *UserUseCase) GetByUsername(ctx kratosx.Context, un string) (*User, error) {
	user, err := u.repo.GetByUsername(ctx, un)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return user, nil
}

// GetByIdCard 获取用户信息
func (u *UserUseCase) GetByIdCard(ctx kratosx.Context, idCard string) (*User, error) {
	user, err := u.repo.GetByIdCard(ctx, idCard)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return user, nil
}

// Page 分页获取用户
func (u *UserUseCase) Page(ctx kratosx.Context, req *types.PageUserRequest) ([]*User, uint32, error) {
	list, total, err := u.repo.PageUser(ctx, req)
	if err != nil {
		return nil, 0, v1.DatabaseErrorFormat(err.Error())
	}
	return list, total, err
}

// Add 添加用户信息
func (u *UserUseCase) Add(ctx kratosx.Context, user *User) (uint32, error) {
	user.Status = proto.Bool(true)
	if user.NickName == "" {
		user.NickName = user.RealName
	}

	id, err := u.repo.Create(ctx, user)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Import 导入用户信息
func (u *UserUseCase) Import(ctx kratosx.Context, users []*User) error {
	if err := u.repo.Import(ctx, users); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Update 添加用户信息
func (u *UserUseCase) Update(ctx kratosx.Context, user *User) error {
	if err := u.repo.Update(ctx, user); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除用户信息
func (u *UserUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Enable 启用用户
func (u *UserUseCase) Enable(ctx kratosx.Context, id uint32) error {
	user := &User{
		BaseModel:   ktypes.BaseModel{ID: id},
		Status:      proto.Bool(true),
		DisableDesc: proto.String(""),
	}
	if err := u.repo.Update(ctx, user); err != nil {
		return v1.DatabaseError()
	}
	return nil
}

// Disable 禁用用户
func (u *UserUseCase) Disable(ctx kratosx.Context, id uint32, desc string) error {
	user := &User{
		BaseModel:   ktypes.BaseModel{ID: id},
		Status:      proto.Bool(false),
		DisableDesc: proto.String(desc),
	}
	if err := u.repo.Update(ctx, user); err != nil {
		return v1.DatabaseError()
	}
	return nil
}

// Offline 对当前登陆用户进行下线
func (u *UserUseCase) Offline(ctx kratosx.Context, id uint32) error {
	tokens := u.repo.GetAllToken(ctx, id)

	//将用户下线
	for _, token := range tokens {
		ctx.JWT().AddBlacklist(token)

	}
	return nil
}
