package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	resource "github.com/limes-cloud/resource/api/v1"
)

func NewResource(ctx context.Context, host string) (resource.ServiceClient, error) {
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint(host))
	if err != nil {
		return nil, err
	}
	return resource.NewServiceClient(conn), nil
}
