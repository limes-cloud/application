package channel

import (
	"sort"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/user-center/api/errors"
	"github.com/limes-cloud/user-center/internal/config"
	"github.com/limes-cloud/user-center/internal/consts"
	"github.com/limes-cloud/user-center/internal/pkg/authorizer"
)

type UseCase struct {
	config *config.Config
	repo   Repo
}

func NewUseCase(config *config.Config, repo Repo) *UseCase {
	return &UseCase{config: config, repo: repo}
}

// All 获取全部登录通道
func (u *UseCase) All(ctx kratosx.Context) ([]*Channel, error) {
	channel, err := u.repo.All(ctx)
	if err != nil {
		return nil, errors.NotRecord()
	}
	return channel, nil
}

// GetTypes 获取可以开通的登录渠道
func (u *UseCase) GetTypes() ([]*Typer, error) {
	list := []*Typer{
		{
			Platform: consts.PasswordChannel,
			Name:     "密码",
		},
		{
			Platform: consts.CaptchaChannel,
			Name:     "验证码",
		},
	}

	var keys []string
	ath := authorizer.New()
	set := ath.GetAuthorizers()
	for key := range set {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		list = append(list, &Typer{
			Platform: key,
			Name:     set[key].Name(),
		})
	}

	return list, nil
}

// GetByPlatform 获取指定登录通道
func (u *UseCase) GetByPlatform(ctx kratosx.Context, platform string) (*Channel, error) {
	channel, err := u.repo.GetByPlatform(ctx, platform)
	if err != nil {
		return nil, errors.NotRecord()
	}
	return channel, nil
}

// Add 添加登录通道信息
func (u *UseCase) Add(ctx kratosx.Context, channel *Channel) (uint32, error) {
	id, err := u.repo.Create(ctx, channel)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	return id, nil
}

// Update 更新登录通道信息
func (u *UseCase) Update(ctx kratosx.Context, channel *Channel) error {
	if err := u.repo.Update(ctx, channel); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// Delete 删除登录通道信息
func (u *UseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}
