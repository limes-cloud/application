package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
)

type UserChannel struct {
	ktypes.CreateModel
	UserID    uint32   `json:"user_id" gorm:"uniqueIndex:uc;not null;comment:用户id"`
	ChannelID uint32   `json:"channel_id"  gorm:"uniqueIndex:uc;not null;comment:渠道id"`
	AuthID    string   `json:"auth_id" gorm:"uniqueIndex:pa;binary;not null;size:64;comment:授权ID"`
	UnionID   *string  `json:"union_id" gorm:"binary;size:64;comment:平台联合ID"`
	Token     *string  `json:"token" gorm:"size:64;comment:平台Token"`
	ExpireAt  int64    `json:"expire_at" gorm:"comment:过期时间"`
	LoginAt   int64    `json:"login_at" gorm:"comment:最近授权"`
	User      *User    `json:"-" gorm:"constraint:onDelete:cascade"`
	Channel   *Channel `json:"-"` // 不允许直接删除channel
}

type UserChannelRepo interface {
	Create(ctx kratosx.Context, uid, aid uint32) (uint32, error)
	Delete(ctx kratosx.Context, uid, aid uint32) error
}

type UserChannelUseCase struct {
	config *config.Config
	repo   UserChannelRepo
}

func NewUserChannelUseCase(config *config.Config, repo UserChannelRepo) *UserChannelUseCase {
	return &UserChannelUseCase{config: config, repo: repo}
}

// Add 添加用户应用信息
func (u *UserChannelUseCase) Add(ctx kratosx.Context, uid, aid uint32) (uint32, error) {
	id, err := u.repo.Create(ctx, uid, aid)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Delete 删除用户应用信息
func (u *UserChannelUseCase) Delete(ctx kratosx.Context, uid, aid uint32) error {
	if err := u.repo.Delete(ctx, uid, aid); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
