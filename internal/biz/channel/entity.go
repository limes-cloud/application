package channel

type Channel struct {
	Id        uint32  `json:"id"`
	Logo      string  `json:"logo"`
	LogoUrl   string  `json:"logoUrl"`
	Keyword   string  `json:"keyword"`
	Name      string  `json:"name"`
	Status    *bool   `json:"status"`
	Ak        *string `json:"ak"`
	Sk        *string `json:"sk"`
	Extra     *string `json:"extra"`
	CreatedAt int64   `json:"createdAt"`
	UpdatedAt int64   `json:"updatedAt"`
}

type Typer struct {
	Keyword string `json:"keyword"`
	Name    string `json:"name"`
}
