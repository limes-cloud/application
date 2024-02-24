package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/user-center/api/errors"
	pb "github.com/limes-cloud/user-center/api/field/v1"
	biz "github.com/limes-cloud/user-center/internal/biz/field"
	"github.com/limes-cloud/user-center/internal/config"
	data "github.com/limes-cloud/user-center/internal/data/field"
)

type FieldService struct {
	pb.UnimplementedServiceServer
	uc   *biz.UseCase
	conf *config.Config
}

func NewField(conf *config.Config) *FieldService {
	return &FieldService{
		conf: conf,
		uc:   biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (s *FieldService) AllFieldType(_ context.Context, _ *empty.Empty) (*pb.AllFieldTypeReply, error) {
	list := s.uc.TypeList()
	reply := pb.AllFieldTypeReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *FieldService) PageField(ctx context.Context, in *pb.PageFieldRequest) (*pb.PageFieldReply, error) {
	var req biz.PageFieldRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	list, total, err := s.uc.Page(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := pb.PageFieldReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *FieldService) AddField(ctx context.Context, in *pb.AddFieldRequest) (*pb.AddFieldReply, error) {
	var field biz.Field
	if err := copier.Copy(&field, in); err != nil {
		return nil, errors.Transform()
	}
	id, err := s.uc.Add(kratosx.MustContext(ctx), &field)
	if err != nil {
		return nil, err
	}
	return &pb.AddFieldReply{Id: id}, nil
}

func (s *FieldService) UpdateField(ctx context.Context, in *pb.UpdateFieldRequest) (*empty.Empty, error) {
	var field biz.Field
	if err := copier.Copy(&field, in); err != nil {
		return nil, errors.Transform()
	}

	return nil, s.uc.Update(kratosx.MustContext(ctx), &field)
}

func (s *FieldService) DeleteField(ctx context.Context, in *pb.DeleteFieldRequest) (*empty.Empty, error) {
	return nil, s.uc.Delete(kratosx.MustContext(ctx), in.Id)
}
