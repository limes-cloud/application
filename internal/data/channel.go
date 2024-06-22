package data

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	biz "github.com/limes-cloud/usercenter/internal/biz/channel"
	"github.com/limes-cloud/usercenter/internal/data/model"
	"github.com/limes-cloud/usercenter/internal/pkg/resource"
)

type channelRepo struct {
}

func NewChannelRepo() biz.Repo {
	return &channelRepo{}
}

// ToChannelEntity model转entity
func (r channelRepo) ToChannelEntity(m *model.Channel) *biz.Channel {
	e := &biz.Channel{
		Id:        m.Id,
		Logo:      m.Logo,
		Keyword:   m.Keyword,
		Name:      m.Name,
		Status:    m.Status,
		Ak:        m.Ak,
		Sk:        m.Sk,
		Extra:     m.Extra,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
	return e
}

// ToChannelModel entity转model
func (r channelRepo) ToChannelModel(e *biz.Channel) *model.Channel {
	m := &model.Channel{
		BaseModel: types.BaseModel{
			Id:        e.Id,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
		Logo:    e.Logo,
		Keyword: e.Keyword,
		Name:    e.Name,
		Status:  e.Status,
		Ak:      e.Ak,
		Sk:      e.Sk,
		Extra:   e.Extra,
	}
	_ = valx.Transform(e, m)
	return m
}

// GetChannelByKeyword 获取指定数据
func (r channelRepo) GetChannelByKeyword(ctx kratosx.Context, keyword string) (*biz.Channel, error) {
	var (
		m  = model.Channel{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	b := r.ToChannelEntity(&m)
	if b.Logo != "" {
		b.LogoUrl = resource.GetURLBySha(ctx, b.Logo)
	}
	return b, nil
}

// GetChannel 获取指定的数据
func (r channelRepo) GetChannel(ctx kratosx.Context, id uint32) (*biz.Channel, error) {
	var (
		m  = model.Channel{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	b := r.ToChannelEntity(&m)
	if b.Logo != "" {
		b.LogoUrl = resource.GetURLBySha(ctx, b.Logo)
	}
	return b, nil
}

// ListChannel 获取列表
func (r channelRepo) ListChannel(ctx kratosx.Context, req *biz.ListChannelRequest) ([]*biz.Channel, uint32, error) {
	var (
		bs    []*biz.Channel
		ms    []*model.Channel
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Channel{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
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
		b := r.ToChannelEntity(m)
		if b.Logo != "" {
			b.LogoUrl = resource.GetURLBySha(ctx, b.Logo)
		}
		bs = append(bs, b)
	}
	return bs, uint32(total), nil
}

// CreateChannel 创建数据
func (r channelRepo) CreateChannel(ctx kratosx.Context, req *biz.Channel) (uint32, error) {
	m := r.ToChannelModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// UpdateChannel 更新数据
func (r channelRepo) UpdateChannel(ctx kratosx.Context, req *biz.Channel) error {
	return ctx.DB().Updates(r.ToChannelModel(req)).Error
}

// UpdateChannelStatus 更新数据状态
func (r channelRepo) UpdateChannelStatus(ctx kratosx.Context, id uint32, status bool) error {
	return ctx.DB().Model(model.Channel{}).Where("id=?", id).Update("status", status).Error
}

// DeleteChannel 删除数据
func (r channelRepo) DeleteChannel(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.Channel{})
	return uint32(db.RowsAffected), db.Error
}
