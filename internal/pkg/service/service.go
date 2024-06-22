package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	file "github.com/limes-cloud/resource/api/resource/file/v1"

	"github.com/limes-cloud/usercenter/api/usercenter/errors"
)

const (
	Resource = "Resource"
)

func NewResourceFile(ctx context.Context) (file.FileClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.ResourceServerError()
	}
	return file.NewFileClient(conn), nil
}
