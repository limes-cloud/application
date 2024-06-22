package app

import (
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

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

// GetApp 获取指定的应用信息
func (u *UseCase) GetApp(ctx kratosx.Context, req *GetAppRequest) (*App, error) {
	var (
		res *App
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetApp(ctx, *req.Id)
	} else if req.Keyword != nil {
		res, err = u.repo.GetAppByKeyword(ctx, *req.Keyword)
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
	if err := u.repo.UpdateApp(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateAppStatus 更新应用信息状态
func (u *UseCase) UpdateAppStatus(ctx kratosx.Context, req *UpdateAppStatusRequest) error {
	if err := u.repo.UpdateAppStatus(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteApp 删除应用信息
func (u *UseCase) DeleteApp(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteApp(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}
