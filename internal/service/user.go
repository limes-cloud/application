package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	pb "github.com/limes-cloud/usercenter/api/usercenter/user/v1"
	"github.com/limes-cloud/usercenter/internal/biz/user"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/data"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *user.UseCase
}

func NewUserService(conf *conf.Config) *UserService {
	return &UserService{
		uc: user.NewUseCase(conf, data.NewUserRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewUserService(c)
		pb.RegisterUserHTTPServer(hs, srv)
		pb.RegisterUserServer(gs, srv)
	})
}

// GetCurrentUser 获取当前用户信息
func (s *UserService) GetCurrentUser(c context.Context, _ *pb.GetCurrentUserRequest) (*pb.GetCurrentUserReply, error) {
	var (
		ctx = kratosx.MustContext(c)
	)

	result, err := s.uc.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	reply := pb.GetCurrentUserReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// UpdateCurrentUser 更新当前用户信息
func (s *UserService) UpdateCurrentUser(c context.Context, req *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserReply, error) {
	return &pb.UpdateCurrentUserReply{}, s.uc.UpdateCurrentUser(kratosx.MustContext(c), &user.UpdateCurrentUserRequest{
		NickName: req.NickName,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
	})
}

// GetUser 获取指定的用户信息
func (s *UserService) GetUser(c context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	var (
		in  = user.GetUserRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.GetUser(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.GetUserReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListUser 获取用户信息列表
func (s *UserService) ListUser(c context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	var (
		in  = user.ListUserRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListUser(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListUserReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateUser 创建用户信息
func (s *UserService) CreateUser(c context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	var (
		in  = user.User{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateUser(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserReply{Id: id}, nil
}

// ImportUser 导入用户信息
func (s *UserService) ImportUser(c context.Context, req *pb.ImportUserRequest) (*pb.ImportUserReply, error) {
	var (
		in  []*user.User
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req.List, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	total, err := s.uc.ImportUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.ImportUserReply{Total: total}, nil
}

// ExportUser 导出用户信息
func (s *UserService) ExportUser(c context.Context, req *pb.ExportUserRequest) (*pb.ExportUserReply, error) {
	var (
		in  = user.ExportUserRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	src, err := s.uc.ExportUser(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.ExportUserReply{Src: src}, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(c context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	var (
		in  = user.User{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateUser(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateUserReply{}, nil
}

// UpdateUserStatus 更新用户信息状态
func (s *UserService) UpdateUserStatus(c context.Context, req *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusReply, error) {
	return &pb.UpdateUserStatusReply{}, s.uc.UpdateUserStatus(kratosx.MustContext(c), &user.UpdateUserStatusRequest{
		Id:          req.Id,
		Status:      req.Status,
		DisableDesc: req.DisableDesc,
	})
}

// DeleteUser 删除用户信息
func (s *UserService) DeleteUser(c context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	total, err := s.uc.DeleteUser(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserReply{Total: total}, nil
}

// GetTrashUser 获取回收站指定的用户信息
func (s *UserService) GetTrashUser(c context.Context, req *pb.GetTrashUserRequest) (*pb.GetTrashUserReply, error) {
	var ctx = kratosx.MustContext(c)

	result, err := s.uc.GetTrashUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	reply := pb.GetTrashUserReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListTrashUser 获取回收站用户信息列表
func (s *UserService) ListTrashUser(c context.Context, req *pb.ListTrashUserRequest) (*pb.ListTrashUserReply, error) {
	var (
		in  = user.ListTrashUserRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListTrashUser(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListTrashUserReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// DeleteTrashUser 删除回收站用户信息
func (s *UserService) DeleteTrashUser(ctx context.Context, req *pb.DeleteTrashUserRequest) (*pb.DeleteTrashUserReply, error) {
	total, err := s.uc.DeleteTrashUser(kratosx.MustContext(ctx), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTrashUserReply{Total: total}, nil
}

// RevertTrashUser 还原回收站用户信息
func (s *UserService) RevertTrashUser(ctx context.Context, req *pb.RevertTrashUserRequest) (*pb.RevertTrashUserReply, error) {
	return &pb.RevertTrashUserReply{}, s.uc.RevertTrashUser(kratosx.MustContext(ctx), req.Id)
}
