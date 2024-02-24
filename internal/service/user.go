package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	resourceV1 "github.com/limes-cloud/resource/api/v1"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/limes-cloud/user-center/api/errors"
	pb "github.com/limes-cloud/user-center/api/user/v1"
	"github.com/limes-cloud/user-center/internal/biz/app"
	fieldbiz "github.com/limes-cloud/user-center/internal/biz/field"
	biz "github.com/limes-cloud/user-center/internal/biz/user"
	"github.com/limes-cloud/user-center/internal/config"
	fielddata "github.com/limes-cloud/user-center/internal/data/field"
	data "github.com/limes-cloud/user-center/internal/data/user"
	"github.com/limes-cloud/user-center/internal/pkg/field"
	"github.com/limes-cloud/user-center/internal/pkg/md"
	"github.com/limes-cloud/user-center/internal/pkg/service"
)

type UserService struct {
	pb.UnimplementedServiceServer
	uc   *biz.UseCase
	fuc  *fieldbiz.UseCase
	conf *config.Config
}

func NewUser(conf *config.Config) *UserService {
	return &UserService{
		conf: conf,
		uc:   biz.NewUseCase(conf, data.NewRepo()),
		fuc:  fieldbiz.NewUseCase(conf, fielddata.NewRepo()),
	}
}

func (s *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	var err error
	var user *biz.User
	switch in.Condition.(type) {
	case *pb.GetUserRequest_Id:
		cond := in.Condition.(*pb.GetUserRequest_Id)
		user, err = s.uc.Get(kratosx.MustContext(ctx), cond.Id)
	case *pb.GetUserRequest_Email:
		cond := in.Condition.(*pb.GetUserRequest_Email)
		user, err = s.uc.GetByEmail(kratosx.MustContext(ctx), cond.Email)
	case *pb.GetUserRequest_Phone:
		cond := in.Condition.(*pb.GetUserRequest_Phone)
		user, err = s.uc.GetByPhone(kratosx.MustContext(ctx), cond.Phone)
	case *pb.GetUserRequest_Username:
		cond := in.Condition.(*pb.GetUserRequest_Username)
		user, err = s.uc.GetByPhone(kratosx.MustContext(ctx), cond.Username)
	default:
		user, err = nil, errors.NotFound()
	}
	if err != nil {
		return nil, err
	}

	return s.transformUserReply(kratosx.MustContext(ctx), user)
}

func (s *UserService) transformUserReply(ctx kratosx.Context, user *biz.User) (*pb.User, error) {
	reply := pb.User{}
	if err := copier.Copy(&reply, user); err != nil {
		return nil, errors.Transform()
	}

	reply.Apps = make([]*pb.User_App, 0)
	reply.Channels = make([]*pb.User_Channel, 0)
	reply.Extra = make(map[string]*structpb.Value)
	reply.ExtraList = []*pb.User_Extra{}

	// 请求资源中心,错了直接忽略，不影响主流程
	resource, rer := service.NewResource(ctx, s.conf.Service.Resource)
	if rer == nil {
		if reply.Avatar != "" {
			reply.Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{
				Sha: reply.Avatar,
			})
		}
	}

	// 组装数据apps
	for _, item := range user.UserApps {
		replyApp := &pb.User_App{
			Id:        item.App.ID,
			Name:      item.App.Name,
			Logo:      item.App.Logo,
			CreatedAt: uint32(item.CreatedAt),
			LoginAt:   uint32(item.LoginAt),
		}
		if rer == nil {
			replyApp.Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.App.Logo})
		}
		reply.Apps = append(reply.Apps, replyApp)
	}

	// 转换授权信息
	for _, item := range user.Auths {
		replyChannel := &pb.User_Channel{
			Id:        item.Channel.ID,
			Name:      item.Channel.Name,
			Logo:      item.Channel.Logo,
			CreatedAt: uint32(item.CreatedAt),
			LoginAt:   uint32(item.LoginAt),
		}
		if rer == nil {
			replyChannel.Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.Channel.Logo})
		}
		reply.Channels = append(reply.Channels, replyChannel)
	}

	// 转换扩展字段
	// 转换extra
	efs := s.fuc.TypeSet(kratosx.MustContext(ctx))
	fir := field.New()

	var userFields []string
	var curApp *app.App

	hasFilter := false
	appId := md.AppID(ctx)
	if appId != 0 {
		for _, item := range user.UserApps {
			if item.App.ID == appId {
				curApp = item.App
			}
		}
	}
	if curApp != nil {
		for _, item := range curApp.Fields {
			userFields = append(userFields, item.Keyword)
		}
		hasFilter = true
	}

	for _, item := range user.UserExtras {
		if hasFilter && !util.InList(userFields, item.Keyword) {
			continue
		}
		if efs[item.Keyword] != nil {
			tp := fir.GetType(efs[item.Keyword].Type)
			reply.Extra[item.Keyword] = tp.ToValue(item.Value)
			reply.ExtraList = append(reply.ExtraList, &pb.User_Extra{
				Name:     efs[item.Keyword].Name,
				Keyword:  efs[item.Keyword].Keyword,
				Type:     efs[item.Keyword].Type,
				TypeName: tp.Name(),
				Value:    reply.Extra[item.Keyword],
			})
		}
	}
	return &reply, nil
}

