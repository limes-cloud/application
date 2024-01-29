package service

import (
	"context"

	resourceV1 "github.com/limes-cloud/resource/api/v1"

	"github.com/limes-cloud/user-center/pkg/service"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/internal/biz"
)

func (s *Service) AllChannel(ctx context.Context, in *empty.Empty) (*v1.AllChannelReply, error) {
	list, err := s.channel.All(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := v1.AllChannelReply{}
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

func (s *Service) AddChannel(ctx context.Context, in *v1.AddChannelRequest) (*v1.AddChannelReply, error) {
	var channel biz.Channel
	if err := copier.Copy(&channel, in); err != nil {
		return nil, v1.TransformError()
	}

	id, err := s.channel.Add(kratosx.MustContext(ctx), &channel)
	if err != nil {
		return nil, err
	}
	return &v1.AddChannelReply{Id: id}, nil
}

func (s *Service) UpdateChannel(ctx context.Context, in *v1.UpdateChannelRequest) (*empty.Empty, error) {
	var channel biz.Channel
	if err := copier.Copy(&channel, in); err != nil {
		return nil, v1.TransformError()
	}
	return nil, s.channel.Update(kratosx.MustContext(ctx), &channel)
}

func (s *Service) DeleteChannel(ctx context.Context, in *v1.DeleteChannelRequest) (*empty.Empty, error) {
	return nil, s.channel.Delete(kratosx.MustContext(ctx), in.Id)
}
