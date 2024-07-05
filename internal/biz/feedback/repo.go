package feedback

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// ListFeedbackCategory 获取反馈建议分类列表
	ListFeedbackCategory(ctx kratosx.Context, req *ListFeedbackCategoryRequest) ([]*FeedbackCategory, uint32, error)

	// CreateFeedbackCategory 创建反馈建议分类
	CreateFeedbackCategory(ctx kratosx.Context, req *FeedbackCategory) (uint32, error)

	// UpdateFeedbackCategory 更新反馈建议分类
	UpdateFeedbackCategory(ctx kratosx.Context, req *FeedbackCategory) error

	// DeleteFeedbackCategory 删除反馈建议分类
	DeleteFeedbackCategory(ctx kratosx.Context, ids []uint32) (uint32, error)

	// IsExistFeedbackByMd5 是否存在反馈
	IsExistFeedbackByMd5(ctx kratosx.Context, md5 string) bool

	// ListFeedback 获取反馈建议列表
	ListFeedback(ctx kratosx.Context, req *ListFeedbackRequest) ([]*Feedback, uint32, error)

	// CreateFeedback 创建反馈建议
	CreateFeedback(ctx kratosx.Context, req *Feedback) (uint32, error)

	// DeleteFeedback 删除反馈建议
	DeleteFeedback(ctx kratosx.Context, ids []uint32) (uint32, error)

	// UpdateFeedback 更新反馈建议
	UpdateFeedback(ctx kratosx.Context, req *Feedback) error
}
