package data

import (
	"errors"
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"

	biz "github.com/limes-cloud/usercenter/internal/biz/userinfo"
	"github.com/limes-cloud/usercenter/internal/data/model"
)

type userinfoRepo struct {
}

func NewUserinfoRepo() biz.Repo {
	return &userinfoRepo{}
}

// ToUserinfoEntity model转entity
func (r userinfoRepo) ToUserinfoEntity(m *model.Userinfo) *biz.Userinfo {
	e := &biz.Userinfo{}
	if m.Field != nil {
		e.Name = m.Field.Name
	}
	_ = valx.Transform(m, e)
	return e
}

// ToUserinfoModel entity转model
func (r userinfoRepo) ToUserinfoModel(e *biz.Userinfo) *model.Userinfo {
	m := &model.Userinfo{}
	_ = valx.Transform(e, m)
	return m
}

// ListUserinfo 获取列表
func (r userinfoRepo) ListUserinfo(ctx kratosx.Context, req *biz.ListUserinfoRequest) ([]*biz.Userinfo, uint32, error) {
	var (
		bs    []*biz.Userinfo
		ms    []*model.Userinfo
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Userinfo{}).Select(fs).Preload("Field", "status=true")

	if req.UserId != nil {
		db = db.Where("user_id = ?", *req.UserId)
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
		bs = append(bs, r.ToUserinfoEntity(m))
	}
	return bs, uint32(total), nil
}

// GetUserinfo 获取指定的数据
func (r userinfoRepo) GetUserinfo(ctx kratosx.Context, id uint32) (*biz.Userinfo, error) {
	var (
		m  = model.Userinfo{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToUserinfoEntity(&m), nil
}

// UpdateUserinfo 更新数据
func (r userinfoRepo) UpdateUserinfo(ctx kratosx.Context, req *biz.Userinfo) error {
	return ctx.DB().Updates(r.ToUserinfoModel(req)).Error
}

func (r userinfoRepo) CheckKeywords(ctx kratosx.Context, appId uint32, keywords []string) error {
	var (
		keys []string
		ids  []uint32
	)
	if err := ctx.DB().Model(model.AppField{}).
		Select("field_id").
		Where("app_id=?", appId).
		Scan(&ids).Error; err != nil {
		return err
	}

	if err := ctx.DB().Model(model.Field{}).
		Select("keyword").
		Where("id in ?", ids).
		Scan(&keys).Error; err != nil {
		return err
	}

	for _, key := range keywords {
		if !valx.InList(keys, key) {
			return errors.New("not exist key:" + key)
		}
	}
	return nil
}

// CreateUserinfo 创建数据
func (r userinfoRepo) CreateUserinfo(ctx kratosx.Context, req *biz.Userinfo) (uint32, error) {
	m := r.ToUserinfoModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// DeleteUserinfo 删除数据
func (r userinfoRepo) DeleteUserinfo(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.Userinfo{})
	return uint32(db.RowsAffected), db.Error
}
