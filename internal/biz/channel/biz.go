package channel

import (
	"sort"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/consts"
	"github.com/limes-cloud/usercenter/internal/pkg/authorizer"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// GetTypes 获取可以开通的登录渠道
func (u *UseCase) GetTypes() []*Typer {
	list := []*Typer{
		{
			Keyword: consts.PasswordChannel,
			Name:    "密码",
		},
		{
			Keyword: consts.EmailChannel,
			Name:    "邮箱",
		},
	}

	var keys []string
	ath := authorizer.New(nil)
	set := ath.GetAuthorizers()
	for key := range set {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		list = append(list, &Typer{
			Keyword: key,
			Name:    set[key].Name(),
		})
	}

	return list
}

// ListChannel 获取登陆渠道列表
func (u *UseCase) ListChannel(ctx kratosx.Context, req *ListChannelRequest) ([]*Channel, uint32, error) {
	list, total, err := u.repo.ListChannel(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateChannel 创建登陆渠道
func (u *UseCase) CreateChannel(ctx kratosx.Context, req *Channel) (uint32, error) {
	req.Status = proto.Bool(false)
	id, err := u.repo.CreateChannel(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateChannel 更新登陆渠道
func (u *UseCase) UpdateChannel(ctx kratosx.Context, req *Channel) error {
	if err := u.repo.UpdateChannel(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateChannelStatus 更新登陆渠道状态
func (u *UseCase) UpdateChannelStatus(ctx kratosx.Context, id uint32, status bool) error {
	if err := u.repo.UpdateChannelStatus(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteChannel 删除登陆渠道
func (u *UseCase) DeleteChannel(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteChannel(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}
