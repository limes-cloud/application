package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "github.com/limes-cloud/usercenter/api/usercenter/app/v1"
	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	"github.com/limes-cloud/usercenter/internal/biz/app"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/data"
)

type AppService struct {
	pb.UnimplementedAppServer
	uc *app.UseCase
}

func NewAppService(conf *conf.Config) *AppService {
	return &AppService{
		uc: app.NewUseCase(conf, data.NewAppRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewAppService(c)
		pb.RegisterAppHTTPServer(hs, srv)
		pb.RegisterAppServer(gs, srv)
	})
}

// GetApp 获取指定的应用信息
func (s *AppService) GetApp(c context.Context, req *pb.GetAppRequest) (*pb.GetAppReply, error) {
	var (
		in  = app.GetAppRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.GetApp(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.GetAppReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListApp 获取应用信息列表
func (s *AppService) ListApp(c context.Context, req *pb.ListAppRequest) (*pb.ListAppReply, error) {
	var (
		in  = app.ListAppRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListApp(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListAppReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateApp 创建应用信息
func (s *AppService) CreateApp(c context.Context, req *pb.CreateAppRequest) (*pb.CreateAppReply, error) {
	var (
		in  = app.App{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	for _, item := range req.ChannelIds {
		in.AppChannels = append(in.AppChannels, &app.AppChannel{
			ChannelId: item,
		})
	}

	for _, item := range req.FieldIds {
		in.AppFields = append(in.AppFields, &app.AppField{
			FieldId: item,
		})
	}

	id, err := s.uc.CreateApp(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAppReply{Id: id}, nil
}

// UpdateApp 更新应用信息
func (s *AppService) UpdateApp(c context.Context, req *pb.UpdateAppRequest) (*pb.UpdateAppReply, error) {
	var (
		in  = app.App{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	for _, item := range req.ChannelIds {
		in.AppChannels = append(in.AppChannels, &app.AppChannel{
			ChannelId: item,
		})
	}

	for _, item := range req.FieldIds {
		in.AppFields = append(in.AppFields, &app.AppField{
			FieldId: item,
		})
	}

	if err := s.uc.UpdateApp(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateAppReply{}, nil
}

// UpdateAppStatus 更新应用信息状态
func (s *AppService) UpdateAppStatus(c context.Context, req *pb.UpdateAppStatusRequest) (*pb.UpdateAppStatusReply, error) {
	return &pb.UpdateAppStatusReply{}, s.uc.UpdateAppStatus(kratosx.MustContext(c), &app.UpdateAppStatusRequest{
		Id:          req.Id,
		Status:      req.Status,
		DisableDesc: req.DisableDesc,
	})
}

// DeleteApp 删除应用信息
func (s *AppService) DeleteApp(c context.Context, req *pb.DeleteAppRequest) (*pb.DeleteAppReply, error) {
	if err := s.uc.DeleteApp(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteAppReply{}, nil
}
