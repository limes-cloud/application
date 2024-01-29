package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	v1 "github.com/limes-cloud/user-center/api/v1"
)

func (s *Service) AddUserApp(ctx context.Context, in *v1.AddUserAppRequest) (*empty.Empty, error) {
	_, err := s.userApp.Add(kratosx.MustContext(ctx), in.UserId, in.AppId)
	return nil, err
}

func (s *Service) DeleteUserApp(ctx context.Context, in *v1.DeleteUserAppRequest) (*empty.Empty, error) {
	return nil, s.userApp.Delete(kratosx.MustContext(ctx), in.UserId, in.AppId)
}
