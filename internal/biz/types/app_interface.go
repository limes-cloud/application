package types

type PageAppInterfaceRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	Title    *string `json:"title"`
}
