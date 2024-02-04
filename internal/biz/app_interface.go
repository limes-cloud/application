package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"

	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/pkg/tree"
)

type AppInterface struct {
	ktypes.CreateModel
	AppID       uint32          `json:"app_id" gorm:"not null;comment:应用id"`
	ParentID    uint32          `json:"parent_id" gorm:"not null;comment:父id"`
	Type        string          `json:"type" gorm:"not null;type:char(32);comment:类型"`
	Title       string          `json:"title" gorm:"not null;size:128;comment:接口标题"`
	Path        *string         `json:"path" gorm:"uniqueIndex:pm;size:128;comment:接口路径"`
	Method      *string         `json:"method" gorm:"uniqueIndex:pm;size:32;comment:接口方法"`
	Description *string         `json:"description" gorm:"size:128;comment:接口描述"`
	Children    []*AppInterface `json:"children" gorm:"-"`
}

// ID 获取菜单树ID
func (m *AppInterface) ID() uint32 {
	return m.CreateModel.ID
}

// Parent 获取父ID
func (m *AppInterface) Parent() uint32 {
	return m.ParentID
}

// AppendChildren 添加子节点
func (m *AppInterface) AppendChildren(child any) {
	ai := child.(*AppInterface)
	m.Children = append(m.Children, ai)
}

// ChildrenNode 获取子节点
func (m *AppInterface) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range m.Children {
		list = append(list, item)
	}
	return list
}

type AppInterfaceRepo interface {
	All(ctx kratosx.Context, appId uint32) ([]*AppInterface, error)
	Create(ctx kratosx.Context, c *AppInterface) (uint32, error)
	Update(ctx kratosx.Context, c *AppInterface) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}

type AppInterfaceUseCase struct {
	config *config.Config
	repo   AppInterfaceRepo
}

func NewAppInterfaceUseCase(config *config.Config, repo AppInterfaceRepo) *AppInterfaceUseCase {
	return &AppInterfaceUseCase{config: config, repo: repo}
}

// Tree 获取全部应用接口
func (u *AppInterfaceUseCase) Tree(ctx kratosx.Context, appId uint32) ([]tree.Tree, error) {
	list, err := u.repo.All(ctx, appId)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	var (
		ais   []tree.Tree
		roots []tree.Tree
		rids  []uint32
	)

	// 构建树枝，并选取根节点
	for _, item := range list {
		ais = append(ais, item)
		if item.ParentID == 0 {
			rids = append(rids, item.ID())
		}
	}

	// 通过根及诶单构造树
	for _, rid := range rids {
		root := tree.BuildTreeByID(ais, rid)
		roots = append(roots, root)
	}

	return roots, nil
}

// Add 添加应用接口信息
func (u *AppInterfaceUseCase) Add(ctx kratosx.Context, app *AppInterface) (uint32, error) {
	id, err := u.repo.Create(ctx, app)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新应用接口信息
func (u *AppInterfaceUseCase) Update(ctx kratosx.Context, app *AppInterface) error {
	if err := u.repo.Update(ctx, app); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除应用接口信息
func (u *AppInterfaceUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
