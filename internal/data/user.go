package data

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"

	biz "github.com/limes-cloud/usercenter/internal/biz/user"
	"github.com/limes-cloud/usercenter/internal/data/model"
	"github.com/limes-cloud/usercenter/internal/pkg/resource"
)

type userRepo struct {
}

func NewUserRepo() biz.Repo {
	return &userRepo{}
}

// ToUserEntity model转entity
func (r userRepo) ToUserEntity(m *model.User) *biz.User {
	e := &biz.User{}
	_ = valx.Transform(m, e)
	return e
}

// ToUserModel entity转model
func (r userRepo) ToUserModel(e *biz.User) *model.User {
	m := &model.User{}
	_ = valx.Transform(e, m)
	return m
}

// GetUserByPhone 获取指定数据
func (r userRepo) GetUserByPhone(ctx kratosx.Context, phone string) (*biz.User, error) {
	var (
		m  = model.User{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("phone = ?", phone).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToUserEntity(&m), nil
}

// GetUserByEmail 获取指定数据
func (r userRepo) GetUserByEmail(ctx kratosx.Context, email string) (*biz.User, error) {
	var (
		m  = model.User{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("email = ?", email).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToUserEntity(&m), nil
}

// GetUserByUsername 获取指定数据
func (r userRepo) GetUserByUsername(ctx kratosx.Context, username string) (*biz.User, error) {
	var (
		m  = model.User{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("username = ?", username).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToUserEntity(&m), nil
}

// GetUser 获取指定的数据
func (r userRepo) GetUser(ctx kratosx.Context, id uint32) (*biz.User, error) {
	var (
		m  = model.User{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}
	b := r.ToUserEntity(&m)
	if b.Avatar != nil && *b.Avatar != "" {
		b.AvatarUrl = proto.String(resource.GetURLBySha(ctx, *b.Avatar))
	}
	return b, nil
}

// ListUser 获取列表
func (r userRepo) ListUser(ctx kratosx.Context, req *biz.ListUserRequest) ([]*biz.User, uint32, error) {
	var (
		bs    []*biz.User
		ms    []*model.User
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.User{}).Select(fs)

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
	if req.App != nil {
		appId := 0
		if err := ctx.DB().Model(model.App{}).
			Select("id").Where("keyword=?", *req.App).
			Scan(&appId).Error; err != nil {
			return nil, 0, err
		}
		db = db.InnerJoins("Auths", ctx.DB().Where("Auths.app_id=?", *req.AppId))
	}
	if req.InIds != nil {
		db = db.Where("id in ?", req.InIds)
	}
	if req.NotInIds != nil {
		db = db.Where("id not in ?", req.NotInIds)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

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

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		b := r.ToUserEntity(m)
		if b.Avatar != nil && *b.Avatar != "" {
			b.AvatarUrl = proto.String(resource.GetURLBySha(ctx, *b.Avatar))
		}
		bs = append(bs, r.ToUserEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateUser 创建数据
func (r userRepo) CreateUser(ctx kratosx.Context, req *biz.User) (uint32, error) {
	m := r.ToUserModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// ImportUser 导入数据
func (r userRepo) ImportUser(ctx kratosx.Context, req []*biz.User) (uint32, error) {
	var (
		ms []*model.User
	)

	for _, item := range req {
		ms = append(ms, r.ToUserModel(item))
	}

	db := ctx.DB().Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(ms, 1000)
	return uint32(len(req)), db.Error
}

// ExportUser 导出数据
func (r userRepo) ExportUser(ctx kratosx.Context, req *biz.ExportUserRequest) (string, error) {
	return "", nil
}

// UpdateUser 更新数据
func (r userRepo) UpdateUser(ctx kratosx.Context, req *biz.User) error {
	return ctx.DB().Updates(r.ToUserModel(req)).Error
}

// UpdateUserStatus 更新数据状态
func (r userRepo) UpdateUserStatus(ctx kratosx.Context, req *biz.UpdateUserStatusRequest) error {
	return ctx.DB().Model(model.User{}).
		Where("id=?", req.Id).
		Updates(map[string]any{
			"status":       req.Status,
			"disable_desc": req.DisableDesc,
		}).Error
}

// DeleteUser 删除数据
func (r userRepo) DeleteUser(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.User{})
	return uint32(db.RowsAffected), db.Error
}

// GetTrashUser 获取垃圾桶指定数据
func (r userRepo) GetTrashUser(ctx kratosx.Context, id uint32) (*biz.User, error) {
	var (
		m  = model.User{}
		fs = []string{"id", "phone", "email", "username", "password", "nick_name", "real_name", "avatar", "gender", "status", "disable_desc", "from", "from_desc", "created_at", "updated_at", "deleted_at"}
	)

	if err := ctx.DB().Unscoped().Select(fs).First(&m, "id=? and deleted_at != 0", id).Error; err != nil {
		return nil, err
	}

	return r.ToUserEntity(&m), nil
}

// ListTrashUser 获取垃圾桶列表
func (r userRepo) ListTrashUser(ctx kratosx.Context, req *biz.ListTrashUserRequest) ([]*biz.User, uint32, error) {
	var (
		bs    []*biz.User
		ms    []*model.User
		total int64
		fs    = []string{"id", "phone", "email", "username", "password", "nick_name", "real_name", "avatar", "gender", "status", "disable_desc", "from", "from_desc", "created_at", "updated_at", "deleted_at"}
	)

	db := ctx.DB().Unscoped().Model(model.User{}).Select(fs)
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
	} else {
		*req.OrderBy = *req.OrderBy + ",id"
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
	if err := db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: *req.OrderBy},
		Desc:   *req.Order == "desc",
	}).Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		b := r.ToUserEntity(m)
		if b.Avatar != nil && *b.Avatar != "" {
			b.AvatarUrl = proto.String(resource.GetURLBySha(ctx, *b.Avatar))
		}
		bs = append(bs, b)
	}
	return bs, uint32(total), nil
}

// DeleteTrashUser 彻底删除数据
func (r userRepo) DeleteTrashUser(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Unscoped().Delete(model.User{}, "id in ?", ids)
	return uint32(db.RowsAffected), db.Error
}

// RevertTrashUser 还原指定的数据
func (r userRepo) RevertTrashUser(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Unscoped().Model(model.User{}).Where("id=?", id).Update("deleted_at", 0).Error
}
