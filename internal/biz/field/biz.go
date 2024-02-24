package field

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/user-center/api/errors"
	"github.com/limes-cloud/user-center/internal/config"
	"github.com/limes-cloud/user-center/internal/pkg/field"
)

type UseCase struct {
	config *config.Config
	repo   Repo
}

func NewUseCase(config *config.Config, repo Repo) *UseCase {
	return &UseCase{config: config, repo: repo}
}

// TypeList 获取全部字段类型
func (u *UseCase) TypeList() []*FieldType {
	var list []*FieldType
	tps := field.New().GetTypes()

	for key, tp := range tps {
		list = append(list, &FieldType{Type: key, Name: tp.Name()})
	}
	return list
}

func (u *UseCase) TypeSet(ctx kratosx.Context) map[string]*Field {
	m := map[string]*Field{}
	efs, err := u.repo.All(ctx)
	if err != nil {
		return m
	}
	for _, ef := range efs {
		m[ef.Keyword] = ef
	}
	return m
}

// Page 获取全部扩展字段
func (u *UseCase) Page(ctx kratosx.Context, req *PageFieldRequest) ([]*Field, uint32, error) {
	ef, total, err := u.repo.Page(ctx, req)
	if err != nil {
		return nil, 0, errors.NotRecord()
	}
	return ef, total, nil
}

// Add 添加扩展字段信息
func (u *UseCase) Add(ctx kratosx.Context, ef *Field) (uint32, error) {
	id, err := u.repo.Create(ctx, ef)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	return id, nil
}

// Update 更新扩展字段信息
func (u *UseCase) Update(ctx kratosx.Context, ef *Field) error {
	if err := u.repo.Update(ctx, ef); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// Delete 删除扩展字段信息
func (u *UseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}
