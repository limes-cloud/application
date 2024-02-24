package user

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/user-center/internal/biz/app"
)

type Repo interface {
	Add(ctx kratosx.Context, user *User) (uint32, error)
	Import(ctx kratosx.Context, list []*User) error
	Get(ctx kratosx.Context, id uint32) (*User, error)
	GetBase(ctx kratosx.Context, id uint32) (*User, error)
	GetByPhone(ctx kratosx.Context, phone string) (*User, error)
	GetByEmail(ctx kratosx.Context, email string) (*User, error)
	GetByUsername(ctx kratosx.Context, un string) (*User, error)
	PageUser(ctx kratosx.Context, req *PageUserRequest) ([]*User, uint32, error)
	Update(ctx kratosx.Context, user *User) error
	Delete(ctx kratosx.Context, id uint32) error
	GetJwtTokens(ctx kratosx.Context, id uint32) []string
	HasUsername(ctx kratosx.Context, un string) bool
	HasUserEmail(ctx kratosx.Context, email string) bool
	HasUserPhone(ctx kratosx.Context, phone string) bool

	GetApp(ctx kratosx.Context, keyword string) (*app.App, error)
	AddUserApp(ctx kratosx.Context, uid, aid uint32) (uint32, error)
	UpdateUserApp(ctx kratosx.Context, in *UserApp) error
	DeleteUserApp(ctx kratosx.Context, uid, aid uint32) error

	GetAuthByCU(ctx kratosx.Context, cid, uid uint32) (*Auth, error)
	GetAuthByCA(ctx kratosx.Context, cid uint32, aid string) (*Auth, error)
	AddAuth(ctx kratosx.Context, channel *Auth) (uint32, error)
	UpdateAuthByCU(ctx kratosx.Context, channel *Auth) error
}
