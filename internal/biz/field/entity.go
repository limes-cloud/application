package field

import ktypes "github.com/limes-cloud/kratosx/types"

type Field struct {
	ktypes.BaseModel
	Keyword     string  `json:"keyword"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
