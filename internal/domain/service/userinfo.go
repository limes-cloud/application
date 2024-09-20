package service

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/application/api/application/errors"
	"github.com/limes-cloud/application/internal/conf"
	"github.com/limes-cloud/application/internal/domain/entity"
	"github.com/limes-cloud/application/internal/domain/repository"
	"github.com/limes-cloud/application/internal/pkg/md"
	"github.com/limes-cloud/application/internal/types"
)

type Userinfo struct {
	conf       *conf.Config
	repo       repository.Userinfo
	permission repository.Permission
}

func NewUserinfo(
	conf *conf.Config,
	repo repository.Userinfo,
	permission repository.Permission,
) *Userinfo {
	return &Userinfo{
		conf:       conf,
		repo:       repo,
		permission: permission,
	}
}

// ListUserinfo 获取用户扩展信息列表
func (u *Userinfo) ListUserinfo(ctx kratosx.Context, req *types.ListUserinfoRequest) ([]*entity.Userinfo, uint32, error) {
	list, total, err := u.repo.ListUserinfo(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list userinfo error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// UpdateUserinfo 更新用户扩展信息
func (u *Userinfo) UpdateUserinfo(ctx kratosx.Context, userinfo *entity.Userinfo) error {
	if err := u.repo.UpdateUserinfo(ctx, userinfo); err != nil {
		ctx.Logger().Warnw("msg", "update userinfo error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// ListCurrentUserinfo 获取当前用户信息列表
func (u *Userinfo) ListCurrentUserinfo(ctx kratosx.Context) ([]*entity.Userinfo, error) {
	info, err := md.Get(ctx)
	if err != nil {
		ctx.Logger().Warnw("msg", "get auth info error", "err", err.Error())
		return nil, errors.SystemError()
	}

	list, _, err := u.repo.ListUserinfo(ctx, &types.ListUserinfoRequest{UserId: info.UserId, AppKeyword: &info.AppKeyword})
	if err != nil {
		ctx.Logger().Warnw("msg", "list userinfo error", "err", err.Error())
		return nil, errors.ListError(err.Error())
	}
	return list, nil
}

// UpdateCurrentUserinfo 更新当前用户信息
func (u *Userinfo) UpdateCurrentUserinfo(ctx kratosx.Context, list []*entity.Userinfo) error {
	info, err := md.Get(ctx)
	if err != nil {
		ctx.Logger().Warnw("msg", "get auth info error", "err", err.Error())
		return errors.SystemError()
	}

	var keys []string
	for _, item := range list {
		item.UserId = info.UserId
		keys = append(keys, item.Keyword)
	}

	if err := u.repo.CheckKeywords(ctx, info.AppId, keys); err != nil {
		return errors.ParamsError(err.Error())
	}

	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		for _, item := range list {
			if err = u.repo.UpdateUserinfo(ctx, item); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		ctx.Logger().Warnw("msg", "update cur userinfo error", "err", err.Error())
		return errors.UpdateError()
	}

	return nil
}

// CreateUserinfo 创建用户扩展信息
func (u *Userinfo) CreateUserinfo(ctx kratosx.Context, req *entity.Userinfo) (uint32, error) {
	id, err := u.repo.CreateUserinfo(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create userinfo error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// DeleteUserinfo 删除用户扩展信息
func (u *Userinfo) DeleteUserinfo(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteUserinfo(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete userinfo error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}
