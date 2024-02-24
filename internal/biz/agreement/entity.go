package agreement

import ktypes "github.com/limes-cloud/kratosx/types"

type Content struct {
	ktypes.BaseModel
	Name        string `json:"name" gorm:"not null;size:32;comment:协议名称"`
	Status      *bool  `json:"status" gorm:"not null;comment:协议状态"`
	Content     string `json:"content" gorm:"type:blob;not null;comment:协议内容"`
	Description string `json:"description" gorm:"not null;size:128;comment:协议描述"`
}

func (t Content) TableName() string {
	return "agreement_content"
}

type Scene struct {
	ktypes.BaseModel
	Keyword       string          `json:"keyword" gorm:"unique;binary;not null;size:32;comment:场景标识"`
	Name          string          `json:"name" gorm:"not null;size:32;comment:场景名称"`
	Description   string          `json:"description" gorm:"not null;size:128;comment:场景描述"`
	Contents      []*Content      `json:"contents" gorm:"many2many:agreement_scene_content;constraint:onDelete:cascade"`
	SceneContents []*SceneContent `json:"scene_contents"`
}

func (t Scene) TableName() string {
	return "agreement_scene"
}

type SceneContent struct {
	SceneID   uint32 `json:"scene_id"`
	ContentID uint32 `json:"content_id"`
}

func (t SceneContent) TableName() string {
	return "agreement_scene_content"
}
