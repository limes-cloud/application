package types

type Auth struct {
	UserID     uint32 `json:"user_id"`
	AppID      uint32 `json:"app_id"`
	AppKeyword string `json:"app_keyword"`
	ChannelID  uint32 `json:"channel_id"`
}
