package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"

	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/internal/biz/types"
	"github.com/limes-cloud/user-center/pkg/md"
)

func (s *Service) CurrentExtraField(ctx context.Context, _ *empty.Empty) (*v1.CurrentExtraFieldReply, error) {
	kCtx := kratosx.MustContext(ctx)
	list, err := s.extraField.AppFields(kCtx, md.AppID(kCtx))
	if err != nil {
		return nil, err
	}
	reply := v1.CurrentExtraFieldReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) AllExtraFieldType(_ context.Context, _ *empty.Empty) (*v1.AllExtraFieldTypeReply, error) {
	list := s.extraField.Types()
	reply := v1.AllExtraFieldTypeReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) PageExtraField(ctx context.Context, in *v1.PageExtraFieldRequest) (*v1.PageExtraFieldReply, error) {
	var req types.PageExtraFieldRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	list, total, err := s.extraField.Page(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageExtraFieldReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) AddExtraField(ctx context.Context, in *v1.AddExtraFieldRequest) (*v1.AddExtraFieldReply, error) {
	var extraField biz.ExtraField
	if err := copier.Copy(&extraField, in); err != nil {
		return nil, v1.TransformError()
	}
	id, err := s.extraField.Add(kratosx.MustContext(ctx), &extraField)
	if err != nil {
		return nil, err
	}
	return &v1.AddExtraFieldReply{Id: id}, nil
}

func (s *Service) UpdateExtraField(ctx context.Context, in *v1.UpdateExtraFieldRequest) (*empty.Empty, error) {
	var extraField biz.ExtraField
	if err := copier.Copy(&extraField, in); err != nil {
		return nil, v1.TransformError()
	}

	return nil, s.extraField.Update(kratosx.MustContext(ctx), &extraField)
}

func (s *Service) DeleteExtraField(ctx context.Context, in *v1.DeleteExtraFieldRequest) (*empty.Empty, error) {
	return nil, s.extraField.Delete(kratosx.MustContext(ctx), in.Id)
}
