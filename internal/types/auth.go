package types

type ListAuthRequest struct {
	Page     uint32   `json:"page"`
	PageSize uint32   `json:"pageSize"`
	Order    *string  `json:"order"`
	OrderBy  *string  `json:"orderBy"`
	UserId   uint32   `json:"userId"`
	AppIds   []uint32 `json:"appIds"`
}

type UpdateAuthStatusRequest struct {
	Id          uint32
	Status      bool
	DisableDesc *string
}

type GenAuthCaptchaRequest struct {
	Type  string `json:"type"`
	Email string `json:"email"`
}

type GenAuthCaptchaResponse struct {
	Id     string `json:"id"`
	Expire int    `json:"expire"`
	Base64 string `json:"base64"`
}

type ListOAuthRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	UserId   uint32  `json:"userId"`
}

type UpdateOAuthStatusRequest struct {
	Id          uint32
	Status      bool
	DisableDesc *string
}

type OAuthLoginRequest struct {
	App     string `json:"app"`
	Code    string `json:"code"`
	Channel string `json:"channel"`
}

type OAuthLoginReply struct {
	IsBind   bool    `json:"isBind"`
	OAuthUid *string `json:"oAuthUid"`
	Token    *string `json:"token"`
	Expire   *uint32 `json:"expire"`
}

type EmailLoginRequest struct {
	Email     string `json:"email"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
	App       string `json:"app"`
}

type PasswordLoginRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
	App       string `json:"app"`
}

type PasswordLoginReply struct {
	Token  string `json:"token"`
	Expire uint32 `json:"expire"`
}

type EmailRegisterRequest struct {
	Email     string  `json:"email"`
	Captcha   string  `json:"captcha"`
	CaptchaId string  `json:"captchaId"`
	App       string  `json:"app"`
	OAuthUid  *string `json:"oAuthUid"`
}

type PasswordRegisterRequest struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Captcha   string  `json:"captcha"`
	CaptchaId string  `json:"captchaId"`
	App       string  `json:"app"`
	OAuthUid  *string `json:"oAuthUid"`
}

type EmailBindRequest struct {
	Email     string `json:"email"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
	App       string `json:"app"`
	OAuthUid  string `json:"oAuthUid"`
}

type PasswordBindRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
	App       string `json:"app"`
	OAuthUid  string `json:"oAuthUid"`
}

type TokenInfo struct {
	Token  string `json:"token"`
	Expire uint32 `json:"expire"`
}

type Password struct {
	Password string `json:"password"`
	Time     int64  `json:"time"`
}
