package types

type PageExtraFieldRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	Keyword  *string `json:"keyword"`
	Name     *string `json:"name"`
}

type ExtraFieldType struct {
	Type string
	Name string
}
