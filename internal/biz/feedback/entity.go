package feedback

type FeedbackCategory struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type App struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type Feedback struct {
	Id              uint32            `json:"id"`
	AppId           uint32            `json:"appId"`
	UserId          uint32            `json:"userId"`
	CategoryId      uint32            `json:"categoryId"`
	Title           string            `json:"title"`
	Content         string            `json:"content"`
	Status          string            `json:"status"`
	Images          *string           `json:"images"`
	ImageUrls       []string          `json:"imageUrls"`
	Contact         *string           `json:"contact"`
	Device          string            `json:"device"`
	Platform        string            `json:"platform"`
	Version         string            `json:"version"`
	Md5             string            `json:"md5"`
	ProcessedBy     *uint32           `json:"processedBy"`
	ProcessedResult *string           `json:"processedResult"`
	CreatedAt       int64             `json:"createdAt"`
	UpdatedAt       int64             `json:"updatedAt"`
	App             *App              `json:"app"`
	User            *User             `json:"user"`
	Category        *FeedbackCategory `json:"category"`
}

type User struct {
	Id       uint32 `json:"id"`
	RealName string `json:"realName"`
	NickName string `json:"nickName"`
}
