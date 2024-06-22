package userinfo

type Userinfo struct {
	Id        uint32 `json:"id"`
	UserId    uint32 `json:"userId"`
	Keyword   string `json:"keyword"`
	Value     string `json:"value"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
