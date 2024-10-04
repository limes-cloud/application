package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	pb "github.com/limes-cloud/application/api/application/user/v1"
	"github.com/limes-cloud/application/internal/conf"
	"github.com/limes-cloud/application/internal/domain/entity"
	"github.com/limes-cloud/application/internal/domain/service"
	"github.com/limes-cloud/application/internal/infra/dbs"
	"github.com/limes-cloud/application/internal/infra/rpc"
	"github.com/limes-cloud/application/internal/types"
)

type User struct {
	pb.UnimplementedUserServer
	srv *service.User
}

func NewUser(conf *conf.Config) *User {
	return &User{
		srv: service.NewUser(
			conf,
			dbs.NewUser(),
			dbs.NewApp(),
			rpc.NewPermission(),
			rpc.NewFile(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewUser(c)
		pb.RegisterUserHTTPServer(hs, srv)
		pb.RegisterUserServer(gs, srv)
	})
}

// GetCurrentUser 获取当前用户信息
func (s *User) GetCurrentUser(c context.Context, _ *pb.GetCurrentUserRequest) (*pb.GetCurrentUserReply, error) {
	user, err := s.srv.GetCurrentUser(kratosx.MustContext(c))
	if err != nil {
		return nil, err
	}

	return &pb.GetCurrentUserReply{
		Id:        user.Id,
		Phone:     user.Phone,
		Email:     user.Email,
		Username:  user.Username,
		NickName:  user.NickName,
		RealName:  user.RealName,
		Avatar:    user.Avatar,
		AvatarUrl: user.AvatarUrl,
		Gender:    user.Gender,
	}, nil
}

// UpdateCurrentUser 更新当前用户信息
func (s *User) UpdateCurrentUser(c context.Context, req *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserReply, error) {
	return &pb.UpdateCurrentUserReply{}, s.srv.UpdateCurrentUser(kratosx.MustContext(c), &types.UpdateCurrentUserRequest{
		NickName: req.NickName,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
	})
}

// GetUser 获取指定的用户信息
func (s *User) GetUser(c context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.srv.GetUser(kratosx.MustContext(c), &types.GetUserRequest{
		Id:       req.Id,
		Phone:    req.Phone,
		Email:    req.Email,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetUserReply{
		Id:          user.Id,
		Phone:       user.Phone,
		Email:       user.Email,
		Username:    user.Username,
		NickName:    user.NickName,
		RealName:    user.RealName,
		Avatar:      user.Avatar,
		AvatarUrl:   user.AvatarUrl,
		Gender:      user.Gender,
		Status:      user.Status,
		DisableDesc: user.DisableDesc,
		From:        user.From,
		FromDesc:    user.FromDesc,
		CreatedAt:   uint32(user.CreatedAt),
		UpdatedAt:   uint32(user.UpdatedAt),
	}, nil
}

// ListUser 获取用户信息列表
func (s *User) ListUser(c context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	list, total, err := s.srv.ListUser(kratosx.MustContext(c), &types.ListUserRequest{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Order:      req.Order,
		OrderBy:    req.OrderBy,
		Phone:      req.Phone,
		Email:      req.Email,
		Username:   req.Username,
		RealName:   req.RealName,
		Gender:     req.Gender,
		Status:     req.Status,
		From:       req.From,
		CreatedAts: req.CreatedAts,
		AppId:      req.AppId,
		App:        req.App,
		InIds:      req.InIds,
		NotInIds:   req.NotInIds,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListUserReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListUserReply_User{
			Id:          item.Id,
			Phone:       item.Phone,
			Email:       item.Email,
			Username:    item.Username,
			NickName:    item.NickName,
			RealName:    item.RealName,
			Avatar:      item.Avatar,
			AvatarUrl:   item.AvatarUrl,
			Gender:      item.Gender,
			Status:      item.Status,
			DisableDesc: item.DisableDesc,
			From:        item.From,
			FromDesc:    item.FromDesc,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateUser 创建用户信息
func (s *User) CreateUser(c context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	id, err := s.srv.CreateUser(kratosx.MustContext(c), &entity.User{
		Phone:    req.Phone,
		Email:    req.Email,
		RealName: req.RealName,
		Gender:   req.Gender,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{Id: id}, nil
}

// ImportUser 导入用户信息
func (s *User) ImportUser(c context.Context, req *pb.ImportUserRequest) (*pb.ImportUserReply, error) {
	var list []*entity.User
	for _, item := range req.List {
		user := &entity.User{
			Phone:    item.Phone,
			Email:    item.Email,
			RealName: item.RealName,
			Gender:   item.Gender,
		}
		if item.AppId != nil {
			user.Auths = append(user.Auths, &entity.Auth{
				AppId:  *item.AppId,
				Status: proto.Bool(true),
			})
		}
		list = append(list, user)
	}

	total, err := s.srv.ImportUser(kratosx.MustContext(c), list)
	if err != nil {
		return nil, err
	}
	return &pb.ImportUserReply{Total: total}, nil
}

// ExportUser 导出用户信息
func (s *User) ExportUser(c context.Context, req *pb.ExportUserRequest) (*pb.ExportUserReply, error) {
	src, err := s.srv.ExportUser(kratosx.MustContext(c), &types.ExportUserRequest{
		Phone:      req.Phone,
		Email:      req.Email,
		Username:   req.Username,
		RealName:   req.RealName,
		Gender:     req.Gender,
		Status:     req.Status,
		From:       req.From,
		CreatedAts: req.CreatedAts,
	})
	if err != nil {
		return nil, err
	}

	return &pb.ExportUserReply{Src: src}, nil
}

// UpdateUser 更新用户信息
func (s *User) UpdateUser(c context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	if err := s.srv.UpdateUser(kratosx.MustContext(c), &entity.User{
		DeleteModel: ktypes.DeleteModel{Id: req.Id},
		Phone:       req.Phone,
		Email:       req.Email,
		RealName:    req.RealName,
		Gender:      req.Gender,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateUserReply{}, nil
}

// UpdateUserStatus 更新用户信息状态
func (s *User) UpdateUserStatus(c context.Context, req *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusReply, error) {
	return &pb.UpdateUserStatusReply{}, s.srv.UpdateUserStatus(kratosx.MustContext(c), &types.UpdateUserStatusRequest{
		Id:          req.Id,
		Status:      req.Status,
		DisableDesc: req.DisableDesc,
	})
}

//
// // DeleteUser 删除用户信息
// func (s *User) DeleteUser(c context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
//	total, err := s.srv.DeleteUser(kratosx.MustContext(c), req.Ids)
//	if err != nil {
//		return nil, err
//	}
//	return &pb.DeleteUserReply{Total: total}, nil
// }
//
// // GetTrashUser 获取回收站指定的用户信息
// func (s *User) GetTrashUser(c context.Context, req *pb.GetTrashUserRequest) (*pb.GetTrashUserReply, error) {
//	user, err := s.srv.GetTrashUser(kratosx.MustContext(c), req.Id)
//	if err != nil {
//		return nil, err
//	}
//
//	return &pb.GetTrashUserReply{
//		Id:        user.Id,
//		Phone:     user.Phone,
//		Email:     user.Email,
//		Username:  user.Username,
//		NickName:  user.NickName,
//		RealName:  user.RealName,
//		Avatar:    user.Avatar,
//		AvatarUrl: user.AvatarUrl,
//		Gender:    user.Gender,
//	}, nil
// }
//
// // ListTrashUser 获取回收站用户信息列表
// func (s *User) ListTrashUser(c context.Context, req *pb.ListTrashUserRequest) (*pb.ListTrashUserReply, error) {
//	list, total, err := s.srv.ListTrashUser(kratosx.MustContext(c), &types.ListTrashUserRequest{
//		Page:       req.Page,
//		PageSize:   req.PageSize,
//		Order:      req.Order,
//		OrderBy:    req.OrderBy,
//		Phone:      req.Phone,
//		Email:      req.Email,
//		Username:   req.Username,
//		RealName:   req.RealName,
//		Gender:     req.Gender,
//		Status:     req.Status,
//		From:       req.From,
//		CreatedAts: req.CreatedAts,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	reply := pb.ListTrashUserReply{Total: total}
//	for _, item := range list {
//		reply.List = append(reply.List, &pb.ListTrashUserReply_User{
//			Id:        item.Id,
//			Phone:     item.Phone,
//			Email:     item.Email,
//			Username:  item.Username,
//			NickName:  item.NickName,
//			RealName:  item.RealName,
//			Avatar:    item.Avatar,
//			AvatarUrl: item.AvatarUrl,
//			Gender:    item.Gender,
//		})
//	}
//	return &reply, nil
// }
//
// // DeleteTrashUser 删除回收站用户信息
// func (s *User) DeleteTrashUser(ctx context.Context, req *pb.DeleteTrashUserRequest) (*pb.DeleteTrashUserReply, error) {
//	total, err := s.srv.DeleteTrashUser(kratosx.MustContext(ctx), req.Ids)
//	if err != nil {
//		return nil, err
//	}
//	return &pb.DeleteTrashUserReply{Total: total}, nil
// }
//
// // RevertTrashUser 还原回收站用户信息
// func (s *User) RevertTrashUser(ctx context.Context, req *pb.RevertTrashUserRequest) (*pb.RevertTrashUserReply, error) {
//	return &pb.RevertTrashUserReply{}, s.srv.RevertTrashUser(kratosx.MustContext(ctx), req.Id)
// }
