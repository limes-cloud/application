package app

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/user-center/api/errors"
	"github.com/limes-cloud/user-center/internal/config"
)

type UseCase struct {
	config *config.Config
	repo   Repo
}

func NewUseCase(config *config.Config, repo Repo) *UseCase {
	return &UseCase{config: config, repo: repo}
}

// GetByID 获取指定应用
func (u *UseCase) GetByID(ctx kratosx.Context, id uint32) (*App, error) {
	app, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.NotRecord()
	}
	return app, nil
}

// GetByKeyword 获取指定应用
func (u *UseCase) GetByKeyword(ctx kratosx.Context, keyword string) (*App, error) {
	app, err := u.repo.GetByKeyword(ctx, keyword)
	if err != nil {
		return nil, errors.NotRecord()
	}
	return app, nil
}

// Page 获取全部登录应用
func (u *UseCase) Page(ctx kratosx.Context, req *PageAppRequest) ([]*App, uint32, error) {
	app, total, err := u.repo.Page(ctx, req)
	if err != nil {
		return nil, 0, errors.NotRecord()
	}
	return app, total, nil
}

// Add 添加登录应用信息
func (u *UseCase) Add(ctx kratosx.Context, app *App) (uint32, error) {
	id, err := u.repo.Create(ctx, app)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	return id, nil
}

// Update 更新登录应用信息
func (u *UseCase) Update(ctx kratosx.Context, app *App) error {
	if err := u.repo.Update(ctx, app); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// Delete 删除登录应用信息
func (u *UseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}
