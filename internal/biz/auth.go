package biz

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/forgoer/openssl"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/biz/types"
	"github.com/limes-cloud/user-center/pkg/md"
	"github.com/limes-cloud/user-center/pkg/util"
)

const (
	loginImage    = "loginImage"
	bindImage     = "bindImage"
	registerImage = "registerImage"

	loginEmail    = "loginEmail"
	bindEmail     = "bindEmail"
	registerEmail = "registerEmail"

	passwordCert = "password"
	PasswordAuth = "password"
	CaptchaAuth  = "captcha"
)

type Auth struct {
	ktypes.BaseModel
	AppID     uint32   `json:"app_id" gorm:"uniqueIndex:auc;not null;comment:应用id"`
	UserID    uint32   `json:"user_id" gorm:"uniqueIndex:auc;not null;comment:用户id"`
	ChannelID uint32   `json:"channel_id" gorm:"uniqueIndex:auc;not null;comment:渠道id"`
	Token     string   `json:"token" gorm:"size:1024;comment:平台Token"`
	ExpireAt  int64    `json:"expire_at" gorm:"comment:过期时间"`
	App       *App     `json:"app" gorm:"constraint:onDelete:cascade"`
	User      *User    `json:"user" gorm:"constraint:onDelete:cascade"`
	Channel   *Channel `json:"channel" gorm:"constraint:onDelete:cascade"`
}

type AuthRepo interface {
	GetUserByUsername(ctx kratosx.Context, username string) (*User, error)
	GetUserByPhone(ctx kratosx.Context, username string) (*User, error)
	GetUserByEmail(ctx kratosx.Context, email string) (*User, error)
	HasUsername(ctx kratosx.Context, un string) bool
	HasUserEmail(ctx kratosx.Context, email string) bool
	GetChannel(ctx kratosx.Context, platform string) (*Channel, error)
	GetApp(ctx kratosx.Context, keyword string) (*App, error)
	Upsert(ctx kratosx.Context, c *Auth) error
	UpsertUserApp(ctx kratosx.Context, c *UserApp) error
	UpsertUserChannel(ctx kratosx.Context, c *UserChannel) error
	Register(ctx kratosx.Context, user *User) (uint32, error)
}

type AuthUseCase struct {
	config *config.Config
	repo   AuthRepo
}

func NewAuthUseCase(config *config.Config, repo AuthRepo) *AuthUseCase {
	return &AuthUseCase{config: config, repo: repo}
}

func (u *AuthUseCase) LoginPlatform() []*types.LoginPlatform {
	return []*types.LoginPlatform{
		{
			Platform: PasswordAuth,
			Name:     "密码登录",
		},
		{
			Platform: CaptchaAuth,
			Name:     "验证码登录",
		},
	}
}

func (u *AuthUseCase) Auth(ctx kratosx.Context, appId uint32) error {
	if md.AppID(ctx) != appId {
		return v1.ForbiddenError()
	}
	return nil
}

func (u *AuthUseCase) GenToken(ctx kratosx.Context, user *User, app *App, channel *Channel) (string, error) {
	// 如果应用不允许注册，则判断是否具有应用权限
	if app.AllowRegistry != nil && !*app.AllowRegistry {
		hasApp := false
		for _, item := range user.UserApps {
			if app.ID == item.ID {
				hasApp = true
			}
		}
		if !hasApp && app.AllowRegistry != nil && !*app.AllowRegistry {
			return "", v1.NotUserAppError()
		}
	}

	// 判断是否具有渠道权限
	hasChannel := false
	for _, item := range app.Channels {
		if channel.ID == item.ID {
			hasChannel = true
		}
	}
	if !hasChannel {
		return "", v1.NotUserChannelError()
	}

	// 生成登陆token
	token, err := ctx.JWT().NewToken(md.New(user.ID, app.ID, channel.ID))
	if err != nil {
		return "", v1.GenTokenErrorFormat(err.Error())
	}

	// 更新用户登录的token
	if err := u.repo.Upsert(ctx, &Auth{
		AppID:     app.ID,
		UserID:    user.ID,
		ChannelID: channel.ID,
		Token:     token,
		ExpireAt:  time.Now().Add(ctx.Config().App().JWT.Expire).Unix(),
	}); err != nil {
		return "", v1.DatabaseError()
	}

	// 更新用户的当前登录时间
	if err := u.repo.UpsertUserApp(ctx, &UserApp{
		UserID:  user.ID,
		AppID:   app.ID,
		LoginAt: time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()),
	}); err != nil {
		return "", v1.DatabaseErrorFormat(err.Error())
	}

	return token, nil
}

