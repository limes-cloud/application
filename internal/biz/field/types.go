package field

type PageFieldRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	Keyword  *string `json:"keyword"`
	Name     *string `json:"name"`
}

type FieldType struct {
	Type string
	Name string
}
