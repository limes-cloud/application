package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	pb "github.com/limes-cloud/usercenter/api/usercenter/field/v1"
	"github.com/limes-cloud/usercenter/internal/biz/field"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/data"
)

type FieldService struct {
	pb.UnimplementedFieldServer
	uc *field.UseCase
}

func NewFieldService(conf *conf.Config) *FieldService {
	return &FieldService{
		uc: field.NewUseCase(conf, data.NewFieldRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewFieldService(c)
		pb.RegisterFieldHTTPServer(hs, srv)
		pb.RegisterFieldServer(gs, srv)
	})
}

// ListFieldType 获取字段类型列表
func (s *FieldService) ListFieldType(c context.Context, req *pb.ListFieldTypeRequest) (*pb.ListFieldTypeReply, error) {
	var (
		reply = pb.ListFieldTypeReply{}
	)

	list := s.uc.ListFieldType()
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListFieldTypeReply_Type{
			Name: item.Name,
			Type: item.Type,
		})
	}

	return &reply, nil
}

// ListField 获取用户字段列表
func (s *FieldService) ListField(c context.Context, req *pb.ListFieldRequest) (*pb.ListFieldReply, error) {
	var (
		in  = field.ListFieldRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListField(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListFieldReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateField 创建用户字段
func (s *FieldService) CreateField(c context.Context, req *pb.CreateFieldRequest) (*pb.CreateFieldReply, error) {
	var (
		in  = field.Field{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateField(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateFieldReply{Id: id}, nil
}

// UpdateField 更新用户字段
func (s *FieldService) UpdateField(c context.Context, req *pb.UpdateFieldRequest) (*pb.UpdateFieldReply, error) {
	var (
		in  = field.Field{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateField(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateFieldReply{}, nil
}

// UpdateFieldStatus 更新用户字段状态
func (s *FieldService) UpdateFieldStatus(c context.Context, req *pb.UpdateFieldStatusRequest) (*pb.UpdateFieldStatusReply, error) {
	return &pb.UpdateFieldStatusReply{}, s.uc.UpdateFieldStatus(kratosx.MustContext(c), req.Id, req.Status)
}

// DeleteField 删除用户字段
func (s *FieldService) DeleteField(c context.Context, req *pb.DeleteFieldRequest) (*pb.DeleteFieldReply, error) {
	total, err := s.uc.DeleteField(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteFieldReply{Total: total}, nil
}
