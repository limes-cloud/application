package data

import (
	"fmt"

	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"

	biz "github.com/limes-cloud/usercenter/internal/biz/feedback"
	"github.com/limes-cloud/usercenter/internal/data/model"
	"github.com/limes-cloud/usercenter/internal/pkg/resource"
)

type feedbackRepo struct {
}

func NewFeedbackRepo() biz.Repo {
	return &feedbackRepo{}
}

// ToFeedbackCategoryEntity model转entity
func (r feedbackRepo) ToFeedbackCategoryEntity(m *model.FeedbackCategory) *biz.FeedbackCategory {
	e := &biz.FeedbackCategory{}
	_ = valx.Transform(m, e)
	return e
}

// ToFeedbackCategoryModel entity转model
func (r feedbackRepo) ToFeedbackCategoryModel(e *biz.FeedbackCategory) *model.FeedbackCategory {
	m := &model.FeedbackCategory{}
	_ = valx.Transform(e, m)
	return m
}

// ListFeedbackCategory 获取列表
func (r feedbackRepo) ListFeedbackCategory(ctx kratosx.Context, req *biz.ListFeedbackCategoryRequest) ([]*biz.FeedbackCategory, uint32, error) {
	var (
		bs    []*biz.FeedbackCategory
		ms    []*model.FeedbackCategory
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.FeedbackCategory{}).Select(fs)

	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
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
		bs = append(bs, r.ToFeedbackCategoryEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateFeedbackCategory 创建数据
func (r feedbackRepo) CreateFeedbackCategory(ctx kratosx.Context, req *biz.FeedbackCategory) (uint32, error) {
	m := r.ToFeedbackCategoryModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// GetFeedbackCategory 获取指定的数据
func (r feedbackRepo) GetFeedbackCategory(ctx kratosx.Context, id uint32) (*biz.FeedbackCategory, error) {
	var (
		m  = model.FeedbackCategory{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToFeedbackCategoryEntity(&m), nil
}

// UpdateFeedbackCategory 更新数据
func (r feedbackRepo) UpdateFeedbackCategory(ctx kratosx.Context, req *biz.FeedbackCategory) error {
	return ctx.DB().Updates(r.ToFeedbackCategoryModel(req)).Error
}

// DeleteFeedbackCategory 删除数据
func (r feedbackRepo) DeleteFeedbackCategory(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.FeedbackCategory{})
	return uint32(db.RowsAffected), db.Error
}

// ToFeedbackEntity model转entity
func (r feedbackRepo) ToFeedbackEntity(m *model.Feedback) *biz.Feedback {
	e := &biz.Feedback{}
	_ = valx.Transform(m, e)
	return e
}

// ToFeedbackModel entity转model
func (r feedbackRepo) ToFeedbackModel(e *biz.Feedback) *model.Feedback {
	m := &model.Feedback{}
	_ = valx.Transform(e, m)
	return m
}

// ListFeedback 获取列表
func (r feedbackRepo) ListFeedback(ctx kratosx.Context, req *biz.ListFeedbackRequest) ([]*biz.Feedback, uint32, error) {
	var (
		bs    []*biz.Feedback
		ms    []*model.Feedback
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Feedback{}).Select(fs)
	db = db.Preload("App").Preload("User").Preload("Category")

	if req.AppId != nil {
		db = db.Where("app_id = ?", *req.AppId)
	}
	if req.CategoryId != nil {
		db = db.Where("category_id = ?", *req.CategoryId)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Platform != nil {
		db = db.Where("platform = ?", *req.Platform)
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
		b := r.ToFeedbackEntity(m)
		if m.Images != nil {
			var list []string
			if json.Unmarshal([]byte(*m.Images), &list) == nil {
				for _, image := range list {
					if url := resource.GetURLBySha(ctx, image); url != "" {
						b.ImageUrls = append(b.ImageUrls, url)
					}
				}
			}
		}

		bs = append(bs, b)
	}
	return bs, uint32(total), nil
}

// CreateFeedback 创建数据
func (r feedbackRepo) CreateFeedback(ctx kratosx.Context, req *biz.Feedback) (uint32, error) {
	m := r.ToFeedbackModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// DeleteFeedback 删除数据
func (r feedbackRepo) DeleteFeedback(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.Feedback{})
	return uint32(db.RowsAffected), db.Error
}

// GetFeedback 获取指定的数据
// func (r feedbackRepo) GetFeedback(ctx kratosx.Context, id uint32) (*biz.Feedback, error) {
//	var (
//		m  = model.Feedback{}
//		fs = []string{"*"}
//	)
//	db := ctx.DB().Select(fs)
//	db = db.Preload("App").Preload("User").Preload("FeedbackCategory")
//	if err := db.First(&m, id).Error; err != nil {
//		return nil, err
//	}
//
//	return r.ToFeedbackEntity(&m), nil
// }

// UpdateFeedback 更新数据
func (r feedbackRepo) UpdateFeedback(ctx kratosx.Context, req *biz.Feedback) error {
	return ctx.DB().Updates(r.ToFeedbackModel(req)).Error
}

func (r feedbackRepo) IsExistFeedbackByMd5(ctx kratosx.Context, md5 string) bool {
	var count int64
	ctx.DB().Model(model.Feedback{}).Where("md5=?", md5).Count(&count)
	return count != 0
}
