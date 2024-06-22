package field

import (
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/pkg/field"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// ListFieldType 获取支持的全部字段类型
func (u *UseCase) ListFieldType() []*FieldType {
	var list []*FieldType
	tps := field.New().GetTypes()

	for key, tp := range tps {
		list = append(list, &FieldType{Type: key, Name: tp.Name()})
	}
	return list
}

// ListField 获取用户字段列表
func (u *UseCase) ListField(ctx kratosx.Context, req *ListFieldRequest) ([]*Field, uint32, error) {
	list, total, err := u.repo.ListField(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateField 创建用户字段
func (u *UseCase) CreateField(ctx kratosx.Context, req *Field) (uint32, error) {
	req.Status = proto.Bool(false)
	id, err := u.repo.CreateField(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateField 更新用户字段
func (u *UseCase) UpdateField(ctx kratosx.Context, req *Field) error {
	if err := u.repo.UpdateField(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateFieldStatus 更新用户字段状态
func (u *UseCase) UpdateFieldStatus(ctx kratosx.Context, id uint32, status bool) error {
	if err := u.repo.UpdateFieldStatus(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteField 删除用户字段
func (u *UseCase) DeleteField(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteField(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}
