package types

type ListUserinfoRequest struct {
	Page       uint32  `json:"page"`
	PageSize   uint32  `json:"pageSize"`
	UserId     uint32  `json:"userId"`
	AppId      *uint32 `json:"appId"`
	AppIds     *uint32 `json:"appIds"`
	AppKeyword *string `json:"appKeyword"`
	Order      *string `json:"order"`
	OrderBy    *string `json:"orderBy"`
}
