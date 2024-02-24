package agreement

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/user-center/api/errors"
	"github.com/limes-cloud/user-center/internal/config"
)

type UseCase struct {
	config *config.Config
	repo   Repo
}

// NewUseCase 创建UseCase实体
func NewUseCase(config *config.Config, repo Repo) *UseCase {
	return &UseCase{config: config, repo: repo}
}

// GetContent 获取详细协议内容
func (u *UseCase) GetContent(ctx kratosx.Context, id uint32) (*Content, error) {
	content, err := u.repo.GetContent(ctx, id)
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}
	return content, nil
}

// PageContent 获取分页协议内容
func (u *UseCase) PageContent(ctx kratosx.Context, req *PageContentRequest) ([]*Content, uint32, error) {
	content, total, err := u.repo.PageContent(ctx, req)
	if err != nil {
		return nil, 0, errors.NotRecord()
	}
	return content, total, nil
}

// AddContent 添加协议内容
func (u *UseCase) AddContent(ctx kratosx.Context, content *Content) (uint32, error) {
	id, err := u.repo.AddContent(ctx, content)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	return id, nil
}

// UpdateContent 删除指定协议内容
func (u *UseCase) UpdateContent(ctx kratosx.Context, content *Content) error {
	if err := u.repo.UpdateContent(ctx, content); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// DeleteContent 删除指定协议内容
func (u *UseCase) DeleteContent(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteContent(ctx, id); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// GetSceneByKeyword 获取场景信息
func (u *UseCase) GetSceneByKeyword(ctx kratosx.Context, keyword string) (*Scene, error) {
	content, err := u.repo.GetSceneByKeyword(ctx, keyword)
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}
	return content, nil
}

// PageScene 获取分页场景
func (u *UseCase) PageScene(ctx kratosx.Context, req *PageSceneRequest) ([]*Scene, uint32, error) {
	scene, total, err := u.repo.PageScene(ctx, req)
	if err != nil {
		return nil, 0, errors.NotRecord()
	}
	return scene, total, nil
}

// AddScene 添加场景信息
func (u *UseCase) AddScene(ctx kratosx.Context, scene *Scene) (uint32, error) {
	id, err := u.repo.AddScene(ctx, scene)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	return id, nil
}

// UpdateScene 更新场景信息
func (u *UseCase) UpdateScene(ctx kratosx.Context, scene *Scene) error {
	if err := u.repo.UpdateScene(ctx, scene); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// DeleteScene 删除场景信息
func (u *UseCase) DeleteScene(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteScene(ctx, id); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}
