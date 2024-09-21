package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"

	pb "github.com/limes-cloud/application/api/application/auth/v1"
	"github.com/limes-cloud/application/internal/conf"
	"github.com/limes-cloud/application/internal/domain/entity"
	"github.com/limes-cloud/application/internal/domain/service"
	"github.com/limes-cloud/application/internal/infra/dbs"
	"github.com/limes-cloud/application/internal/infra/rpc"
	"github.com/limes-cloud/application/internal/types"
)

type Auth struct {
	pb.UnimplementedAuthServer
	srv *service.Auth
}

func NewAuth(conf *conf.Config) *Auth {
	return &Auth{
		srv: service.NewAuth(
			conf,
			dbs.NewAuth(),
			dbs.NewUser(),
			dbs.NewApp(),
			dbs.NewChannel(),
			rpc.NewPermission(),
			rpc.NewFile(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewAuth(c)
		pb.RegisterAuthHTTPServer(hs, srv)
		pb.RegisterAuthServer(gs, srv)
	})
}

func (s *Auth) Auth(ctx context.Context, _ *empty.Empty) (*pb.AuthReply, error) {
	res, err := s.srv.Auth(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	return &pb.AuthReply{
		UserId:     res.UserId,
		AppKeyword: res.AppKeyword,
	}, nil
}

// ListAuth 获取应用授权信息列表
func (s *Auth) ListAuth(c context.Context, req *pb.ListAuthRequest) (*pb.ListAuthReply, error) {
	list, total, err := s.srv.ListAuth(kratosx.MustContext(c), &types.ListAuthRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		UserId:   req.UserId,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListAuthReply{Total: total}
	for _, item := range list {
		if item.App == nil {
			continue
		}
		reply.List = append(reply.List, &pb.ListAuthReply_Auth{
			Id:          item.Id,
			AppId:       item.AppId,
			Status:      item.Status,
			DisableDesc: item.DisableDesc,
			LoggedAt:    uint32(item.LoggedAt),
			ExpiredAt:   uint32(item.ExpiredAt),
			CreatedAt:   uint32(item.CreatedAt),
			App: &pb.ListAuthReply_App{
				Logo:    item.App.Logo,
				Keyword: item.App.Keyword,
				Name:    item.App.Name,
			},
		})
	}
	return &reply, nil
}

// CreateAuth 创建应用授权信息
func (s *Auth) CreateAuth(c context.Context, req *pb.CreateAuthRequest) (*pb.CreateAuthReply, error) {
	id, err := s.srv.CreateAuth(kratosx.MustContext(c), &entity.Auth{
		UserId: req.UserId,
		AppId:  req.AppId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAuthReply{Id: id}, nil
}

// UpdateAuthStatus 更新应用授权信息状态
func (s *Auth) UpdateAuthStatus(c context.Context, req *pb.UpdateAuthStatusRequest) (*pb.UpdateAuthStatusReply, error) {
	return &pb.UpdateAuthStatusReply{}, s.srv.UpdateAuthStatus(kratosx.MustContext(c), &types.UpdateAuthStatusRequest{
		Id:          req.Id,
		Status:      req.Status,
		DisableDesc: req.DisableDesc,
	})
}

// DeleteAuth 删除应用授权信息
func (s *Auth) DeleteAuth(c context.Context, req *pb.DeleteAuthRequest) (*pb.DeleteAuthReply, error) {
	if err := s.srv.DeleteAuth(kratosx.MustContext(c), req.UserId, req.AppId); err != nil {
		return nil, err
	}
	return &pb.DeleteAuthReply{}, nil
}

// ListOAuth 获取应用授权信息列表
func (s *Auth) ListOAuth(c context.Context, req *pb.ListOAuthRequest) (*pb.ListOAuthReply, error) {
	list, total, err := s.srv.ListOAuth(kratosx.MustContext(c), &types.ListOAuthRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		UserId:   req.UserId,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListOAuthReply{Total: total}
	for _, item := range list {
		if item.Channel == nil {
			continue
		}
		reply.List = append(reply.List, &pb.ListOAuthReply_OAuth{
			Id:        item.Id,
			ChannelId: item.ChannelId,
			LoggedAt:  uint32(item.LoggedAt),
			ExpiredAt: uint32(item.ExpiredAt),
			CreatedAt: uint32(item.CreatedAt),
			Channel: &pb.ListOAuthReply_Channel{
				Logo:    item.Channel.Logo,
				Keyword: item.Channel.Keyword,
				Name:    item.Channel.Name,
			},
		})
	}
	return &reply, nil
}

// DeleteOAuth 删除应用授权信息
func (s *Auth) DeleteOAuth(c context.Context, req *pb.DeleteOAuthRequest) (*pb.DeleteOAuthReply, error) {
	if err := s.srv.DeleteOAuth(kratosx.MustContext(c), req.UserId, req.ChannelId); err != nil {
		return nil, err
	}
	return &pb.DeleteOAuthReply{}, nil
}

func (s *Auth) GenAuthCaptcha(c context.Context, req *pb.GenAuthCaptchaRequest) (*pb.GenAuthCaptchaReply, error) {
	resp, err := s.srv.GenAuthCaptcha(kratosx.MustContext(c), &types.GenAuthCaptchaRequest{
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

func (s *Auth) OAuthLogin(c context.Context, req *pb.OAuthLoginRequest) (*pb.OAuthLoginReply, error) {
	resp, err := s.srv.OAuthLogin(kratosx.MustContext(c), &types.OAuthLoginRequest{
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

func (s *Auth) EmailLogin(c context.Context, req *pb.EmailLoginRequest) (*pb.EmailLoginReply, error) {
	resp, err := s.srv.EmailLogin(kratosx.MustContext(c), &types.EmailLoginRequest{
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

func (s *Auth) PasswordLogin(c context.Context, req *pb.PasswordLoginRequest) (*pb.PasswordLoginReply, error) {
	resp, err := s.srv.PasswordLogin(kratosx.MustContext(c), &types.PasswordLoginRequest{
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

func (s *Auth) EmailRegister(c context.Context, req *pb.EmailRegisterRequest) (*pb.EmailRegisterReply, error) {
	resp, err := s.srv.EmailRegister(kratosx.MustContext(c), &types.EmailRegisterRequest{
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

func (s *Auth) PasswordRegister(c context.Context, req *pb.PasswordRegisterRequest) (*pb.PasswordRegisterReply, error) {
	resp, err := s.srv.PasswordRegister(kratosx.MustContext(c), &types.PasswordRegisterRequest{
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

func (s *Auth) EmailBind(c context.Context, req *pb.EmailBindRequest) (*pb.EmailBindReply, error) {
	resp, err := s.srv.EmailBind(kratosx.MustContext(c), &types.EmailBindRequest{
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

func (s *Auth) PasswordBind(c context.Context, req *pb.PasswordBindRequest) (*pb.PasswordBindReply, error) {
	resp, err := s.srv.PasswordBind(kratosx.MustContext(c), &types.PasswordBindRequest{
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

func (s *Auth) RefreshToken(ctx context.Context, _ *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	resp, err := s.srv.RefreshToken(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	return &pb.RefreshTokenReply{
		Token:  resp.Token,
		Expire: resp.Expire,
	}, nil
}

func (s *Auth) Logout(ctx context.Context, _ *pb.LogoutRequest) (*pb.LogoutReply, error) {
	err := s.srv.Logout(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	return &pb.LogoutReply{}, nil
}