func (s *UserService) GetSimpleUser(ctx context.Context, in *pb.GetSimpleUserRequest) (*pb.SimpleUser, error) {
	user, err := s.uc.GetBase(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	reply := pb.SimpleUser{}
	if err := copier.Copy(&reply, user); err != nil {
		return nil, errors.Transform()
	}

	resource, err := service.NewResource(ctx, s.conf.Service.Resource)
	if err == nil {
		reply.Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: reply.Avatar})
	}

	return &reply, nil
}

func (s *UserService) GetBaseUser(ctx context.Context, in *pb.GetBaseUserRequest) (*pb.BaseUser, error) {
	user, err := s.uc.GetBase(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	reply := pb.BaseUser{}
	if err := copier.Copy(&reply, user); err != nil {
		return nil, errors.Transform()
	}

	resource, err := service.NewResource(ctx, s.conf.Service.Resource)
	if err == nil {
		reply.Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: reply.Avatar})
	}
	return &reply, nil
}

func (s *UserService) GetCurrentUser(ctx context.Context, _ *empty.Empty) (*pb.User, error) {
	kCtx := kratosx.MustContext(ctx)
	user, err := s.uc.Get(kCtx, md.UserID(kratosx.MustContext(ctx)))
	if err != nil {
		return nil, err
	}
	return s.transformUserReply(kCtx, user)
}

func (s *UserService) UpdateCurrentUser(ctx context.Context, in *pb.UpdateCurrentUserRequest) (*empty.Empty, error) {
	user := biz.User{}
	if err := copier.Copy(&user, in); err != nil {
		return nil, errors.Transform()
	}

	kCtx := kratosx.MustContext(ctx)
	user.ID = md.UserID(kCtx)

	// 转换extra
	efs := s.fuc.TypeSet(kCtx)
	fd := field.New()

	for key, value := range in.Extra {
		if efs[key] == nil {
			continue
		}
		tp := fd.GetType(efs[key].Type)
		user.UserExtras = append(user.UserExtras, &biz.UserExtra{
			Keyword: key,
			Value:   tp.ToString(value),
		})
	}

	return nil, s.uc.Update(kratosx.MustContext(ctx), &user)
}

func (s *UserService) PageUser(ctx context.Context, in *pb.PageUserRequest) (*pb.PageUserReply, error) {
	var req biz.PageUserRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	list, total, err := s.uc.Page(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := pb.PageUserReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.Transform()
	}

	// 请求资源中心,错了直接忽略，不影响主流程
	resource, err := service.NewResource(ctx, s.conf.Service.Resource)
	if err == nil {
		for index, item := range reply.List {
			if item.Avatar != "" {
				reply.List[index].Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{
					Sha: item.Avatar,
				})
			}
		}
	}
	return &reply, nil
}

func (s *UserService) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserReply, error) {
	var user biz.User
	if err := copier.Copy(&user, in); err != nil {
		return nil, errors.Transform()
	}

	id, err := s.uc.Add(kratosx.MustContext(ctx), &user)
	if err != nil {
		return nil, err
	}
	return &pb.AddUserReply{Id: id}, nil
}

//	func (s *UserService) ImportUser(ctx context.Context, in *pb.ImportUserRequest) (*empty.Empty, error) {
//		var users []*biz.User
//		if err := copier.Copy(&users, in.List); err != nil {
//			return nil, errors.Transform()
//		}
//		return nil, s.uc.Import(kratosx.MustContext(ctx), users)
//
// }

func (s *UserService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*empty.Empty, error) {
	var user biz.User
	if err := copier.Copy(&user, in); err != nil {
		return nil, errors.Transform()
	}
	return nil, s.uc.Update(kratosx.MustContext(ctx), &user)
}

func (s *UserService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*empty.Empty, error) {
	return nil, s.uc.Delete(kratosx.MustContext(ctx), in.Id)
}

func (s *UserService) DisableUser(ctx context.Context, in *pb.DisableUserRequest) (*empty.Empty, error) {
	return nil, s.uc.Disable(kratosx.MustContext(ctx), in.Id, in.Desc)
}

func (s *UserService) EnableUser(ctx context.Context, in *pb.EnableUserRequest) (*empty.Empty, error) {
	return nil, s.uc.Enable(kratosx.MustContext(ctx), in.Id)
}

func (s *UserService) OfflineUser(ctx context.Context, in *pb.OfflineUserRequest) (*empty.Empty, error) {
	return nil, s.uc.Offline(kratosx.MustContext(ctx), in.Id)
}

func (s *UserService) AddUserApp(ctx context.Context, in *pb.AddUserAppRequest) (*empty.Empty, error) {
	_, err := s.uc.AddApp(kratosx.MustContext(ctx), in.UserId, in.AppId)
	return nil, err
}

