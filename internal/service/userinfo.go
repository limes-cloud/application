package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	pb "github.com/limes-cloud/usercenter/api/usercenter/userinfo/v1"
	"github.com/limes-cloud/usercenter/internal/biz/userinfo"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/data"
)

type UserinfoService struct {
	pb.UnimplementedUserinfoServer
	uc *userinfo.UseCase
}

func NewUserinfoService(conf *conf.Config) *UserinfoService {
	return &UserinfoService{
		uc: userinfo.NewUseCase(conf, data.NewUserinfoRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewUserinfoService(c)
		pb.RegisterUserinfoHTTPServer(hs, srv)
		pb.RegisterUserinfoServer(gs, srv)
	})
}

// ListUserinfo 获取用户扩展信息列表
func (s *UserinfoService) ListUserinfo(c context.Context, req *pb.ListUserinfoRequest) (*pb.ListUserinfoReply, error) {
	var (
		in  = userinfo.ListUserinfoRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListUserinfo(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListUserinfoReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// UpdateUserinfo 更新用户扩展信息
func (s *UserinfoService) UpdateUserinfo(c context.Context, req *pb.UpdateUserinfoRequest) (*pb.UpdateUserinfoReply, error) {
	var (
		in  = userinfo.Userinfo{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateUserinfo(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateUserinfoReply{}, nil
}

//
// // ListCurrentUserinfo 获取当前用户信息列表
// func (s *UserinfoService) ListCurrentUserinfo(c context.Context, _ *emptypb.Empty) (*pb.ListUserinfoReply, error) {
//	var (
//		ctx = kratosx.MustContext(c)
//	)
//	result, err := s.uc.ListCurrentUserinfo(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	reply := pb.ListUserinfoReply{}
//	if err := valx.Transform(result, &reply.List); err != nil {
//		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
//		return nil, errors.TransformError()
//	}
//
//	return &reply, nil
// }
//
// // UpdateCurrentUserinfo 更新当前用户信息
// func (s *UserinfoService) UpdateCurrentUserinfo(c context.Context, req *pb.UpdateCurrentUserinfoRequest) (*pb.UpdateCurrentUserinfoReply, error) {
//	var (
//		in  []*userinfo.Userinfo
//		ctx = kratosx.MustContext(c)
//	)
//
//	for _, item := range req.List {
//		in = append(in, &userinfo.Userinfo{
//			Keyword: item.Keyword,
//			Value:   item.Value,
//		})
//	}
//
//	if err := s.uc.UpdateCurrentUserinfo(ctx, in); err != nil {
//		return nil, err
//	}
//
//	return &pb.UpdateCurrentUserinfoReply{}, nil
// }

// CreateUserinfo 创建用户扩展信息
func (s *UserinfoService) CreateUserinfo(c context.Context, req *pb.CreateUserinfoRequest) (*pb.CreateUserinfoReply, error) {
	var (
		in  = userinfo.Userinfo{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateUserinfo(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserinfoReply{Id: id}, nil
}

// DeleteUserinfo 删除用户扩展信息
func (s *UserinfoService) DeleteUserinfo(c context.Context, req *pb.DeleteUserinfoRequest) (*pb.DeleteUserinfoReply, error) {
	total, err := s.uc.DeleteUserinfo(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserinfoReply{Total: total}, nil
}
