package auth

type ListAuthRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	UserId   uint32  `json:"userId"`
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

type EmailLoginReply struct {
	Token  string `json:"token"`
	Expire uint32 `json:"expire"`
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

type EmailRegisterReply struct {
	Token  string `json:"token"`
	Expire uint32 `json:"expire"`
}

type PasswordRegisterRequest struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Captcha   string  `json:"captcha"`
	CaptchaId string  `json:"captchaId"`
	App       string  `json:"app"`
	OAuthUid  *string `json:"oAuthUid"`
}

type PasswordRegisterReply struct {
	Token  string `json:"token"`
	Expire uint32 `json:"expire"`
}

type EmailBindRequest struct {
	Email     string `json:"email"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
	App       string `json:"app"`
	OAuthUid  string `json:"oAuthUid"`
}

type EmailBindReply struct {
	Token  string `json:"token"`
	Expire uint32 `json:"expire"`
}

type PasswordBindRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
	App       string `json:"app"`
	OAuthUid  string `json:"oAuthUid"`
}

type PasswordBindReply struct {
	Token  string `json:"token"`
	Expire uint32 `json:"expire"`
}
