package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/biz/types"
)

type App struct {
	ktypes.BaseModel
	Keyword       string        `json:"keyword" gorm:"unique;not null;size:32;comment:应用标识"`
	Logo          string        `json:"logo" gorm:"not null;size:128;comment:应用logo"`
	Name          string        `json:"name" gorm:"not null;size:32;comment:应用名称"`
	Status        *bool         `json:"status" gorm:"not null;comment:应用状态"`
	Version       string        `json:"version" gorm:"size:32;comment:应用版本"`
	Copyright     string        `json:"copyright" gorm:"size:128;comment:应用版权"`
	AllowRegistry *bool         `json:"allow_registry" gorm:"not null;comment:是否允许注册"`
	Description   string        `json:"description" gorm:"not null;size:128;comment:应用描述"`
	Channels      []*Channel    `json:"channels" gorm:"many2many:app_channel;constraint:onDelete:cascade"`
	AppChannels   []*AppChannel `json:"app_channels"`
}

type AppChannel struct {
	AppID     uint32 `json:"app_id"`
	ChannelID uint32 `json:"channel_id"`
}

type AppRepo interface {
	GetByKeyword(ctx kratosx.Context, keyword string) (*App, error)
	Page(ctx kratosx.Context, req *types.PageAppRequest) ([]*App, uint32, error)
	Create(ctx kratosx.Context, c *App) (uint32, error)
	Update(ctx kratosx.Context, c *App) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}

type AppUseCase struct {
	config *config.Config
	repo   AppRepo
}

func NewAppUseCase(config *config.Config, repo AppRepo) *AppUseCase {
	return &AppUseCase{config: config, repo: repo}
}

// GetByKeyword 获取全部应用
func (u *AppUseCase) GetByKeyword(ctx kratosx.Context, keyword string) (*App, error) {
	app, err := u.repo.GetByKeyword(ctx, keyword)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return app, nil
}

// Page 获取全部登录应用
func (u *AppUseCase) Page(ctx kratosx.Context, req *types.PageAppRequest) ([]*App, uint32, error) {
	app, total, err := u.repo.Page(ctx, req)
	if err != nil {
		return nil, 0, v1.NotRecordError()
	}
	return app, total, nil
}

// Add 添加登录应用信息
func (u *AppUseCase) Add(ctx kratosx.Context, app *App) (uint32, error) {
	id, err := u.repo.Create(ctx, app)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新登录应用信息
func (u *AppUseCase) Update(ctx kratosx.Context, app *App) error {
	if err := u.repo.Update(ctx, app); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除登录应用信息
func (u *AppUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
