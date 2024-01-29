package types

type PageAgreementRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	Name     *string `json:"name"`
}
