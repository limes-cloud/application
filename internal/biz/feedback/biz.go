package feedback

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/manager/api/manager/auth"

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

// ListFeedbackCategory 获取反馈建议分类列表
func (u *UseCase) ListFeedbackCategory(ctx kratosx.Context, req *ListFeedbackCategoryRequest) ([]*FeedbackCategory, uint32, error) {
	list, total, err := u.repo.ListFeedbackCategory(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateFeedbackCategory 创建反馈建议分类
func (u *UseCase) CreateFeedbackCategory(ctx kratosx.Context, req *FeedbackCategory) (uint32, error) {
	id, err := u.repo.CreateFeedbackCategory(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateFeedbackCategory 更新反馈建议分类
func (u *UseCase) UpdateFeedbackCategory(ctx kratosx.Context, req *FeedbackCategory) error {
	if err := u.repo.UpdateFeedbackCategory(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteFeedbackCategory 删除反馈建议分类
func (u *UseCase) DeleteFeedbackCategory(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteFeedbackCategory(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// ListFeedback 获取反馈建议列表
func (u *UseCase) ListFeedback(ctx kratosx.Context, req *ListFeedbackRequest) ([]*Feedback, uint32, error) {
	list, total, err := u.repo.ListFeedback(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateFeedback 创建反馈建议
func (u *UseCase) CreateFeedback(ctx kratosx.Context, req *Feedback) (uint32, error) {
	content := req.Title + req.Content + req.Device
	if req.Images != nil {
		content += *req.Images
	}
	if req.Contact != nil {
		content += *req.Contact
	}
	req.UserId = md.UserID(ctx)
	req.Status = StatusUntreated
	req.Md5 = crypto.MD5([]byte(content))

	if u.repo.IsExistFeedbackByMd5(ctx, req.Md5) {
		return 0, errors.ExistFeedbackError()
	}

	id, err := u.repo.CreateFeedback(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// DeleteFeedback 删除反馈建议
func (u *UseCase) DeleteFeedback(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteFeedback(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// UpdateFeedback 更新反馈建议
func (u *UseCase) UpdateFeedback(ctx kratosx.Context, req *Feedback) error {
	adminInfo, err := auth.GetAuthInfo(ctx)
	if err != nil {
		return err
	}
	req.ProcessedBy = &adminInfo.UserId
	if err := u.repo.UpdateFeedback(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
