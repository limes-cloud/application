package channel

import "github.com/limes-cloud/kratosx"

type Repo interface {
	All(ctx kratosx.Context) ([]*Channel, error)
	GetByPlatform(ctx kratosx.Context, platform string) (*Channel, error)
	Create(ctx kratosx.Context, c *Channel) (uint32, error)
	Update(ctx kratosx.Context, c *Channel) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}