// LoginByPassword 通过账号密码登录
func (u *AuthUseCase) LoginByPassword(ctx kratosx.Context, req *types.LoginByPasswordRequest) (string, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyImage(loginImage, ctx.ClientIP(), req.CaptchaID, req.Captcha); err != nil {
		return "", v1.VerifyCaptchaError()
	}

	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(req.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		return "", v1.RsaDecodeErrorFormat(err.Error())
	}

	// 序列化密码
	var pw types.Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return "", v1.PasswordFormatError()
	}

	// 判断当前时间戳是否过期,超过3s则拒绝
	if time.Now().UnixMilli()-pw.Time > 3*1000 {
		return "", v1.PasswordExpireError()
	}

	// 获取用户信息
	var user *User
	if util.IsPhone(req.Username) {
		user, err = u.repo.GetUserByPhone(ctx, req.Username)
	} else if util.IsEmail(req.Username) {
		user, err = u.repo.GetUserByEmail(ctx, req.Username)
	} else {
		user, err = u.repo.GetUserByUsername(ctx, req.Username)
	}
	// 查询不到用户信息
	if err != nil {
		return "", v1.UsernameOrPasswordError()
	}

	// 用户被禁用则拒绝登陆
	if user.Status == nil && !*user.Status {
		if user.DisableDesc == nil {
			user.DisableDesc = proto.String("用户已被禁用")
		}
		return "", v1.UserDisableErrorFormat(*user.DisableDesc)
	}

	// 对比用户密码，错误则拒绝登陆
	if !util.CompareHashPwd(user.Password, pw.Password) {
		return "", v1.UsernameOrPasswordError()
	}

	// 查询当前用户的所属应用
	app, err := u.repo.GetApp(ctx, req.App)
	if err != nil {
		return "", v1.NotAppError()
	}

	// 查询当前用户登录的渠道
	channel, err := u.repo.GetChannel(ctx, CaptchaAuth)
	if err != nil {
		return "", v1.NotAppError()
	}

	// 更新完了用户渠道信息
	if err := u.repo.UpsertUserChannel(ctx, &UserChannel{
		UserID:    user.ID,
		ChannelID: channel.ID,
		AuthID:    req.Username,
		LoginAt:   time.Now().Unix(),
	}); err != nil {
		return "", v1.DatabaseError()
	}

	return u.GenToken(ctx, user, app, channel)
}

