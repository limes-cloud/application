package agreement

import (
	"github.com/limes-cloud/kratosx"

	biz "github.com/limes-cloud/user-center/internal/biz/channel"
)

type repo struct {
}

func NewRepo() biz.Repo {
	return &repo{}
}

func (u *repo) All(ctx kratosx.Context) ([]*biz.Channel, error) {
	var list []*biz.Channel
	return list, ctx.DB().Find(&list).Error
}

func (u *repo) GetByPlatform(ctx kratosx.Context, platform string) (*biz.Channel, error) {
	var channel biz.Channel
	return &channel, ctx.DB().First(&channel, "platform=?", platform).Error
}

func (u *repo) Create(ctx kratosx.Context, channel *biz.Channel) (uint32, error) {
	return channel.ID, ctx.DB().Create(channel).Error
}

func (u *repo) Update(ctx kratosx.Context, channel *biz.Channel) error {
	return ctx.DB().Updates(channel).Error
}

func (u *repo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Channel{}, "id=?", id).Error
}
