package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/user-center/internal/biz"
)

type channelRepo struct {
}

func NewChannelRepo() biz.ChannelRepo {
	return &channelRepo{}
}

func (u *channelRepo) All(ctx kratosx.Context) ([]*biz.Channel, error) {
	var list []*biz.Channel
	return list, ctx.DB().Find(&list).Error
}

func (u *channelRepo) GetByPlatform(ctx kratosx.Context, platform string) (*biz.Channel, error) {
	var channel biz.Channel
	return &channel, ctx.DB().First(&channel, "platform=?", platform).Error
}

func (u *channelRepo) Create(ctx kratosx.Context, channel *biz.Channel) (uint32, error) {
	return channel.ID, ctx.DB().Create(channel).Error
}

func (u *channelRepo) Update(ctx kratosx.Context, channel *biz.Channel) error {
	return ctx.DB().Updates(channel).Error
}

func (u *channelRepo) Delete(ctx kratosx.Context, id uint32) error {
	//if err := ctx.DB().First(biz.UserChannel{}, "channel_id=?", id); err == nil {
	//	return errors.New("渠道正在使用中，无法删除")
	//}
	return ctx.DB().Delete(biz.Channel{}, "id=?", id).Error
}
