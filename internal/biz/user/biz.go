package user

import (
	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/pkg/md"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// GetCurrentUser 获取当前用户信息
func (u *UseCase) GetCurrentUser(ctx kratosx.Context) (*User, error) {
	user, err := u.repo.GetUser(ctx, md.UserID(ctx))

	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return user, nil
}

// UpdateCurrentUser 获取当前用户信息
func (u *UseCase) UpdateCurrentUser(ctx kratosx.Context, req *UpdateCurrentUserRequest) error {
	if err := u.repo.UpdateUser(ctx, &User{
		Id:       md.UserID(ctx),
		NickName: req.NickName,
		Avatar:   &req.Avatar,
		Gender:   &req.Gender,
	}); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// GetUser 获取指定的用户信息
func (u *UseCase) GetUser(ctx kratosx.Context, req *GetUserRequest) (*User, error) {
	var (
		res *User
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
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// ListUser 获取用户信息列表
func (u *UseCase) ListUser(ctx kratosx.Context, req *ListUserRequest) ([]*User, uint32, error) {
	list, total, err := u.repo.ListUser(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateUser 创建用户信息
func (u *UseCase) CreateUser(ctx kratosx.Context, req *User) (uint32, error) {
	req.Status = proto.Bool(true)
	md5 := crypto.MD5([]byte(uuid.NewString()))
	req.NickName = u.conf.DefaultNickName + md5[:8]
	id, err := u.repo.CreateUser(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// ImportUser 导入用户信息
func (u *UseCase) ImportUser(ctx kratosx.Context, req []*User) (uint32, error) {
	total, err := u.repo.ImportUser(ctx, req)
	if err != nil {
		return 0, errors.ImportError(err.Error())
	}
	return total, nil
}

// ExportUser 导出用户信息
func (u *UseCase) ExportUser(ctx kratosx.Context, req *ExportUserRequest) (string, error) {
	id, err := u.repo.ExportUser(ctx, req)
	if err != nil {
		return "", errors.ExportError(err.Error())
	}
	return id, nil
}

// UpdateUser 更新用户信息
func (u *UseCase) UpdateUser(ctx kratosx.Context, req *User) error {
	if err := u.repo.UpdateUser(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateUserStatus 更新用户信息状态
func (u *UseCase) UpdateUserStatus(ctx kratosx.Context, req *UpdateUserStatusRequest) error {
	if err := u.repo.UpdateUserStatus(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteUser 删除用户信息
func (u *UseCase) DeleteUser(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteUser(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// GetTrashUser 获取指定的用户信息
func (u *UseCase) GetTrashUser(ctx kratosx.Context, id uint32) (*User, error) {
	user, err := u.repo.GetTrashUser(ctx, id)
	if err != nil {
		return nil, errors.GetTrashError(err.Error())
	}
	return user, nil
}

// ListTrashUser 获取用户信息列表
func (u *UseCase) ListTrashUser(ctx kratosx.Context, req *ListTrashUserRequest) ([]*User, uint32, error) {
	list, total, err := u.repo.ListTrashUser(ctx, req)
	if err != nil {
		return nil, 0, errors.ListTrashError(err.Error())
	}
	return list, total, nil
}

// DeleteTrashUser 彻底删除用户信息
func (u *UseCase) DeleteTrashUser(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteTrashUser(ctx, ids)
	if err != nil {
		return 0, errors.DeleteTrashError(err.Error())
	}
	return total, nil
}

// RevertTrashUser 还原删除用户信息
func (u *UseCase) RevertTrashUser(ctx kratosx.Context, id uint32) error {
	if err := u.repo.RevertTrashUser(ctx, id); err != nil {
		return errors.RevertTrashError(err.Error())
	}
	return nil
}
