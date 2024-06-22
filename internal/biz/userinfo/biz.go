package userinfo

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	"github.com/limes-cloud/usercenter/internal/conf"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// ListUserinfo 获取用户扩展信息列表
func (u *UseCase) ListUserinfo(ctx kratosx.Context, req *ListUserinfoRequest) ([]*Userinfo, uint32, error) {
	list, total, err := u.repo.ListUserinfo(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// UpdateUserinfo 更新用户扩展信息
func (u *UseCase) UpdateUserinfo(ctx kratosx.Context, req *Userinfo) error {
	if err := u.repo.UpdateUserinfo(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// // ListCurrentUserinfo 获取当前用户信息列表
// func (u *UseCase) ListCurrentUserinfo(ctx kratosx.Context) ([]*Userinfo, error) {
//	info, err := md.Get(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	list, err := u.repo.ListUserinfo(ctx, &ListUserinfoRequest{UserId: info.UserId, AppId: &info.AppId})
//	if err != nil {
//		return nil, errors.ListError(err.Error())
//	}
//	return list, nil
// }
//
// // UpdateCurrentUserinfo 更新当前用户信息
// func (u *UseCase) UpdateCurrentUserinfo(ctx kratosx.Context, list []*Userinfo) error {
//	info, err := md.Get(ctx)
//	if err != nil {
//		return err
//	}
//
//	var keys []string
//	for _, item := range list {
//		item.UserId = info.UserId
//		keys = append(keys, item.Keyword)
//	}
//	if err := u.repo.CheckKeywords(ctx, info.AppId, keys); err != nil {
//		return errors.ParamsError(err.Error())
//	}
//
//	if _, err = u.repo.UpdateUserinfo(ctx, list); err != nil {
//		return errors.UpdateError(err.Error())
//	}
//	return nil
// }

// CreateUserinfo 创建用户扩展信息
func (u *UseCase) CreateUserinfo(ctx kratosx.Context, req *Userinfo) (uint32, error) {
	id, err := u.repo.CreateUserinfo(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// DeleteUserinfo 删除用户扩展信息
func (u *UseCase) DeleteUserinfo(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteUserinfo(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}
