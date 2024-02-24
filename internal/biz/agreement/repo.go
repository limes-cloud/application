package agreement

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	GetContent(ctx kratosx.Context, id uint32) (*Content, error)
	PageContent(ctx kratosx.Context, req *PageContentRequest) ([]*Content, uint32, error)
	AddContent(ctx kratosx.Context, c *Content) (uint32, error)
	UpdateContent(ctx kratosx.Context, c *Content) error
	DeleteContent(ctx kratosx.Context, id uint32) error

	GetSceneByKeyword(ctx kratosx.Context, keyword string) (*Scene, error)
	PageScene(ctx kratosx.Context, req *PageSceneRequest) ([]*Scene, uint32, error)
	AddScene(ctx kratosx.Context, c *Scene) (uint32, error)
	UpdateScene(ctx kratosx.Context, c *Scene) error
	DeleteScene(ctx kratosx.Context, id uint32) error
}
