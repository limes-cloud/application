package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/biz/types"
)

type Agreement struct {
	ktypes.BaseModel
	Name        string `json:"name" gorm:"not null;size:32;comment:协议名称"`
	Status      *bool  `json:"status" gorm:"not null;comment:协议状态"`
	Content     string `json:"content" gorm:"type:blob;not null;comment:协议内容"`
	Description string `json:"description" gorm:"not null;size:128;comment:协议描述"`
}

type AgreementRepo interface {
	Get(ctx kratosx.Context, id uint32) (*Agreement, error)
	Page(ctx kratosx.Context, req *types.PageAgreementRequest) ([]*Agreement, uint32, error)
	Create(ctx kratosx.Context, c *Agreement) (uint32, error)
	Update(ctx kratosx.Context, c *Agreement) error
	Delete(ctx kratosx.Context, id uint32) error
}

type AgreementUseCase struct {
	config *config.Config
	repo   AgreementRepo
}

func NewAgreementUseCase(config *config.Config, repo AgreementRepo) *AgreementUseCase {
	return &AgreementUseCase{config: config, repo: repo}
}

// Get 获取歇息信息
func (u *AgreementUseCase) Get(ctx kratosx.Context, id uint32) (*Agreement, error) {
	agreement, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return agreement, nil
}

// Page 获取全部登录协议
func (u *AgreementUseCase) Page(ctx kratosx.Context, req *types.PageAgreementRequest) ([]*Agreement, uint32, error) {
	agreement, total, err := u.repo.Page(ctx, req)
	if err != nil {
		return nil, 0, v1.NotRecordError()
	}
	return agreement, total, nil
}

// Add 添加登录协议信息
func (u *AgreementUseCase) Add(ctx kratosx.Context, agreement *Agreement) (uint32, error) {
	id, err := u.repo.Create(ctx, agreement)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新登录协议信息
func (u *AgreementUseCase) Update(ctx kratosx.Context, agreement *Agreement) error {
	if err := u.repo.Update(ctx, agreement); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除登录协议信息
func (u *AgreementUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
