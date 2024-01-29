package service

import (
	"context"

	"github.com/limes-cloud/user-center/internal/biz/types"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	v1 "github.com/limes-cloud/user-center/api/v1"
)

func (s *Service) AllLoginPlatform(_ context.Context, _ *empty.Empty) (*v1.AllLoginPlatformReply, error) {
	list := s.auth.LoginPlatform()

	reply := v1.AllLoginPlatformReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) LoginImageCaptcha(ctx context.Context, _ *empty.Empty) (*v1.ImageCaptchaReply, error) {
	res, err := s.auth.LoginImageCaptcha(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := v1.ImageCaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) LoginByPassword(ctx context.Context, in *v1.LoginByPasswordRequest) (*v1.LoginReply, error) {
	var req types.LoginByPasswordRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	token, err := s.auth.LoginByPassword(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	return &v1.LoginReply{
		Token: token,
	}, nil
}

func (s *Service) BindImageCaptcha(ctx context.Context, _ *empty.Empty) (*v1.ImageCaptchaReply, error) {
	res, err := s.auth.BindImageCaptcha(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := v1.ImageCaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) RegisterImageCaptcha(ctx context.Context, _ *empty.Empty) (*v1.ImageCaptchaReply, error) {
	res, err := s.auth.RegisterImageCaptcha(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := v1.ImageCaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) RegisterUsernameCheck(ctx context.Context, in *v1.RegisterUsernameCheckRequest) (*v1.RegisterUsernameCheckReply, error) {
	return &v1.RegisterUsernameCheckReply{
		Allow: s.auth.RegisterUsernameCheck(kratosx.MustContext(ctx), in.Username),
	}, nil
}

func (s *Service) RegisterByPassword(ctx context.Context, in *v1.RegisterByPasswordRequest) (*v1.RegisterReply, error) {
	var req types.RegisterByPasswordRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	info, err := s.auth.RegisterByPassword(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	return &v1.RegisterReply{
		Id:    info.ID,
		Token: info.Token,
	}, nil
}

func (s *Service) LoginEmailCaptcha(ctx context.Context, in *v1.LoginEmailCaptchaRequest) (*v1.EmailCaptchaReply, error) {
	res, err := s.auth.LoginEmailCaptcha(kratosx.MustContext(ctx), in.Email)
	if err != nil {
		return nil, err
	}

	reply := v1.EmailCaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) LoginByEmail(ctx context.Context, in *v1.LoginByEmailRequest) (*v1.LoginReply, error) {
	var req types.LoginByEmailRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	token, err := s.auth.LoginByEmail(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	return &v1.LoginReply{
		Token: token,
	}, nil
}

func (s *Service) BindEmailCaptcha(ctx context.Context, in *v1.BindEmailCaptchaRequest) (*v1.EmailCaptchaReply, error) {
	res, err := s.auth.BindEmailCaptcha(kratosx.MustContext(ctx), in.Email)
	if err != nil {
		return nil, err
	}

	reply := v1.EmailCaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) RegisterEmailCaptcha(ctx context.Context, in *v1.RegisterEmailCaptchaRequest) (*v1.EmailCaptchaReply, error) {
	res, err := s.auth.RegisterEmailCaptcha(kratosx.MustContext(ctx), in.Email)
	if err != nil {
		return nil, err
	}

	reply := v1.EmailCaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) RegisterByEmail(ctx context.Context, in *v1.RegisterByEmailRequest) (*v1.RegisterReply, error) {
	var req types.RegisterByEmailRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	info, err := s.auth.RegisterByEmail(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	return &v1.RegisterReply{
		Id:    info.ID,
		Token: info.Token,
	}, nil
}

func (s *Service) Auth(ctx context.Context, in *v1.AuthRequest) (*empty.Empty, error) {
	return nil, s.auth.Auth(kratosx.MustContext(ctx), in.AppId)
}
