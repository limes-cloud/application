package types

type PageUserRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	App      *string `json:"app"`
	Username *string `json:"username"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
	IdCard   *string `json:"id_card"`
}
