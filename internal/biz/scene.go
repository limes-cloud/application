package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/biz/types"
)

type Scene struct {
	ktypes.BaseModel
	Keyword         string            `json:"keyword" gorm:"unique;binary;not null;size:32;comment:场景标识"`
	Name            string            `json:"name" gorm:"not null;size:32;comment:场景名称"`
	Description     string            `json:"description" gorm:"not null;size:128;comment:场景描述"`
	Agreements      []*Agreement      `json:"agreements" gorm:"many2many:agreement_scene;constraint:onDelete:cascade"`
	AgreementScenes []*AgreementScene `json:"agreement_scenes"`
}

type AgreementScene struct {
	SceneID     uint32 `json:"scene_id"`
	AgreementID uint32 `json:"agreement_id"`
}

type SceneRepo interface {
	GetByKeyword(ctx kratosx.Context, keyword string) (*Scene, error)
	Page(ctx kratosx.Context, req *types.PageSceneRequest) ([]*Scene, uint32, error)
	Create(ctx kratosx.Context, c *Scene) (uint32, error)
	Update(ctx kratosx.Context, c *Scene) error
	Delete(ctx kratosx.Context, id uint32) error
}

type SceneUseCase struct {
	config *config.Config
	repo   SceneRepo
}

func NewSceneUseCase(config *config.Config, repo SceneRepo) *SceneUseCase {
	return &SceneUseCase{config: config, repo: repo}
}

// GetByKeyword 获取场景信息
func (u *SceneUseCase) GetByKeyword(ctx kratosx.Context, keyword string) (*Scene, error) {
	agreement, err := u.repo.GetByKeyword(ctx, keyword)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return agreement, nil
}

// Page 获取全部场景
func (u *SceneUseCase) Page(ctx kratosx.Context, req *types.PageSceneRequest) ([]*Scene, uint32, error) {
	agreement, total, err := u.repo.Page(ctx, req)
	if err != nil {
		return nil, 0, v1.NotRecordError()
	}
	return agreement, total, nil
}

// Add 添加场景信息
func (u *SceneUseCase) Add(ctx kratosx.Context, agreement *Scene) (uint32, error) {
	id, err := u.repo.Create(ctx, agreement)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新场景信息
func (u *SceneUseCase) Update(ctx kratosx.Context, agreement *Scene) error {
	if err := u.repo.Update(ctx, agreement); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除场景信息
func (u *SceneUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
