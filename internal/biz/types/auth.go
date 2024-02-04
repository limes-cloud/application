package types

type LoginPlatform struct {
	Platform string
	Name     string
}

type LoginByPasswordRequest struct {
	App       string `json:"app"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captcha_id"`
}

type RegisterByPasswordRequest struct {
	App       string `json:"app"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captcha_id"`
}

type RegisterReply struct {
	ID    uint32 `json:"id"`
	Token string `json:"token"`
}

type LoginByEmailRequest struct {
	App       string `json:"app"`
	Email     string `json:"email"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captcha_id"`
}

type RegisterByEmailRequest struct {
	App       string `json:"app"`
	Email     string `json:"email"`
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

type AuthRequest struct {
	Path   string
	Method string
}
