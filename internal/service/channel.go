package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "github.com/limes-cloud/usercenter/api/usercenter/channel/v1"
	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	"github.com/limes-cloud/usercenter/internal/biz/channel"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/data"
)

type ChannelService struct {
	pb.UnimplementedChannelServer
	uc *channel.UseCase
}

func NewChannelService(conf *conf.Config) *ChannelService {
	return &ChannelService{
		uc: channel.NewUseCase(conf, data.NewChannelRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewChannelService(c)
		pb.RegisterChannelHTTPServer(hs, srv)
		pb.RegisterChannelServer(gs, srv)
	})
}

// ListChannelType 获取登陆渠道可用列表
func (s *ChannelService) ListChannelType(_ context.Context, _ *pb.ListChannelTypeRequest) (*pb.ListChannelTypeReply, error) {
	types := s.uc.GetTypes()
	reply := pb.ListChannelTypeReply{}
	if err := valx.Transform(types, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListChannel 获取登陆渠道列表
func (s *ChannelService) ListChannel(c context.Context, req *pb.ListChannelRequest) (*pb.ListChannelReply, error) {
	var (
		in  = channel.ListChannelRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListChannel(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListChannelReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateChannel 创建登陆渠道
func (s *ChannelService) CreateChannel(c context.Context, req *pb.CreateChannelRequest) (*pb.CreateChannelReply, error) {
	var (
		in  = channel.Channel{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateChannel(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateChannelReply{Id: id}, nil
}

// UpdateChannel 更新登陆渠道
func (s *ChannelService) UpdateChannel(c context.Context, req *pb.UpdateChannelRequest) (*pb.UpdateChannelReply, error) {
	var (
		in  = channel.Channel{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateChannel(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateChannelReply{}, nil
}

// UpdateChannelStatus 更新登陆渠道状态
func (s *ChannelService) UpdateChannelStatus(c context.Context, req *pb.UpdateChannelStatusRequest) (*pb.UpdateChannelStatusReply, error) {
	return &pb.UpdateChannelStatusReply{}, s.uc.UpdateChannelStatus(kratosx.MustContext(c), req.Id, req.Status)
}

// DeleteChannel 删除登陆渠道
func (s *ChannelService) DeleteChannel(c context.Context, req *pb.DeleteChannelRequest) (*pb.DeleteChannelReply, error) {
	total, err := s.uc.DeleteChannel(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteChannelReply{Total: total}, nil
}
