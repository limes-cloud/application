package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	resourcev1 "github.com/limes-cloud/manager/api/manager/resource/v1"
	file "github.com/limes-cloud/resource/api/resource/file/v1"

	"github.com/limes-cloud/usercenter/api/usercenter/errors"
)

const (
	Resource = "Resource"
	Manager  = "Manager"
)

func NewResourceFile(ctx context.Context) (file.FileClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.ResourceServerError()
	}
	return file.NewFileClient(conn), nil
}

func NewManagerResource(ctx kratosx.Context) (resourcev1.ResourceClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Manager)
	if err != nil {
		return nil, errors.ManagerServerError()
	}
	return resourcev1.NewResourceClient(conn), nil
}
