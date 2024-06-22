package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "github.com/limes-cloud/usercenter/api/usercenter/auth/v1"
	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	"github.com/limes-cloud/usercenter/internal/biz/auth"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/data"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	uc *auth.UseCase
}

func NewAuthService(conf *conf.Config) *AuthService {
	return &AuthService{
		uc: auth.NewUseCase(conf, data.NewAuthRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewAuthService(c)
		pb.RegisterAuthHTTPServer(hs, srv)
		pb.RegisterAuthServer(gs, srv)
	})
}

func (s *AuthService) Auth(ctx context.Context, _ *empty.Empty) (*pb.AuthReply, error) {
	res, err := s.uc.Auth(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	return &pb.AuthReply{
		UserId:     res.UserId,
		AppKeyword: res.AppKeyword,
	}, nil
}

// ListAuth 获取应用授权信息列表
func (s *AuthService) ListAuth(c context.Context, req *pb.ListAuthRequest) (*pb.ListAuthReply, error) {
	var (
		in  = auth.ListAuthRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListAuth(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListAuthReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateAuth 创建应用授权信息
func (s *AuthService) CreateAuth(c context.Context, req *pb.CreateAuthRequest) (*pb.CreateAuthReply, error) {
	var (
		in  = auth.Auth{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateAuth(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAuthReply{Id: id}, nil
}

// UpdateAuthStatus 更新应用授权信息状态
func (s *AuthService) UpdateAuthStatus(c context.Context, req *pb.UpdateAuthStatusRequest) (*pb.UpdateAuthStatusReply, error) {
	return &pb.UpdateAuthStatusReply{}, s.uc.UpdateAuthStatus(kratosx.MustContext(c), &auth.UpdateAuthStatusRequest{
		Id:          req.Id,
		Status:      req.Status,
		DisableDesc: req.DisableDesc,
	})
}

// DeleteAuth 删除应用授权信息
func (s *AuthService) DeleteAuth(c context.Context, req *pb.DeleteAuthRequest) (*pb.DeleteAuthReply, error) {
	if err := s.uc.DeleteAuth(kratosx.MustContext(c), req.UserId, req.AppId); err != nil {
		return nil, err
	}
	return &pb.DeleteAuthReply{}, nil
}

// ListOAuth 获取应用授权信息列表
func (s *AuthService) ListOAuth(c context.Context, req *pb.ListOAuthRequest) (*pb.ListOAuthReply, error) {
	var (
		in  = auth.ListOAuthRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListOAuth(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListOAuthReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// DeleteOAuth 删除应用授权信息
func (s *AuthService) DeleteOAuth(c context.Context, req *pb.DeleteOAuthRequest) (*pb.DeleteOAuthReply, error) {
	if err := s.uc.DeleteOAuth(kratosx.MustContext(c), req.UserId, req.ChannelId); err != nil {
		return nil, err
	}
	return &pb.DeleteOAuthReply{}, nil
}

func (s *AuthService) GenAuthCaptcha(c context.Context, req *pb.GenAuthCaptchaRequest) (*pb.GenAuthCaptchaReply, error) {
	resp, err := s.uc.GenAuthCaptcha(kratosx.MustContext(c), &auth.GenAuthCaptchaRequest{
		Type:  req.Type,
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GenAuthCaptchaReply{
		Id:     resp.Id,
		Expire: uint32(resp.Expire),
		Base64: resp.Base64,
	}, nil
}

func (s *AuthService) OAuthLogin(c context.Context, req *pb.OAuthLoginRequest) (*pb.OAuthLoginReply, error) {
	resp, err := s.uc.OAuthLogin(kratosx.MustContext(c), &auth.OAuthLoginRequest{
		App:     req.App,
		Code:    req.Code,
		Channel: req.Channel,
	})
	if err != nil {
		return nil, err
	}
	return &pb.OAuthLoginReply{
		IsBind:   resp.IsBind,
		OAuthUid: resp.OAuthUid,
		Token:    resp.Token,
		Expire:   resp.Expire,
	}, nil
}

func (s *AuthService) EmailLogin(c context.Context, req *pb.EmailLoginRequest) (*pb.EmailLoginReply, error) {
	resp, err := s.uc.EmailLogin(kratosx.MustContext(c), &auth.EmailLoginRequest{
		Email:     req.Email,
		Captcha:   req.Captcha,
		CaptchaId: req.CaptchaId,
		App:       req.App,
	})
	if err != nil {
		return nil, err
	}
	return &pb.EmailLoginReply{
		Token:  resp.Token,
		Expire: resp.Expire,
	}, nil
}

func (s *AuthService) PasswordLogin(c context.Context, req *pb.PasswordLoginRequest) (*pb.PasswordLoginReply, error) {
	resp, err := s.uc.PasswordLogin(kratosx.MustContext(c), &auth.PasswordLoginRequest{
		Username:  req.Username,
		Password:  req.Password,
		Captcha:   req.Captcha,
		CaptchaId: req.CaptchaId,
		App:       req.App,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PasswordLoginReply{
		Token:  resp.Token,
		Expire: resp.Expire,
	}, nil
}

func (s *AuthService) EmailRegister(c context.Context, req *pb.EmailRegisterRequest) (*pb.EmailRegisterReply, error) {
	resp, err := s.uc.EmailRegister(kratosx.MustContext(c), &auth.EmailRegisterRequest{
		Email:     req.Email,
		Captcha:   req.Captcha,
		CaptchaId: req.CaptchaId,
		App:       req.App,
	})
	if err != nil {
		return nil, err
	}
	return &pb.EmailRegisterReply{
		Token:  resp.Token,
		Expire: resp.Expire,
	}, nil
}

func (s *AuthService) PasswordRegister(c context.Context, req *pb.PasswordRegisterRequest) (*pb.PasswordRegisterReply, error) {
	resp, err := s.uc.PasswordRegister(kratosx.MustContext(c), &auth.PasswordRegisterRequest{
		Username:  req.Username,
		Password:  req.Password,
		Captcha:   req.Captcha,
		CaptchaId: req.CaptchaId,
		App:       req.App,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PasswordRegisterReply{
		Token:  resp.Token,
		Expire: resp.Expire,
	}, nil
}

func (s *AuthService) EmailBind(c context.Context, req *pb.EmailBindRequest) (*pb.EmailBindReply, error) {
	resp, err := s.uc.EmailBind(kratosx.MustContext(c), &auth.EmailBindRequest{
		Email:     req.Email,
		Captcha:   req.Captcha,
		CaptchaId: req.CaptchaId,
		App:       req.App,
		OAuthUid:  req.OAuthUid,
	})
	if err != nil {
		return nil, err
	}
	return &pb.EmailBindReply{
		Token:  resp.Token,
		Expire: resp.Expire,
	}, nil
}

func (s *AuthService) PasswordBind(c context.Context, req *pb.PasswordBindRequest) (*pb.PasswordBindReply, error) {
	resp, err := s.uc.PasswordBind(kratosx.MustContext(c), &auth.PasswordBindRequest{
		Username:  req.Username,
		Password:  req.Password,
		Captcha:   req.Captcha,
		CaptchaId: req.CaptchaId,
		App:       req.App,
		OAuthUid:  req.OAuthUid,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PasswordBindReply{
		Token:  resp.Token,
		Expire: resp.Expire,
	}, nil
}
