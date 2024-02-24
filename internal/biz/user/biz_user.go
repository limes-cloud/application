package user

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	v1 "github.com/limes-cloud/resource/api/v1"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/user-center/api/errors"
)

// GetBase 获取用户基础信息
func (u *UseCase) GetBase(ctx kratosx.Context, id uint32) (*User, error) {
	user, err := u.repo.GetBase(ctx, id)
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}
	return user, nil
}

// Get 获取用户信息
func (u *UseCase) Get(ctx kratosx.Context, id uint32) (*User, error) {
	user, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}
	return user, nil
}

// GetByEmail 获取用户信息
func (u *UseCase) GetByEmail(ctx kratosx.Context, email string) (*User, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}
	return user, nil
}

// GetByPhone 获取用户信息
func (u *UseCase) GetByPhone(ctx kratosx.Context, phone string) (*User, error) {
	user, err := u.repo.GetByPhone(ctx, phone)
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}
	return user, nil
}

// GetByUsername 获取用户信息
func (u *UseCase) GetByUsername(ctx kratosx.Context, un string) (*User, error) {
	user, err := u.repo.GetByUsername(ctx, un)
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}
	return user, nil
}

// Page 分页获取用户
func (u *UseCase) Page(ctx kratosx.Context, req *PageUserRequest) ([]*User, uint32, error) {
	list, total, err := u.repo.PageUser(ctx, req)
	if err != nil {
		return nil, 0, errors.DatabaseFormat(err.Error())
	}
	return list, total, err
}

// Add 添加用户信息
func (u *UseCase) Add(ctx kratosx.Context, user *User) (uint32, error) {
	user.Status = proto.Bool(true)
	if user.NickName == "" {
		user.NickName = user.RealName
	}

	id, err := u.repo.Add(ctx, user)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	return id, nil
}

// Import 导入用户信息
func (u *UseCase) Import(ctx kratosx.Context, users []*User) error {
	if err := u.repo.Import(ctx, users); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// Update 添加用户信息
func (u *UseCase) Update(ctx kratosx.Context, user *User) error {
	if err := u.repo.Update(ctx, user); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// Delete 删除用户信息
func (u *UseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// Enable 启用用户
func (u *UseCase) Enable(ctx kratosx.Context, id uint32) error {
	user := &User{
		BaseModel:   ktypes.BaseModel{ID: id},
		Status:      proto.Bool(true),
		DisableDesc: proto.String(""),
	}
	if err := u.repo.Update(ctx, user); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// Disable 禁用用户
func (u *UseCase) Disable(ctx kratosx.Context, id uint32, desc string) error {
	user := &User{
		BaseModel:   ktypes.BaseModel{ID: id},
		Status:      proto.Bool(false),
		DisableDesc: proto.String(desc),
	}
	if err := u.repo.Update(ctx, user); err != nil {
		return v1.DatabaseError()
	}
	return nil
}

// Offline 对当前登陆用户进行下线
func (u *UseCase) Offline(ctx kratosx.Context, id uint32) error {
	tokens := u.repo.GetJwtTokens(ctx, id)
	for _, token := range tokens {
		ctx.JWT().AddBlacklist(token)
	}
	return nil
}

// AddApp 添加用户应用
func (u *UseCase) AddApp(ctx kratosx.Context, uid, aid uint32) (uint32, error) {
	id, err := u.repo.AddUserApp(ctx, uid, aid)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// DeleteApp 删除用户应用
func (u *UseCase) DeleteApp(ctx kratosx.Context, uid, aid uint32) error {
	if err := u.repo.DeleteUserApp(ctx, uid, aid); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
