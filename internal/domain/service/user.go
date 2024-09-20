package service

import (
	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	ktypes "github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/application/api/application/errors"
	"github.com/limes-cloud/application/internal/conf"
	"github.com/limes-cloud/application/internal/domain/entity"
	"github.com/limes-cloud/application/internal/domain/repository"
	"github.com/limes-cloud/application/internal/pkg/md"
	"github.com/limes-cloud/application/internal/types"
)

type User struct {
	conf       *conf.Config
	repo       repository.User
	permission repository.Permission
	file       repository.File
}

func NewUser(
	conf *conf.Config,
	repo repository.User,
	permission repository.Permission,
	file repository.File,
) *User {
	return &User{
		conf:       conf,
		repo:       repo,
		permission: permission,
		file:       file,
	}
}

// GetCurrentUser 获取当前用户信息
func (u *User) GetCurrentUser(ctx kratosx.Context) (*entity.User, error) {
	user, err := u.repo.GetUser(ctx, md.UserID(ctx))
	if err != nil {
		ctx.Logger().Warnw("msg", "get cur user error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	return user, nil
}

// UpdateCurrentUser 获取当前用户信息
func (u *User) UpdateCurrentUser(ctx kratosx.Context, req *types.UpdateCurrentUserRequest) error {
	if err := u.repo.UpdateUser(ctx, &entity.User{
		DeleteModel: ktypes.DeleteModel{Id: md.UserID(ctx)},
		NickName:    req.NickName,
		Avatar:      &req.Avatar,
		Gender:      &req.Gender,
	}); err != nil {
		ctx.Logger().Warnw("msg", "update cur user error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// GetUser 获取指定的用户信息
func (u *User) GetUser(ctx kratosx.Context, req *types.GetUserRequest) (*entity.User, error) {
	var (
		res *entity.User
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetUser(ctx, *req.Id)
	} else if req.Phone != nil {
		res, err = u.repo.GetUserByPhone(ctx, *req.Phone)
	} else if req.Email != nil {
		res, err = u.repo.GetUserByEmail(ctx, *req.Email)
	} else if req.Username != nil {
		res, err = u.repo.GetUserByUsername(ctx, *req.Username)
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		ctx.Logger().Warnw("msg", "get user error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// ListUser 获取用户信息列表
func (u *User) ListUser(ctx kratosx.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error) {
	list, total, err := u.repo.ListUser(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list user error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateUser 创建用户信息
func (u *User) CreateUser(ctx kratosx.Context, req *entity.User) (uint32, error) {
	req.Status = proto.Bool(true)
	md5 := crypto.MD5([]byte(uuid.NewString()))
	req.NickName = u.conf.DefaultNickName + md5[:8]
	id, err := u.repo.CreateUser(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create user error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// ImportUser 导入用户信息
func (u *User) ImportUser(ctx kratosx.Context, req []*entity.User) (uint32, error) {
	total, err := u.repo.ImportUser(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "import user error", "err", err.Error())
		return 0, errors.ImportError(err.Error())
	}
	return total, nil
}

// ExportUser 导出用户信息
func (u *User) ExportUser(ctx kratosx.Context, req *types.ExportUserRequest) (string, error) {
	id, err := u.repo.ExportUser(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "export user error", "err", err.Error())
		return "", errors.ExportError(err.Error())
	}
	return id, nil
}

// UpdateUser 更新用户信息
func (u *User) UpdateUser(ctx kratosx.Context, req *entity.User) error {
	if err := u.repo.UpdateUser(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update user error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateUserStatus 更新用户信息状态
func (u *User) UpdateUserStatus(ctx kratosx.Context, req *types.UpdateUserStatusRequest) error {
	if err := u.repo.UpdateUserStatus(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update user status error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteUser 删除用户信息
func (u *User) DeleteUser(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteUser(ctx, ids)
	if err != nil {
		ctx.Logger().Warnw("msg", "delete user error", "err", err.Error())
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// GetTrashUser 获取指定的用户信息
func (u *User) GetTrashUser(ctx kratosx.Context, id uint32) (*entity.User, error) {
	user, err := u.repo.GetTrashUser(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get trash user error", "err", err.Error())
		return nil, errors.GetTrashError(err.Error())
	}
	return user, nil
}

// ListTrashUser 获取用户信息列表
func (u *User) ListTrashUser(ctx kratosx.Context, req *types.ListTrashUserRequest) ([]*entity.User, uint32, error) {
	list, total, err := u.repo.ListTrashUser(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list trash user error", "err", err.Error())
		return nil, 0, errors.ListTrashError(err.Error())
	}
	return list, total, nil
}

// DeleteTrashUser 彻底删除用户信息
func (u *User) DeleteTrashUser(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteTrashUser(ctx, ids)
	if err != nil {
		ctx.Logger().Warnw("msg", "delete trash user error", "err", err.Error())
		return 0, errors.DeleteTrashError(err.Error())
	}
	return total, nil
}

// RevertTrashUser 还原删除用户信息
func (u *User) RevertTrashUser(ctx kratosx.Context, id uint32) error {
	if err := u.repo.RevertTrashUser(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "revert trash user error", "err", err.Error())
		return errors.RevertTrashError(err.Error())
	}
	return nil
}
