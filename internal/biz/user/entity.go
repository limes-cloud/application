package user

type User struct {
	Id          uint32  `json:"id"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	Username    *string `json:"username"`
	Password    *string `json:"password"`
	NickName    string  `json:"nickName"`
	RealName    *string `json:"realName"`
	Avatar      *string `json:"avatar"`
	AvatarUrl   *string `json:"avatarUrl"`
	Gender      *string `json:"gender"`
	Status      *bool   `json:"status"`
	DisableDesc *string `json:"disableDesc"`
	From        string  `json:"from"`
	FromDesc    string  `json:"fromDesc"`
	CreatedAt   int64   `json:"createdAt"`
	UpdatedAt   int64   `json:"updatedAt"`
	DeletedAt   int64   `json:"deletedAt"`
}
