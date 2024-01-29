package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/internal/biz/types"
)

func (s *Service) PageAgreement(ctx context.Context, in *v1.PageAgreementRequest) (*v1.PageAgreementReply, error) {
	var req types.PageAgreementRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	list, total, err := s.agreement.Page(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageAgreementReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) GetAgreement(ctx context.Context, in *v1.GetAgreementRequest) (*v1.Agreement, error) {
	agreement, err := s.agreement.Get(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}

	reply := v1.Agreement{}
	if err := copier.Copy(&reply, agreement); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) AddAgreement(ctx context.Context, in *v1.AddAgreementRequest) (*v1.AddAgreementReply, error) {
	var agreement biz.Agreement
	if err := copier.Copy(&agreement, in); err != nil {
		return nil, v1.TransformError()
	}

	id, err := s.agreement.Add(kratosx.MustContext(ctx), &agreement)
	if err != nil {
		return nil, err
	}
	return &v1.AddAgreementReply{Id: id}, nil
}

func (s *Service) UpdateAgreement(ctx context.Context, in *v1.UpdateAgreementRequest) (*empty.Empty, error) {
	var agreement biz.Agreement
	if err := copier.Copy(&agreement, in); err != nil {
		return nil, v1.TransformError()
	}

	return nil, s.agreement.Update(kratosx.MustContext(ctx), &agreement)
}

func (s *Service) DeleteAgreement(ctx context.Context, in *v1.DeleteAgreementRequest) (*empty.Empty, error) {
	return nil, s.agreement.Delete(kratosx.MustContext(ctx), in.Id)
}
