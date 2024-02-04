package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"

	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/pkg/util"
)

func (s *Service) GetAppInterfaceTree(ctx context.Context, in *v1.GetAppInterfaceTreeRequest) (*v1.GetAppInterfaceTreeReply, error) {
	tree, err := s.appInterface.Tree(kratosx.MustContext(ctx), in.AppId)
	if err != nil {
		return nil, err
	}

	reply := v1.GetAppInterfaceTreeReply{}

	if err := util.Transform(tree, &reply.List); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) AddAppInterface(ctx context.Context, in *v1.AddAppInterfaceRequest) (*v1.AddAppInterfaceReply, error) {
	var appInterface biz.AppInterface
	if err := copier.Copy(&appInterface, in); err != nil {
		return nil, v1.TransformError()
	}

	id, err := s.appInterface.Add(kratosx.MustContext(ctx), &appInterface)
	if err != nil {
		return nil, err
	}
	return &v1.AddAppInterfaceReply{Id: id}, nil
}

func (s *Service) UpdateAppInterface(ctx context.Context, in *v1.UpdateAppInterfaceRequest) (*empty.Empty, error) {
	var appInterface biz.AppInterface
	if err := copier.Copy(&appInterface, in); err != nil {
		return nil, v1.TransformError()
	}

	return nil, s.appInterface.Update(kratosx.MustContext(ctx), &appInterface)
}

func (s *Service) DeleteAppInterface(ctx context.Context, in *v1.DeleteAppInterfaceRequest) (*empty.Empty, error) {
	return nil, s.appInterface.Delete(kratosx.MustContext(ctx), in.Id)
}
