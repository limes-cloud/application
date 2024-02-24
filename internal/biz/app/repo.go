package app

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	GetByID(ctx kratosx.Context, id uint32) (*App, error)
	GetByKeyword(ctx kratosx.Context, keyword string) (*App, error)
	Page(ctx kratosx.Context, req *PageAppRequest) ([]*App, uint32, error)
	Create(ctx kratosx.Context, c *App) (uint32, error)
	Update(ctx kratosx.Context, c *App) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}
