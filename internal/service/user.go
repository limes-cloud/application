package service

import (
	"context"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	resourceV1 "github.com/limes-cloud/resource/api/v1"
	"google.golang.org/protobuf/types/known/structpb"

	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/internal/biz/types"
	"github.com/limes-cloud/user-center/pkg/field"
	"github.com/limes-cloud/user-center/pkg/md"
	"github.com/limes-cloud/user-center/pkg/service"
	"github.com/limes-cloud/user-center/pkg/util"
)

func (s *Service) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.User, error) {
	var err error
	var user *biz.User
	switch in.Condition.(type) {
	case *v1.GetUserRequest_Id:
		cond := in.Condition.(*v1.GetUserRequest_Id)
		user, err = s.user.Get(kratosx.MustContext(ctx), cond.Id)
	case *v1.GetUserRequest_Email:
		cond := in.Condition.(*v1.GetUserRequest_Email)
		user, err = s.user.GetByEmail(kratosx.MustContext(ctx), cond.Email)
	case *v1.GetUserRequest_Phone:
		cond := in.Condition.(*v1.GetUserRequest_Phone)
		user, err = s.user.GetByPhone(kratosx.MustContext(ctx), cond.Phone)
	case *v1.GetUserRequest_Username:
		cond := in.Condition.(*v1.GetUserRequest_Username)
		user, err = s.user.GetByPhone(kratosx.MustContext(ctx), cond.Username)
	default:
		user, err = nil, v1.NotFoundError()
	}
	if err != nil {
		return nil, err
	}

	return s.transformUserReply(ctx, user, nil)
}

func (s *Service) transformUserReply(ctx context.Context, user *biz.User, app *biz.App) (*v1.User, error) {
	reply := v1.User{}
	if err := copier.Copy(&reply, user); err != nil {
		return nil, v1.TransformError()
	}

	reply.Apps = []*v1.User_App{}
	reply.Channels = []*v1.User_Channel{}
	reply.Extra = make(map[string]*structpb.Value)
	reply.ExtraList = []*v1.User_Extra{}

	// 请求资源中心,错了直接忽略，不影响主流程
	resource, rer := service.NewResource(ctx, s.conf.Service.Resource)

	if reply.Avatar != "" {
		reply.Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{
			Sha: reply.Avatar,
		})
	}

	// 组装数据apps
	uaSet := map[uint32]*biz.UserApp{}
	for _, ua := range user.UserApps {
		uaSet[ua.AppID] = ua
	}
	for _, item := range user.Apps {
		app := &v1.User_App{
			Id:         item.ID,
			Name:       item.Name,
			Logo:       item.Logo,
			RegistryAt: uint32(uaSet[item.ID].CreatedAt),
			LoginAt:    uint32(uaSet[item.ID].LoginAt),
		}
		if rer == nil {
			app.Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.Logo})
		}
		reply.Apps = append(reply.Apps, app)
	}

	// 组装channels
	ucSet := map[uint32]*biz.UserChannel{}
	for _, uc := range user.UserChannels {
		ucSet[uc.ChannelID] = uc
	}
	for _, item := range user.Channels {
		channel := &v1.User_Channel{
			Id:      item.ID,
			Name:    item.Name,
			Logo:    item.Logo,
			AuthAt:  uint32(ucSet[item.ID].CreatedAt),
			LoginAt: uint32(ucSet[item.ID].LoginAt),
		}
		if rer == nil {
			channel.Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.Logo})
		}
		reply.Channels = append(reply.Channels, channel)
	}

	// 转换extra
	efs := s.extraField.FiledTypeSet(kratosx.MustContext(ctx))
	fir := field.New()
	// 通过app过滤字段
	var userFields []string
	hasFilter := false
	if app != nil {
		userFields = strings.Split(app.UserFields, ",")
		hasFilter = true
	}

	for _, item := range user.UserExtras {
		if !hasFilter || !util.InList(userFields, item.Keyword) {
			continue
		}
		if efs[item.Keyword] != nil {
			tp := fir.GetType(efs[item.Keyword].Type)
			reply.Extra[item.Keyword] = tp.ToValue(item.Value)
			reply.ExtraList = append(reply.ExtraList, &v1.User_Extra{
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

func (s *Service) GetSimpleUser(ctx context.Context, in *v1.GetSimpleUserRequest) (*v1.SimpleUser, error) {
	user, err := s.user.GetBase(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	reply := v1.SimpleUser{}
	if err := copier.Copy(&reply, user); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) GetBaseUser(ctx context.Context, in *v1.GetBaseUserRequest) (*v1.BaseUser, error) {
	user, err := s.user.GetBase(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	reply := v1.BaseUser{}
	if err := copier.Copy(&reply, user); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) GetCurrentUser(ctx context.Context, _ *empty.Empty) (*v1.User, error) {
	kCtx := kratosx.MustContext(ctx)
	user, err := s.user.Get(kCtx, md.UserID(kratosx.MustContext(ctx)))
	if err != nil {
		return nil, err
	}

	// 获取当前的app
	app, err := s.app.GetByID(kCtx, md.AppID(kCtx))
	if err != nil {
		return nil, err
	}

	return s.transformUserReply(ctx, user, app)
}

func (s *Service) UpdateCurrentUser(ctx context.Context, in *v1.UpdateCurrentUserRequest) (*empty.Empty, error) {
	user := biz.User{}
	if err := copier.Copy(&user, in); err != nil {
		return nil, v1.TransformError()
	}

	kCtx := kratosx.MustContext(ctx)
	user.ID = md.UserID(kCtx)

	// 转换extra
	efs := s.extraField.FiledTypeSet(kCtx)
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

	return nil, s.user.Update(kratosx.MustContext(ctx), &user)
}

func (s *Service) PageUser(ctx context.Context, in *v1.PageUserRequest) (*v1.PageUserReply, error) {
	var req types.PageUserRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	list, total, err := s.user.Page(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageUserReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) AddUser(ctx context.Context, in *v1.AddUserRequest) (*v1.AddUserReply, error) {
	var user biz.User
	if err := copier.Copy(&user, in); err != nil {
		return nil, v1.TransformError()
	}

	id, err := s.user.Add(kratosx.MustContext(ctx), &user)
	if err != nil {
		return nil, err
	}
	return &v1.AddUserReply{Id: id}, nil
}

//	func (s *Service) ImportUser(ctx context.Context, in *v1.ImportUserRequest) (*empty.Empty, error) {
//		var users []*biz.User
//		if err := copier.Copy(&users, in.List); err != nil {
//			return nil, v1.TransformError()
//		}
//		return nil, s.user.Import(kratosx.MustContext(ctx), users)
//
// }
func (s *Service) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*empty.Empty, error) {
	var user biz.User
	if err := copier.Copy(&user, in); err != nil {
		return nil, v1.TransformError()
	}
	return nil, s.user.Update(kratosx.MustContext(ctx), &user)
}

func (s *Service) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*empty.Empty, error) {
	return nil, s.user.Delete(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) DisableUser(ctx context.Context, in *v1.DisableUserRequest) (*empty.Empty, error) {
	return nil, s.user.Disable(kratosx.MustContext(ctx), in.Id, in.Desc)
}

func (s *Service) EnableUser(ctx context.Context, in *v1.EnableUserRequest) (*empty.Empty, error) {
	return nil, s.user.Enable(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) OfflineUser(ctx context.Context, in *v1.OfflineUserRequest) (*empty.Empty, error) {
	return nil, s.user.Offline(kratosx.MustContext(ctx), in.Id)
}
