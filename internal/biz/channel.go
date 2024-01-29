package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
)

type Channel struct {
	ktypes.CreateModel
	Logo     string `json:"logo" gorm:"not null;size:128;comment:渠道logo"`
	Platform string `json:"platform" gorm:"unique;not null;binary;type:char(32);comment:渠道标识"`
	Name     string `json:"name" gorm:"not null;size:32;comment:渠道名称"`
	Ak       string `json:"ak" gorm:"size:32;comment:渠道ak"`
	Sk       string `json:"sk" gorm:"size:32;comment:渠道sk"`
	Extra    string `json:"extra" gorm:"size:256;comment:渠道状态"`
	Status   *bool  `json:"status" gorm:"not null;comment:渠道状态"`
}

type ChannelRepo interface {
	All(ctx kratosx.Context) ([]*Channel, error)
	GetByPlatform(ctx kratosx.Context, platform string) (*Channel, error)
	Create(ctx kratosx.Context, c *Channel) (uint32, error)
	Update(ctx kratosx.Context, c *Channel) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}

type ChannelUseCase struct {
	config *config.Config
	repo   ChannelRepo
}

func NewChannelUseCase(config *config.Config, repo ChannelRepo) *ChannelUseCase {
	return &ChannelUseCase{config: config, repo: repo}
}

// All 获取全部登录通道
func (u *ChannelUseCase) All(ctx kratosx.Context) ([]*Channel, error) {
	channel, err := u.repo.All(ctx)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return channel, nil
}

// GetByPlatform 获取指定登录通道
func (u *ChannelUseCase) GetByPlatform(ctx kratosx.Context, platform string) (*Channel, error) {
	channel, err := u.repo.GetByPlatform(ctx, platform)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return channel, nil
}

// Add 添加登录通道信息
func (u *ChannelUseCase) Add(ctx kratosx.Context, channel *Channel) (uint32, error) {
	id, err := u.repo.Create(ctx, channel)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新登录通道信息
func (u *ChannelUseCase) Update(ctx kratosx.Context, channel *Channel) error {
	if err := u.repo.Update(ctx, channel); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除登录通道信息
func (u *ChannelUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
