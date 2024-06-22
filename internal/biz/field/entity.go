package field

type Field struct {
	Id          uint32  `json:"id"`
	Keyword     string  `json:"keyword"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Status      *bool   `json:"status"`
	Description *string `json:"description"`
	CreatedAt   int64   `json:"createdAt"`
	UpdatedAt   int64   `json:"updatedAt"`
}
