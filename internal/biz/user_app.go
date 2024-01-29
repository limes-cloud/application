package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
)

type UserApp struct {
	ktypes.CreateModel
	UserID  uint32 `json:"user_id" gorm:"uniqueIndex:ua;not null;comment:用户id"`
	AppID   uint32 `json:"app_id" gorm:"uniqueIndex:ua;not null;comment:渠道id"`
	LoginAt int64  `json:"login_at" gorm:"comment:最近登录"`
	User    *User  `json:"user" gorm:"constraint:onDelete:cascade"`
	App     *App   `json:"app"` // 不允许直接删除app
}

type UserAppRepo interface {
	Create(ctx kratosx.Context, uid, aid uint32) (uint32, error)
	Delete(ctx kratosx.Context, uid, aid uint32) error
}

type UserAppUseCase struct {
	config *config.Config
	repo   UserAppRepo
}

func NewUserAppUseCase(config *config.Config, repo UserAppRepo) *UserAppUseCase {
	return &UserAppUseCase{config: config, repo: repo}
}

// Add 添加用户渠道信息
func (u *UserAppUseCase) Add(ctx kratosx.Context, uid, aid uint32) (uint32, error) {
	id, err := u.repo.Create(ctx, uid, aid)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Delete 删除用户渠道信息
func (u *UserAppUseCase) Delete(ctx kratosx.Context, uid, aid uint32) error {
	if err := u.repo.Delete(ctx, uid, aid); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
