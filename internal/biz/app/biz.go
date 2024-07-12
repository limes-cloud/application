package app

import (
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/pkg/permission"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// GetApp 获取指定的应用信息
func (u *UseCase) GetApp(ctx kratosx.Context, req *GetAppRequest) (*App, error) {
	var (
		res *App
		err error
	)

	if req.Id != nil {
		if !permission.HasApp(ctx, *req.Id) {
			return nil, errors.NotPermissionError()
		}
		res, err = u.repo.GetApp(ctx, *req.Id)
	} else if req.Keyword != nil {
		res, err = u.repo.GetAppByKeyword(ctx, *req.Keyword)
		if !permission.HasApp(ctx, res.Id) {
			return nil, errors.NotPermissionError()
		}
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// ListApp 获取应用信息列表
func (u *UseCase) ListApp(ctx kratosx.Context, req *ListAppRequest) ([]*App, uint32, error) {
	all, scopes, err := permission.GetApp(ctx)
	if err != nil {
		return nil, 0, err
	}
	if !all {
		req.Ids = scopes
	}
	list, total, err := u.repo.ListApp(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateApp 创建应用信息
func (u *UseCase) CreateApp(ctx kratosx.Context, req *App) (uint32, error) {
	req.Status = proto.Bool(false)
	req.DisableDesc = proto.String("应用未发布")
	id, err := u.repo.CreateApp(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateApp 更新应用信息
func (u *UseCase) UpdateApp(ctx kratosx.Context, req *App) error {
	if !permission.HasApp(ctx, req.Id) {
		return errors.NotPermissionError()
	}
	if err := u.repo.UpdateApp(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateAppStatus 更新应用信息状态
func (u *UseCase) UpdateAppStatus(ctx kratosx.Context, req *UpdateAppStatusRequest) error {
	if !permission.HasApp(ctx, req.Id) {
		return errors.NotPermissionError()
	}

	if err := u.repo.UpdateAppStatus(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteApp 删除应用信息
func (u *UseCase) DeleteApp(ctx kratosx.Context, id uint32) error {
	if !permission.HasApp(ctx, id) {
		return errors.NotPermissionError()
	}
	if err := u.repo.DeleteApp(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
