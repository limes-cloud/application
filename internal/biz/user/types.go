package user

type GetUserRequest struct {
	Id       *uint32 `json:"id"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
	Username *string `json:"username"`
}

type ListUserRequest struct {
	Page       uint32  `json:"page"`
	PageSize   uint32  `json:"pageSize"`
	Order      *string `json:"order"`
	OrderBy    *string `json:"orderBy"`
	Phone      *string `json:"phone"`
	Email      *string `json:"email"`
	Username   *string `json:"username"`
	RealName   *string `json:"realName"`
	Gender     *string `json:"gender"`
	Status     *bool   `json:"status"`
	From       *string `json:"from"`
	CreatedAts []int64 `json:"createdAts"`
	AppId      *uint32 `json:"appId"`
}

type ListTrashUserRequest struct {
	Page       uint32  `json:"page"`
	PageSize   uint32  `json:"pageSize"`
	Order      *string `json:"order"`
	OrderBy    *string `json:"orderBy"`
	Phone      *string `json:"phone"`
	Email      *string `json:"email"`
	Username   *string `json:"username"`
	RealName   *string `json:"realName"`
	Gender     *string `json:"gender"`
	Status     *bool   `json:"status"`
	From       *string `json:"from"`
	CreatedAts []int64 `json:"createdAts"`
	AppId      *uint32 `json:"appId"`
}

type ExportUserRequest struct {
	Phone      *string `json:"phone"`
	Email      *string `json:"email"`
	Username   *string `json:"username"`
	RealName   *string `json:"realName"`
	Gender     *string `json:"gender"`
	Status     *bool   `json:"status"`
	From       *string `json:"from"`
	CreatedAts []int64 `json:"createdAts"`
	AppId      *uint32 `json:"appId"`
}

type UpdateUserStatusRequest struct {
	Id          uint32  `json:"id"`
	Status      bool    `json:"status"`
	DisableDesc *string `json:"disableDesc"`
}

type UpdateCurrentUserRequest struct {
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Gender   string `json:"gender"`
}