func (s *UserService) DeleteUserApp(ctx context.Context, in *pb.DeleteUserAppRequest) (*empty.Empty, error) {
	return nil, s.uc.DeleteApp(kratosx.MustContext(ctx), in.UserId, in.AppId)
}

func (s *UserService) OAuthLogin(ctx context.Context, in *pb.OAuthLoginRequest) (*pb.LoginReply, error) {
	var req biz.OAuthLoginRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	token, err := s.uc.OAuthLogin(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		Token: token,
	}, nil
}

func (s *UserService) OAuthBindByPassword(ctx context.Context, in *pb.OAuthBindByPasswordRequest) (*pb.BindReply, error) {
	var req biz.OAuthBindByPasswordRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	token, err := s.uc.OAuthBindByPassword(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	return &pb.BindReply{
		Token: token,
	}, nil
}

func (s *UserService) OAuthBindByCaptcha(ctx context.Context, in *pb.OAuthBindByCaptchaRequest) (*pb.BindReply, error) {
	var req biz.OAuthBindByCaptchaRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	token, err := s.uc.OAuthBindByCaptcha(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	return &pb.BindReply{
		Token: token,
	}, nil
}

func (s *UserService) OAuthBindImageCaptcha(ctx context.Context, _ *empty.Empty) (*pb.CaptchaReply, error) {
	res, err := s.uc.OAuthBindCaptcha(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := pb.CaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *UserService) OAuthBindEmail(ctx context.Context, in *pb.OAuthBindEmailCaptchaRequest) (*pb.CaptchaReply, error) {
	res, err := s.uc.OAuthBindEmail(kratosx.MustContext(ctx), in.Email)
	if err != nil {
		return nil, err
	}

	reply := pb.CaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *UserService) PasswordLogin(ctx context.Context, in *pb.PasswordLoginRequest) (*pb.LoginReply, error) {
	var req biz.PasswordLoginRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	token, err := s.uc.PasswordLogin(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		Token: token,
	}, nil
}

func (s *UserService) PasswordLoginCaptcha(ctx context.Context, _ *empty.Empty) (*pb.CaptchaReply, error) {
	res, err := s.uc.PasswordLoginCaptcha(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := pb.CaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *UserService) PasswordRegister(ctx context.Context, in *pb.PasswordRegisterRequest) (*pb.RegisterReply, error) {
	var req biz.PasswordRegisterRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	token, err := s.uc.PasswordRegister(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterReply{
		Token: token,
	}, nil
}

func (s *UserService) PasswordRegisterCaptcha(ctx context.Context, _ *empty.Empty) (*pb.CaptchaReply, error) {
	res, err := s.uc.PasswordRegisterCaptcha(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := pb.CaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *UserService) PasswordRegisterCheck(ctx context.Context, in *pb.PasswordRegisterCheckRequest) (*pb.PasswordRegisterCheckReply, error) {
	return &pb.PasswordRegisterCheckReply{
		Allow: s.uc.PasswordRegisterCheck(kratosx.MustContext(ctx), in.Username),
	}, nil
}

func (s *UserService) CaptchaLogin(ctx context.Context, in *pb.CaptchaLoginRequest) (*pb.LoginReply, error) {
	var req biz.CaptchaLoginRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	token, err := s.uc.CaptchaLogin(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		Token: token,
	}, nil
}

func (s *UserService) CaptchaLoginEmail(ctx context.Context, in *pb.CaptchaLoginEmailRequest) (*pb.CaptchaReply, error) {
	res, err := s.uc.CaptchaLoginEmail(kratosx.MustContext(ctx), in.Email)
	if err != nil {
		return nil, err
	}

	reply := pb.CaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *UserService) CaptchaRegisterEmail(ctx context.Context, in *pb.CaptchaRegisterEmailRequest) (*pb.CaptchaReply, error) {
	res, err := s.uc.CaptchaRegisterEmail(kratosx.MustContext(ctx), in.Email)
	if err != nil {
		return nil, err
	}

	reply := pb.CaptchaReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *UserService) CaptchaRegister(ctx context.Context, in *pb.CaptchaRegisterRequest) (*pb.RegisterReply, error) {
	var req biz.CaptchaRegisterRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	token, err := s.uc.CaptchaRegister(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterReply{
		Token: token,
	}, nil
}

func (s *UserService) ParseToken(ctx context.Context, _ *empty.Empty) (*pb.ParseTokenReply, error) {
	res, err := s.uc.ParseToken(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := pb.ParseTokenReply{}
	if err := copier.Copy(&reply, res); err != nil {
		return nil, errors.Transform()
	}
	return &reply, nil
}

func (s *UserService) RefreshToken(ctx context.Context, _ *empty.Empty) (*pb.LoginReply, error) {
	token, err := s.uc.RefreshToken(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Token: token}, nil
}
