package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	resourcepb "github.com/limes-cloud/resource/api/file/v1"

	pb "github.com/limes-cloud/user-center/api/channel/v1"
	"github.com/limes-cloud/user-center/api/errors"
	biz "github.com/limes-cloud/user-center/internal/biz/channel"
	"github.com/limes-cloud/user-center/internal/config"
	data "github.com/limes-cloud/user-center/internal/data/channel"
	"github.com/limes-cloud/user-center/internal/pkg/service"
)

type ChannelService struct {
	pb.UnimplementedServiceServer
	uc   *biz.UseCase
	conf *config.Config
}

func NewChannel(conf *config.Config) *ChannelService {
	return &ChannelService{
		conf: conf,
		uc:   biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (s *ChannelService) AllChannel(ctx context.Context, _ *empty.Empty) (*pb.AllChannelReply, error) {
	list, err := s.uc.All(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := pb.AllChannelReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.Transform()
	}

	// 请求资源中心,错了直接忽略，不影响主流程
	resource, err := service.NewResource(ctx)
	if err == nil {
		for ind, item := range reply.List {
			reply.List[ind].Resource, _ = resource.GetFileBySha(ctx, &resourcepb.GetFileByShaRequest{Sha: item.Logo})
		}
	}
	return &reply, nil
}

func (s *ChannelService) GetTypes(_ context.Context, _ *empty.Empty) (*pb.GetTypesReply, error) {
	list, err := s.uc.GetTypes()
	if err != nil {
		return nil, err
	}

	reply := pb.GetTypesReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *ChannelService) AddChannel(ctx context.Context, in *pb.AddChannelRequest) (*pb.AddChannelReply, error) {
	var channel biz.Channel
	if err := copier.Copy(&channel, in); err != nil {
		return nil, errors.Transform()
	}

	id, err := s.uc.Add(kratosx.MustContext(ctx), &channel)
	if err != nil {
		return nil, err
	}
	return &pb.AddChannelReply{Id: id}, nil
}

func (s *ChannelService) UpdateChannel(ctx context.Context, in *pb.UpdateChannelRequest) (*empty.Empty, error) {
	var channel biz.Channel
	if err := copier.Copy(&channel, in); err != nil {
		return nil, errors.Transform()
	}
	return nil, s.uc.Update(kratosx.MustContext(ctx), &channel)
}

func (s *ChannelService) DeleteChannel(ctx context.Context, in *pb.DeleteChannelRequest) (*empty.Empty, error) {
	return nil, s.uc.Delete(kratosx.MustContext(ctx), in.Id)
}
