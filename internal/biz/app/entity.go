package app

type App struct {
	Id            uint32        `json:"id"`
	Logo          string        `json:"logo"`
	LogoUrl       string        `json:"logoUrl"`
	Keyword       string        `json:"keyword"`
	Name          string        `json:"name"`
	Status        *bool         `json:"status"`
	DisableDesc   *string       `json:"disableDesc"`
	AllowRegistry *bool         `json:"allowRegistry"`
	Version       string        `json:"version"`
	Copyright     string        `json:"copyright"`
	Extra         *string       `json:"extra"`
	Description   *string       `json:"description"`
	CreatedAt     int64         `json:"createdAt"`
	UpdatedAt     int64         `json:"updatedAt"`
	AppChannels   []*AppChannel `json:"appChannels"`
	AppFields     []*AppField   `json:"appFields"`
	Channels      []*Channel    `json:"channels"`
	Fields        []*Field      `json:"fields"`
}

type AppChannel struct {
	ChannelId uint32 `json:"channelId"`
}

type AppField struct {
	FieldId uint32 `json:"fieldId"`
}

type Channel struct {
	Id      uint32 `json:"id"`
	Logo    string `json:"logo"`
	Name    string `json:"name"`
	Keyword string `json:"keyword"`
}

type Field struct {
	Id      uint32 `json:"id"`
	Keyword string `json:"keyword"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}
