package agreement

type PageContentRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	Name     *string `json:"name"`
}

type PageSceneRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	Name     *string `json:"name"`
}
