package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/limes-cloud/application/api/application/userinfo/v1"
	"github.com/limes-cloud/application/internal/conf"
	"github.com/limes-cloud/application/internal/domain/entity"
	"github.com/limes-cloud/application/internal/domain/service"
	"github.com/limes-cloud/application/internal/infra/dbs"
	"github.com/limes-cloud/application/internal/infra/rpc"
	"github.com/limes-cloud/application/internal/types"
)

type Userinfo struct {
	pb.UnimplementedUserinfoServer
	srv *service.Userinfo
}

func NewUserinfo(conf *conf.Config) *Userinfo {
	return &Userinfo{
		srv: service.NewUserinfo(
			conf,
			dbs.NewUserinfo(),
			rpc.NewPermission(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewUserinfo(c)
		pb.RegisterUserinfoHTTPServer(hs, srv)
		pb.RegisterUserinfoServer(gs, srv)
	})
}

// ListUserinfo 获取用户扩展信息列表
func (s *Userinfo) ListUserinfo(c context.Context, req *pb.ListUserinfoRequest) (*pb.ListUserinfoReply, error) {
	list, total, err := s.srv.ListUserinfo(kratosx.MustContext(c), &types.ListUserinfoRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		UserId:   req.UserId,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListUserinfoReply{Total: total}
	for _, item := range list {
		if item.Field == nil {
			continue
		}
		reply.List = append(reply.List, &pb.ListUserinfoReply_Userinfo{
			Id:        item.Id,
			UserId:    item.UserId,
			Keyword:   item.Keyword,
			Value:     item.Value,
			Name:      item.Field.Name,
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// ListCurrentUserinfo 获取当前用户信息列表
func (s *Userinfo) ListCurrentUserinfo(c context.Context, _ *emptypb.Empty) (*pb.ListUserinfoReply, error) {
	list, err := s.srv.ListCurrentUserinfo(kratosx.MustContext(c))
	if err != nil {
		return nil, err
	}

	reply := pb.ListUserinfoReply{Total: uint32(len(list))}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListUserinfoReply_Userinfo{
			Id:        item.Id,
			UserId:    item.UserId,
			Keyword:   item.Keyword,
			Value:     item.Value,
			Name:      item.Field.Name,
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// UpdateCurrentUserinfo 更新当前用户信息
func (s *Userinfo) UpdateCurrentUserinfo(c context.Context, req *pb.UpdateCurrentUserinfoRequest) (*pb.UpdateCurrentUserinfoReply, error) {
	var list []*entity.Userinfo
	for _, item := range req.List {
		list = append(list, &entity.Userinfo{
			Keyword: item.Keyword,
			Value:   item.Value,
		})
	}

	if err := s.srv.UpdateCurrentUserinfo(kratosx.MustContext(c), list); err != nil {
		return nil, err
	}

	return &pb.UpdateCurrentUserinfoReply{}, nil
}

// CreateUserinfo 创建用户扩展信息
func (s *Userinfo) CreateUserinfo(c context.Context, req *pb.CreateUserinfoRequest) (*pb.CreateUserinfoReply, error) {
	id, err := s.srv.CreateUserinfo(kratosx.MustContext(c), &entity.Userinfo{
		UserId:  req.UserId,
		Keyword: req.Keyword,
		Value:   req.Value,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserinfoReply{Id: id}, nil
}

// UpdateUserinfo 更新用户扩展信息
func (s *Userinfo) UpdateUserinfo(c context.Context, req *pb.UpdateUserinfoRequest) (*pb.UpdateUserinfoReply, error) {
	if err := s.srv.UpdateUserinfo(kratosx.MustContext(c), &entity.Userinfo{
		BaseModel: ktypes.BaseModel{Id: req.Id},
		Keyword:   req.Keyword,
		Value:     req.Value,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateUserinfoReply{}, nil
}

// DeleteUserinfo 删除用户扩展信息
func (s *Userinfo) DeleteUserinfo(c context.Context, req *pb.DeleteUserinfoRequest) (*pb.DeleteUserinfoReply, error) {
	if err := s.srv.DeleteUserinfo(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteUserinfoReply{}, nil
}
