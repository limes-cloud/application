package dbs

import (
	"fmt"
	"sync"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/library/db/gormtranserror"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"

	"github.com/limes-cloud/application/internal/domain/entity"
	"github.com/limes-cloud/application/internal/types"
)

type User struct {
}

var (
	userIns  *User
	userOnce sync.Once
)

func NewUser() *User {
	userOnce.Do(func() {
		userIns = &User{}
	})
	return userIns
}

// HasUserByEmail 是否存在指定的用户邮箱
func (r User) HasUserByEmail(ctx kratosx.Context, email string) bool {
	var count int64 = 0
	_ = ctx.DB().Model(entity.User{}).Where("email=?", email).Count(&count)
	return count > 0
}

// HasUserByUsername 是否存在指定的用户账户
func (r User) HasUserByUsername(ctx kratosx.Context, username string) bool {
	var count int64 = 0
	_ = ctx.DB().Model(entity.User{}).Where("username=?", username).Count(&count)
	return count > 0
}

// GetUserByPhone 获取指定数据
func (r User) GetUserByPhone(ctx kratosx.Context, phone string) (*entity.User, error) {
	var (
		user = entity.User{}
		fs   = []string{"*"}
	)

	return &user, ctx.DB().Select(fs).Where("phone = ?", phone).First(&user).Error
}

// GetUserByEmail 获取指定数据
func (r User) GetUserByEmail(ctx kratosx.Context, email string) (*entity.User, error) {
	var (
		user = entity.User{}
		fs   = []string{"*"}
	)
	return &user, ctx.DB().Select(fs).Where("email = ?", email).First(&user).Error
}

// GetUserByUsername 获取指定数据
func (r User) GetUserByUsername(ctx kratosx.Context, username string) (*entity.User, error) {
	var (
		user = entity.User{}
		fs   = []string{"*"}
	)
	return &user, ctx.DB().Select(fs).Where("username = ?", username).First(&user).Error
}

// GetUser 获取指定的数据
func (r User) GetUser(ctx kratosx.Context, id uint32) (*entity.User, error) {
	var (
		user = entity.User{}
		fs   = []string{"*"}
	)
	return &user, ctx.DB().Select(fs).First(&user, id).Error
}

// ListUser 获取列表
func (r User) ListUser(ctx kratosx.Context, req *types.ListUserRequest) ([]*entity.User, uint32, error) {
	var (
		list  []*entity.User
		total int64
		fs    = []string{"user.id", "user.username", "user.nick_name", "user.real_name", "user.phone", "user.email", "user.gender", "user.from",
			"user.from_desc", "user.avatar", "user.status", "user.disable_desc", "user.created_at", "user.updated_at"}
	)

	db := ctx.DB().Model(entity.User{}).Select(fs)

	if req.Phone != nil {
		db = db.Where("phone = ?", *req.Phone)
	}
	if req.Email != nil {
		db = db.Where("email = ?", *req.Email)
	}
	if req.Username != nil {
		db = db.Where("username = ?", *req.Username)
	}
	if req.RealName != nil {
		db = db.Where("real_name LIKE ?", *req.RealName+"%")
	}
	if req.Gender != nil {
		db = db.Where("gender = ?", *req.Gender)
	}
	if req.Status != nil {
		db = db.Where("user.status = ?", *req.Status)
	}
	if req.From != nil {
		db = db.Where("from = ?", *req.From)
	}
	if len(req.CreatedAts) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", req.CreatedAts[0], req.CreatedAts[1])
	}
	if req.AppId != nil {
		db = db.InnerJoins("Auths", ctx.DB().Where("Auths.app_id=?", *req.AppId))
	}
	if req.AppId == nil && req.AppIds != nil {
		db = db.InnerJoins("Auths", ctx.DB().Where("Auths.app_id in ?", req.AppIds))
	}

	if req.InIds != nil {
		db = db.Where("user.id in ?", req.InIds)
	}
	if req.NotInIds != nil {
		db = db.Where("user.id not in ?", req.NotInIds)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("user.id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "user.id" {
		db = db.Order("id asc")
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateUser 创建数据
func (r User) CreateUser(ctx kratosx.Context, user *entity.User) (uint32, error) {
	return user.Id, ctx.DB().Create(user).Error
}

// ImportUser 导入数据
func (r User) ImportUser(ctx kratosx.Context, users []*entity.User) (uint32, error) {
	for _, user := range users {
		var oldUser entity.User
		if err := ctx.DB().Where("phone=?", user.Phone).First(&oldUser).Error; err != nil {
			if !gormtranserror.Is(err, gorm.ErrRecordNotFound) {
				return 0, err
			}
			if err := ctx.DB().Create(user).Error; err != nil {
				return 0, err
			}
		} else {
			user.Id = oldUser.Id
			if err := ctx.DB().Updates(user).Error; err != nil {
				return 0, err
			}
		}
	}
	return uint32(len(users)), nil
}

// ExportUser 导出数据
func (r User) ExportUser(ctx kratosx.Context, req *types.ExportUserRequest) (string, error) {
	return "", nil
}

// UpdateUser 更新数据
func (r User) UpdateUser(ctx kratosx.Context, user *entity.User) error {
	return ctx.DB().Updates(user).Error
}

// UpdateUserStatus 更新数据状态
func (r User) UpdateUserStatus(ctx kratosx.Context, req *types.UpdateUserStatusRequest) error {
	return ctx.DB().Model(entity.User{}).
		Where("id=?", req.Id).
		Updates(map[string]any{
			"status":       req.Status,
			"disable_desc": req.DisableDesc,
		}).Error
}

// DeleteUser 删除数据
func (r User) DeleteUser(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&entity.User{})
	return uint32(db.RowsAffected), db.Error
}

// GetTrashUser 获取垃圾桶指定数据
func (r User) GetTrashUser(ctx kratosx.Context, id uint32) (*entity.User, error) {
	var (
		user = entity.User{}
		fs   = []string{"*"}
	)

	if err := ctx.DB().Unscoped().Select(fs).First(&user, "id=? and deleted_at != 0", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// ListTrashUser 获取垃圾桶列表
func (r User) ListTrashUser(ctx kratosx.Context, req *types.ListTrashUserRequest) ([]*entity.User, uint32, error) {
	var (
		list  []*entity.User
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Unscoped().Model(entity.User{}).Select(fs)
	db = db.Where("deleted_at != 0")
	if req.Phone != nil {
		db = db.Where("phone = ?", *req.Phone)
	}
	if req.Email != nil {
		db = db.Where("email = ?", *req.Email)
	}
	if req.Username != nil {
		db = db.Where("username = ?", *req.Username)
	}
	if req.RealName != nil {
		db = db.Where("real_name LIKE ?", *req.RealName+"%")
	}
	if req.Gender != nil {
		db = db.Where("gender = ?", *req.Gender)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.From != nil {
		db = db.Where("from = ?", *req.From)
	}
	if len(req.CreatedAts) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", req.CreatedAts[0], req.CreatedAts[1])
	}
	if req.AppId != nil {
		db = db.InnerJoins("Auths", ctx.DB().Where("Auths.app_id=?", *req.AppId))
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
	return list, uint32(total), db.Find(&list).Error
}

// DeleteTrashUser 彻底删除数据
func (r User) DeleteTrashUser(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Unscoped().Delete(entity.User{}, "id in ?", ids)
	return uint32(db.RowsAffected), db.Error
}

// RevertTrashUser 还原指定的数据
func (r User) RevertTrashUser(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Unscoped().Model(entity.User{}).Where("id=?", id).Update("deleted_at", 0).Error
}
