package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	pb "github.com/limes-cloud/usercenter/api/usercenter/feedback/v1"
	"github.com/limes-cloud/usercenter/internal/biz/feedback"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/data"
)

type FeedbackService struct {
	pb.UnimplementedFeedbackServer
	uc *feedback.UseCase
}

func NewFeedbackService(conf *conf.Config) *FeedbackService {
	return &FeedbackService{
		uc: feedback.NewUseCase(conf, data.NewFeedbackRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewFeedbackService(c)
		pb.RegisterFeedbackHTTPServer(hs, srv)
		pb.RegisterFeedbackServer(gs, srv)
	})
}

// ListFeedbackCategory 获取反馈建议分类列表
func (s *FeedbackService) ListFeedbackCategory(c context.Context, req *pb.ListFeedbackCategoryRequest) (*pb.ListFeedbackCategoryReply, error) {
	var (
		in  = feedback.ListFeedbackCategoryRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListFeedbackCategory(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListFeedbackCategoryReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateFeedbackCategory 创建反馈建议分类
func (s *FeedbackService) CreateFeedbackCategory(c context.Context, req *pb.CreateFeedbackCategoryRequest) (*pb.CreateFeedbackCategoryReply, error) {
	var (
		in  = feedback.FeedbackCategory{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateFeedbackCategory(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateFeedbackCategoryReply{Id: id}, nil
}

// UpdateFeedbackCategory 更新反馈建议分类
func (s *FeedbackService) UpdateFeedbackCategory(c context.Context, req *pb.UpdateFeedbackCategoryRequest) (*pb.UpdateFeedbackCategoryReply, error) {
	var (
		in  = feedback.FeedbackCategory{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateFeedbackCategory(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateFeedbackCategoryReply{}, nil
}

// DeleteFeedbackCategory 删除反馈建议分类
func (s *FeedbackService) DeleteFeedbackCategory(c context.Context, req *pb.DeleteFeedbackCategoryRequest) (*pb.DeleteFeedbackCategoryReply, error) {
	total, err := s.uc.DeleteFeedbackCategory(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteFeedbackCategoryReply{Total: total}, nil
}

// ListFeedback 获取反馈建议列表
func (s *FeedbackService) ListFeedback(c context.Context, req *pb.ListFeedbackRequest) (*pb.ListFeedbackReply, error) {
	var (
		in  = feedback.ListFeedbackRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListFeedback(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListFeedbackReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateFeedback 创建反馈建议
func (s *FeedbackService) CreateFeedback(c context.Context, req *pb.CreateFeedbackRequest) (*pb.CreateFeedbackReply, error) {
	var (
		in  = feedback.Feedback{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateFeedback(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateFeedbackReply{Id: id}, nil
}

// DeleteFeedback 删除反馈建议
func (s *FeedbackService) DeleteFeedback(c context.Context, req *pb.DeleteFeedbackRequest) (*pb.DeleteFeedbackReply, error) {
	total, err := s.uc.DeleteFeedback(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteFeedbackReply{Total: total}, nil
}

// UpdateFeedback 更新反馈建议
func (s *FeedbackService) UpdateFeedback(c context.Context, req *pb.UpdateFeedbackRequest) (*pb.UpdateFeedbackReply, error) {
	var (
		in  = feedback.Feedback{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateFeedback(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateFeedbackReply{}, nil
}