// LoginImageCaptcha 登录验证码
func (u *AuthUseCase) LoginImageCaptcha(ctx kratosx.Context) (*types.CaptchaResponse, error) {
	res, err := ctx.Captcha().Image(loginImage, ctx.ClientIP())
	if err != nil {
		return nil, v1.GenCaptchaErrorFormat(err.Error())
	}
	return &types.CaptchaResponse{
		ID:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}

// BindImageCaptcha 绑定验证码
func (u *AuthUseCase) BindImageCaptcha(ctx kratosx.Context) (*types.CaptchaResponse, error) {
	res, err := ctx.Captcha().Image(bindImage, ctx.ClientIP())
	if err != nil {
		return nil, v1.GenCaptchaErrorFormat(err.Error())
	}
	return &types.CaptchaResponse{
		ID:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}

// RegisterImageCaptcha 注册验证码
func (u *AuthUseCase) RegisterImageCaptcha(ctx kratosx.Context) (*types.CaptchaResponse, error) {
	res, err := ctx.Captcha().Image(registerImage, ctx.ClientIP())
	if err != nil {
		return nil, v1.GenCaptchaErrorFormat(err.Error())
	}
	return &types.CaptchaResponse{
		ID:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}

// RegisterByPassword 密码注册
func (u *AuthUseCase) RegisterByPassword(ctx kratosx.Context, req *types.RegisterByPasswordRequest) (*types.RegisterReply, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyImage(loginImage, ctx.ClientIP(), req.CaptchaID, req.Captcha); err != nil {
		return nil, v1.VerifyCaptchaError()
	}

	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(req.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		return nil, v1.RsaDecodeErrorFormat(err.Error())
	}

	// 序列化密码
	var pw types.Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return nil, v1.PasswordFormatError()
	}

	// 判断当前时间戳是否过期,超过3s则拒绝
	if time.Now().UnixMilli()-pw.Time > 3*1000 {
		return nil, v1.PasswordExpireError()
	}

	// 判断是否存在这个用户
	if u.repo.HasUsername(ctx, req.Username) {
		return nil, v1.AlreadyExistUsernameError()
	}

	// 获取应用信息
	app, err := u.repo.GetApp(ctx, req.App)
	if err != nil {
		return nil, v1.NotAppError()
	}

	// 判断app是否能注册
	if app.AllowRegistry != nil && !*app.AllowRegistry {
		return nil, v1.NotUserAppError()
	}

	// 注册
	id, err := u.repo.Register(ctx, &User{
		Username: &req.Username,
		Password: pw.Password,
		Status:   proto.Bool(true),
		UserApps: []*UserApp{{
			AppID: app.ID,
		}},
	})
	if err != nil {
		return nil, v1.RegisterError()
	}

	// 生成token
	user, _ := u.repo.GetUserByUsername(ctx, req.Username)

	// 获取当前渠道
	channel, _ := u.repo.GetChannel(ctx, CaptchaAuth)

	token, err := u.GenToken(ctx, user, app, channel)
	return &types.RegisterReply{ID: id, Token: token}, err
}

// RegisterUsernameCheck 注册用户名检测
func (u *AuthUseCase) RegisterUsernameCheck(ctx kratosx.Context, username string) bool {
	return !u.repo.HasUsername(ctx, username)
}

func (u *AuthUseCase) LoginByEmail(ctx kratosx.Context, req *types.LoginByEmailRequest) (string, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyEmail(loginEmail, ctx.ClientIP(), req.CaptchaID, req.Captcha); err != nil {
		return "", v1.VerifyCaptchaError()
	}

	// 查询用户信息
	user, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", v1.NotExistEmailError()
	}
	if user.Status == nil || !*user.Status {
		if user.DisableDesc == nil {
			user.DisableDesc = proto.String("用户已被禁用")
		}
		return "", v1.UserDisableErrorFormat(*user.DisableDesc)
	}

	// 查询当前用户的所属应用
	app, err := u.repo.GetApp(ctx, req.App)
	if err != nil {
		return "", v1.NotAppError()
	}

	// 查询当前用户登录的渠道
	channel, err := u.repo.GetChannel(ctx, CaptchaAuth)
	if err != nil {
		return "", v1.NotAppError()
	}

	// 更新完了用户渠道信息
	if err := u.repo.UpsertUserChannel(ctx, &UserChannel{
		UserID:    user.ID,
		ChannelID: channel.ID,
		AuthID:    req.Email,
		LoginAt:   time.Now().Unix(),
	}); err != nil {
		return "", v1.DatabaseError()
	}

	return u.GenToken(ctx, user, app, channel)
}

// LoginEmailCaptcha 登录验证码
func (u *AuthUseCase) LoginEmailCaptcha(ctx kratosx.Context, email string) (*types.CaptchaResponse, error) {
	// 查询不到用户信息
	if !u.repo.HasUserEmail(ctx, email) {
		return nil, v1.NotExistEmailError()
	}
	res, err := ctx.Captcha().Email(loginEmail, ctx.ClientIP(), email)
	if err != nil {
		return nil, v1.GenCaptchaErrorFormat(err.Error())
	}
	return &types.CaptchaResponse{
		ID:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}

// BindEmailCaptcha 绑定验证码
func (u *AuthUseCase) BindEmailCaptcha(ctx kratosx.Context, email string) (*types.CaptchaResponse, error) {
	// 查询不到用户信息
	if !u.repo.HasUserEmail(ctx, email) {
		return nil, v1.NotExistEmailError()
	}
	res, err := ctx.Captcha().Email(bindEmail, ctx.ClientIP(), email)
	if err != nil {
		return nil, v1.GenCaptchaErrorFormat(err.Error())
	}
	return &types.CaptchaResponse{
		ID:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}

// RegisterEmailCaptcha 注册验证码
func (u *AuthUseCase) RegisterEmailCaptcha(ctx kratosx.Context, email string) (*types.CaptchaResponse, error) {
	res, err := ctx.Captcha().Email(registerEmail, ctx.ClientIP(), email)
	if err != nil {
		return nil, v1.GenCaptchaErrorFormat(err.Error())
	}
	return &types.CaptchaResponse{
		ID:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}

// RegisterByEmail 密码注册
func (u *AuthUseCase) RegisterByEmail(ctx kratosx.Context, req *types.RegisterByEmailRequest) (*types.RegisterReply, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyEmail(registerEmail, ctx.ClientIP(), req.CaptchaID, req.Captcha); err != nil {
		return nil, v1.VerifyCaptchaError()
	}

	// 判断是否存在这个用户
	if u.repo.HasUserEmail(ctx, req.Email) {
		return nil, v1.AlreadyExistEmailError()
	}

	// 获取应用信息
	app, err := u.repo.GetApp(ctx, req.App)
	if err != nil {
		return nil, v1.NotAppError()
	}

	// 判断app是否能注册
	if app.AllowRegistry != nil && !*app.AllowRegistry {
		return nil, v1.NotUserAppError()
	}

	// 注册
	id, err := u.repo.Register(ctx, &User{
		Email:  &req.Email,
		Status: proto.Bool(true),
		UserApps: []*UserApp{{
			AppID: app.ID,
		}},
	})
	if err != nil {
		return nil, v1.RegisterError()
	}

	// 获取用户信息
	user, _ := u.repo.GetUserByEmail(ctx, req.Email)

	// 获取当前渠道
	channel, _ := u.repo.GetChannel(ctx, CaptchaAuth)

	// 生成token
	token, err := u.GenToken(ctx, user, app, channel)
	return &types.RegisterReply{ID: id, Token: token}, err
}
