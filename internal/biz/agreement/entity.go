package agreement

import ktypes "github.com/limes-cloud/kratosx/types"

type Content struct {
	ktypes.BaseModel
	Name        string `json:"name"`
	Status      *bool  `json:"status"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

func (t Content) TableName() string {
	return "agreement_content"
}

type Scene struct {
	ktypes.BaseModel
	Keyword       string          `json:"keyword"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Contents      []*Content      `json:"contents" gorm:"many2many:agreement_scene_content"`
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
