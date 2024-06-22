package channel

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// GetChannel 获取指定的登陆渠道
	GetChannel(ctx kratosx.Context, id uint32) (*Channel, error)

	// ListChannel 获取登陆渠道列表
	ListChannel(ctx kratosx.Context, req *ListChannelRequest) ([]*Channel, uint32, error)

	// CreateChannel 创建登陆渠道
	CreateChannel(ctx kratosx.Context, req *Channel) (uint32, error)

	// UpdateChannel 更新登陆渠道
	UpdateChannel(ctx kratosx.Context, req *Channel) error

	// UpdateChannelStatus 更新登陆渠道状态
	UpdateChannelStatus(ctx kratosx.Context, id uint32, status bool) error

	// DeleteChannel 删除登陆渠道
	DeleteChannel(ctx kratosx.Context, ids []uint32) (uint32, error)

	// GetChannelByKeyword 获取指定的登陆渠道
	GetChannelByKeyword(ctx kratosx.Context, keyword string) (*Channel, error)
}
