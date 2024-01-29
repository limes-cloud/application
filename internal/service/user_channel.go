package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	v1 "github.com/limes-cloud/user-center/api/v1"
)

func (s *Service) AddUserChannel(ctx context.Context, in *v1.AddUserChannelRequest) (*empty.Empty, error) {
	_, err := s.userChannel.Add(kratosx.MustContext(ctx), in.UserId, in.ChannelId)
	return nil, err
}

func (s *Service) DeleteUserChannel(ctx context.Context, in *v1.DeleteUserChannelRequest) (*empty.Empty, error) {
	return nil, s.userChannel.Delete(kratosx.MustContext(ctx), in.UserId, in.ChannelId)
}
