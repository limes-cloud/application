package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	resourceV1 "github.com/limes-cloud/resource/api/v1"

	pb "github.com/limes-cloud/user-center/api/app/v1"
	"github.com/limes-cloud/user-center/api/errors"
	biz "github.com/limes-cloud/user-center/internal/biz/app"
	"github.com/limes-cloud/user-center/internal/config"
	data "github.com/limes-cloud/user-center/internal/data/app"
	"github.com/limes-cloud/user-center/internal/pkg/service"
)

type AppService struct {
	pb.UnimplementedServiceServer
	uc   *biz.UseCase
	conf *config.Config
}

func NewApp(conf *config.Config) *AppService {
	return &AppService{
		conf: conf,
		uc:   biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (s *AppService) PageApp(ctx context.Context, in *pb.PageAppRequest) (*pb.PageAppReply, error) {
	var req biz.PageAppRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	list, total, err := s.uc.Page(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := pb.PageAppReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.Transform()
	}

	// 请求资源中心,错了直接忽略，不影响主流程
	resource, err := service.NewResource(ctx, s.conf.Service.Resource)
	if err == nil {
		for ind, item := range reply.List {
			reply.List[ind].Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.Logo})
		}
	}

	return &reply, nil
}

func (s *AppService) GetAppByKeyword(ctx context.Context, in *pb.GetAppByKeywordRequest) (*pb.App, error) {
	app, err := s.uc.GetByKeyword(kratosx.MustContext(ctx), in.Keyword)
	if err != nil {
		return nil, err
	}

	reply := pb.App{}
	if err := copier.Copy(&reply, app); err != nil {
		return nil, errors.Transform()
	}

	// 请求资源中心,错了直接忽略，不影响主流程
	resource, err := service.NewResource(ctx, s.conf.Service.Resource)
	if err == nil {
		reply.Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: app.Logo})
		for ind, channel := range reply.Channels {
			if res, _ := resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: channel.Logo}); res != nil {
				reply.Channels[ind].Logo = res.Src
			}
		}
	}

	return &reply, nil
}

func (s *AppService) AddApp(ctx context.Context, in *pb.AddAppRequest) (*pb.AddAppReply, error) {
	var app biz.App
	if err := copier.Copy(&app, in); err != nil {
		return nil, errors.Transform()
	}

	for _, id := range in.ChannelIds {
		app.AppChannels = append(app.AppChannels, &biz.AppChannel{
			ChannelID: id,
		})
	}
	for _, id := range in.FieldIds {
		app.AppFields = append(app.AppFields, &biz.AppField{
			FieldID: id,
		})
	}

	id, err := s.uc.Add(kratosx.MustContext(ctx), &app)
	if err != nil {
		return nil, err
	}
	return &pb.AddAppReply{Id: id}, nil
}

func (s *AppService) UpdateApp(ctx context.Context, in *pb.UpdateAppRequest) (*empty.Empty, error) {
	var app biz.App
	if err := copier.Copy(&app, in); err != nil {
		return nil, errors.Transform()
	}

	for _, id := range in.ChannelIds {
		app.AppChannels = append(app.AppChannels, &biz.AppChannel{
			ChannelID: id,
		})
	}
	for _, id := range in.FieldIds {
		app.AppFields = append(app.AppFields, &biz.AppField{
			FieldID: id,
		})
	}

	return nil, s.uc.Update(kratosx.MustContext(ctx), &app)
}

func (s *AppService) DeleteApp(ctx context.Context, in *pb.DeleteAppRequest) (*empty.Empty, error) {
	return nil, s.uc.Delete(kratosx.MustContext(ctx), in.Id)
}
