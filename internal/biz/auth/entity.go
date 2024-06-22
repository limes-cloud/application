package auth

type App struct {
	Id            uint32 `json:"id"`
	Logo          string `json:"logo"`
	Keyword       string `json:"keyword"`
	Name          string `json:"name"`
	AllowRegister *bool  `json:"allowRegister"`
}

type Auth struct {
	Id          uint32  `json:"id"`
	UserId      uint32  `json:"userId"`
	AppId       uint32  `json:"appId"`
	Status      *bool   `json:"status"`
	DisableDesc *string `json:"disableDesc"`
	Token       string  `json:"token"`
	LoggedAt    int64   `json:"loggedAt"`
	ExpiredAt   int64   `json:"expiredAt"`
	CreatedAt   int64   `json:"createdAt"`
	App         *App
}

type Channel struct {
	Id      uint32 `json:"id"`
	Ak      string `json:"ak"`
	Sk      string `json:"sk"`
	Extra   string `json:"extra"`
	Logo    string `json:"logo"`
	Keyword string `json:"keyword"`
	Name    string `json:"name"`
}

type OAuth struct {
	Id        uint32  `json:"id"`
	UserId    *uint32 `json:"userId"`
	ChannelId uint32  `json:"channelId"`
	AuthId    string  `json:"authId"`
	UnionId   *string `json:"unionId"`
	Token     string  `json:"token"`
	LoggedAt  int64   `json:"loggedAt"`
	ExpiredAt int64   `json:"expiredAt"`
	CreatedAt int64   `json:"createdAt"`
	Channel   *Channel
}

type User struct {
	Id       uint32   `json:"id"`
	Nickname string   `json:"nickname"`
	Email    *string  `json:"email"`
	Username *string  `json:"username"`
	Password *string  `json:"password"`
	Status   *bool    `json:"status"`
	From     string   `json:"from"`
	FromDesc string   `json:"from_desc"`
	AppIds   []uint32 `json:"appIds"`
}

type Password struct {
	Password string `json:"password"`
	Time     int64  `json:"time"`
}
