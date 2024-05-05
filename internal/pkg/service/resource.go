package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	resource "github.com/limes-cloud/resource/api/file/v1"

	"github.com/limes-cloud/user-center/api/errors"
)

const (
	Resource = "Resource"
)

func NewResource(ctx context.Context) (resource.ServiceClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.ResourceServer()
	}
	return resource.NewServiceClient(conn), nil
}
