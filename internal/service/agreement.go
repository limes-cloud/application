package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"

	pb "github.com/limes-cloud/user-center/api/agreement/v1"
	"github.com/limes-cloud/user-center/api/errors"
	biz "github.com/limes-cloud/user-center/internal/biz/agreement"
	"github.com/limes-cloud/user-center/internal/config"
	data "github.com/limes-cloud/user-center/internal/data/agreement"
)

type AgreementService struct {
	pb.UnimplementedServiceServer
	uc *biz.UseCase
}

func NewAgreement(conf *config.Config) *AgreementService {
	return &AgreementService{
		uc: biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (s *AgreementService) PageContent(ctx context.Context, in *pb.PageContentRequest) (*pb.PageContentReply, error) {
	var req biz.PageContentRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	list, total, err := s.uc.PageContent(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := pb.PageContentReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *AgreementService) GetContent(ctx context.Context, in *pb.GetContentRequest) (*pb.Content, error) {
	content, err := s.uc.GetContent(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}

	reply := pb.Content{}
	if err := copier.Copy(&reply, content); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *AgreementService) AddContent(ctx context.Context, in *pb.AddContentRequest) (*pb.AddContentReply, error) {
	var content biz.Content
	if err := copier.Copy(&content, in); err != nil {
		return nil, errors.Transform()
	}

	id, err := s.uc.AddContent(kratosx.MustContext(ctx), &content)
	if err != nil {
		return nil, err
	}
	return &pb.AddContentReply{Id: id}, nil
}

func (s *AgreementService) UpdateContent(ctx context.Context, in *pb.UpdateContentRequest) (*empty.Empty, error) {
	var content biz.Content
	if err := copier.Copy(&content, in); err != nil {
		return nil, errors.Transform()
	}

	return nil, s.uc.UpdateContent(kratosx.MustContext(ctx), &content)
}

func (s *AgreementService) DeleteContent(ctx context.Context, in *pb.DeleteContentRequest) (*empty.Empty, error) {
	return nil, s.uc.DeleteContent(kratosx.MustContext(ctx), in.Id)
}

func (s *AgreementService) PageScene(ctx context.Context, in *pb.PageSceneRequest) (*pb.PageSceneReply, error) {
	var req biz.PageSceneRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	list, total, err := s.uc.PageScene(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := pb.PageSceneReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *AgreementService) GetSceneByKeyword(ctx context.Context, in *pb.GetSceneByKeywordRequest) (*pb.Scene, error) {
	scene, err := s.uc.GetSceneByKeyword(kratosx.MustContext(ctx), in.Keyword)
	if err != nil {
		return nil, err
	}

	reply := pb.Scene{}
	if err := copier.Copy(&reply, scene); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *AgreementService) AddScene(ctx context.Context, in *pb.AddSceneRequest) (*pb.AddSceneReply, error) {
	var scene biz.Scene
	if err := copier.Copy(&scene, in); err != nil {
		return nil, errors.Transform()
	}

	for _, id := range in.ContentIds {
		scene.SceneContents = append(scene.SceneContents, &biz.SceneContent{
			ContentID: id,
		})
	}

	id, err := s.uc.AddScene(kratosx.MustContext(ctx), &scene)
	if err != nil {
		return nil, err
	}
	return &pb.AddSceneReply{Id: id}, nil
}

func (s *AgreementService) UpdateScene(ctx context.Context, in *pb.UpdateSceneRequest) (*empty.Empty, error) {
	var scene biz.Scene
	if err := copier.Copy(&scene, in); err != nil {
		return nil, errors.Transform()
	}

	for _, id := range in.ContentIds {
		scene.SceneContents = append(scene.SceneContents, &biz.SceneContent{
			ContentID: id,
		})
	}

	return nil, s.uc.UpdateScene(kratosx.MustContext(ctx), &scene)
}

func (s *AgreementService) DeleteScene(ctx context.Context, in *pb.DeleteSceneRequest) (*empty.Empty, error) {
	return nil, s.uc.DeleteScene(kratosx.MustContext(ctx), in.Id)
}
