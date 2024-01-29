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

func (s *Service) PageScene(ctx context.Context, in *v1.PageSceneRequest) (*v1.PageSceneReply, error) {
	var req types.PageSceneRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	list, total, err := s.scene.Page(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageSceneReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) GetSceneByKeyword(ctx context.Context, in *v1.GetSceneByKeywordRequest) (*v1.Scene, error) {
	scene, err := s.scene.GetByKeyword(kratosx.MustContext(ctx), in.Keyword)
	if err != nil {
		return nil, err
	}

	reply := v1.Scene{}
	if err := copier.Copy(&reply, scene); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) AddScene(ctx context.Context, in *v1.AddSceneRequest) (*v1.AddSceneReply, error) {
	var scene biz.Scene
	if err := copier.Copy(&scene, in); err != nil {
		return nil, v1.TransformError()
	}

	for _, id := range in.AgreementIds {
		scene.AgreementScenes = append(scene.AgreementScenes, &biz.AgreementScene{
			AgreementID: id,
		})
	}

	id, err := s.scene.Add(kratosx.MustContext(ctx), &scene)
	if err != nil {
		return nil, err
	}
	return &v1.AddSceneReply{Id: id}, nil
}

func (s *Service) UpdateScene(ctx context.Context, in *v1.UpdateSceneRequest) (*empty.Empty, error) {
	var scene biz.Scene
	if err := copier.Copy(&scene, in); err != nil {
		return nil, v1.TransformError()
	}

	for _, id := range in.AgreementIds {
		scene.AgreementScenes = append(scene.AgreementScenes, &biz.AgreementScene{
			AgreementID: id,
		})
	}

	return nil, s.scene.Update(kratosx.MustContext(ctx), &scene)
}

func (s *Service) DeleteScene(ctx context.Context, in *v1.DeleteSceneRequest) (*empty.Empty, error) {
	return nil, s.scene.Delete(kratosx.MustContext(ctx), in.Id)
}
