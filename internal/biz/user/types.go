package user

type PageUserRequest struct {
	Page     uint32   `json:"page"`
	PageSize uint32   `json:"page_size"`
	App      *string  `json:"app"`
	Username *string  `json:"username"`
	Phone    *string  `json:"phone"`
	Email    *string  `json:"email"`
	InIds    []uint32 `json:"in_ids"`
	NotInIds []uint32 `json:"not_in_ids"`
}

type OAuthLoginRequest struct {
	App      string `json:"app"`
	Code     string `json:"code"`
	Platform string `json:"platform"`
}

type OAuthBindByPasswordRequest struct {
	Username  string
	Password  string
	Captcha   string
	CaptchaID string
	App       string
	Code      string
	Platform  string
}

type OAuthBindByCaptchaRequest struct {
	Username  string
	Captcha   string
	CaptchaID string
	App       string
	Code      string
	Platform  string
}

type PasswordLoginRequest struct {
	App       string `json:"app"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captcha_id"`
}

type PasswordRegisterRequest struct {
	App       string `json:"app"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captcha_id"`
}

type CaptchaLoginRequest struct {
	App       string `json:"app"`
	Username  string `json:"username"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captcha_id"`
}

type CaptchaRegisterRequest struct {
	App       string `json:"app"`
	Username  string `json:"username"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captcha_id"`
}

type Password struct {
	Password string `json:"password"`
	Time     int64  `json:"time"`
}

type CaptchaResponse struct {
	ID     string `json:"id"`
	Expire int    `json:"expire"`
	Base64 string `json:"base64"`
}
