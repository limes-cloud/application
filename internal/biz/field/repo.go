package field

import "github.com/limes-cloud/kratosx"

type Repo interface {
	Page(ctx kratosx.Context, req *PageFieldRequest) ([]*Field, uint32, error)
	All(ctx kratosx.Context) ([]*Field, error)
	Create(ctx kratosx.Context, c *Field) (uint32, error)
	Update(ctx kratosx.Context, c *Field) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}
