package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"

	v1 "github.com/limes-cloud/user-center/api/v1"
)

func (s *Service) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.User, error) {
	return s.user.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) PageUser(ctx context.Context, in *v1.PageUserRequest) (*v1.PageUserReply, error) {
	return s.user.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) AddUser(ctx context.Context, in *v1.AddUserRequest) (*v1.AddUserReply, error) {
	return s.user.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) ImportUser(ctx context.Context, in *v1.ImportUserRequest) (*empty.Empty, error) {
	return s.user.Import(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*empty.Empty, error) {
	return s.user.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*empty.Empty, error) {
	return s.user.Delete(kratosx.MustContext(ctx), in)
}

func (s *Service) DisableUser(ctx context.Context, in *v1.DisableUserRequest) (*empty.Empty, error) {
	return s.user.Disable(kratosx.MustContext(ctx), in)
}

func (s *Service) EnableUser(ctx context.Context, in *v1.EnableUserRequest) (*empty.Empty, error) {
	return s.user.Enable(kratosx.MustContext(ctx), in)
}

func (s *Service) OfflineUser(ctx context.Context, in *v1.OfflineUserRequest) (*empty.Empty, error) {
	return s.user.Offline(kratosx.MustContext(ctx), in)
}
