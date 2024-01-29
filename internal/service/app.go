package service

import (
	"context"

	resourceV1 "github.com/limes-cloud/resource/api/v1"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/internal/biz/types"
	"github.com/limes-cloud/user-center/pkg/service"
)

func (s *Service) PageApp(ctx context.Context, in *v1.PageAppRequest) (*v1.PageAppReply, error) {
	var req types.PageAppRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	list, total, err := s.app.Page(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageAppReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
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

func (s *Service) GetAppByKeyword(ctx context.Context, in *v1.GetAppByKeywordRequest) (*v1.App, error) {
	app, err := s.app.GetByKeyword(kratosx.MustContext(ctx), in.Keyword)
	if err != nil {
		return nil, err
	}

	reply := v1.App{}
	if err := copier.Copy(&reply, app); err != nil {
		return nil, v1.TransformError()
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

func (s *Service) AddApp(ctx context.Context, in *v1.AddAppRequest) (*v1.AddAppReply, error) {
	var app biz.App
	if err := copier.Copy(&app, in); err != nil {
		return nil, v1.TransformError()
	}

	for _, id := range in.ChannelIds {
		app.AppChannels = append(app.AppChannels, &biz.AppChannel{
			ChannelID: id,
		})
	}

	id, err := s.app.Add(kratosx.MustContext(ctx), &app)
	if err != nil {
		return nil, err
	}
	return &v1.AddAppReply{Id: id}, nil
}

func (s *Service) UpdateApp(ctx context.Context, in *v1.UpdateAppRequest) (*empty.Empty, error) {
	var app biz.App
	if err := copier.Copy(&app, in); err != nil {
		return nil, v1.TransformError()
	}

	for _, id := range in.ChannelIds {
		app.AppChannels = append(app.AppChannels, &biz.AppChannel{
			ChannelID: id,
		})
	}

	return nil, s.app.Update(kratosx.MustContext(ctx), &app)
}

func (s *Service) DeleteApp(ctx context.Context, in *v1.DeleteAppRequest) (*empty.Empty, error) {
	return nil, s.app.Delete(kratosx.MustContext(ctx), in.Id)
}
