package user

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/forgoer/openssl"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/user-center/api/errors"
	"github.com/limes-cloud/user-center/internal/biz/app"
	"github.com/limes-cloud/user-center/internal/biz/channel"
	"github.com/limes-cloud/user-center/internal/consts"
	"github.com/limes-cloud/user-center/internal/pkg/authorizer"
	"github.com/limes-cloud/user-center/internal/pkg/md"
	"github.com/limes-cloud/user-center/types"
)

const (
	passwordCert = "password"

	loginImage    = "loginImage"
	bindImage     = "bindImage"
	registerImage = "registerImage"

	loginCaptcha    = "loginEmail"
	bindCaptcha     = "bindEmail"
	registerCaptcha = "registerEmail"
)

func (u *UseCase) bind(ctx kratosx.Context, user *User, app, platform, code string) (string, error) {
	// 获取应用和渠道信息
	curApp, err := u.repo.GetApp(ctx, app)
	if err != nil {
		return "", errors.NotApp()
	}
	var ac *channel.Channel
	for _, item := range curApp.Channels {
		if item.Platform == platform {
			ac = item
		}
	}
	if ac == nil {
		return "", errors.NotAppChannel()
	}

	af := authorizer.New(&authorizer.Config{
		Platform: ac.Platform,
		Ak:       ac.Ak,
		Sk:       ac.Sk,
		Code:     code,
	})

	ti, err := af.GetToken(ctx)
	if err != nil {
		return "", errors.GetAuthInfoFormat(err.Error())
	}
	ai, err := af.GetAuthInfo(ctx, ti.Token)
	if err != nil {
		return "", errors.GetAuthInfoFormat(err.Error())
	}

	if _, err := u.repo.AddAuth(ctx, &Auth{
		UserID:    user.ID,
		AuthID:    &ai.AuthId,
		UnionID:   ai.UnionId,
		ChannelID: ac.ID,
		LoginAt:   time.Now().Unix(),
	}); err != nil {
		return "", errors.Database()
	}
	// 生成用户token
	token, err := u.GenToken(ctx, user, curApp, ac)
	if err != nil {
		return "", errors.GenToken()
	}
	return token, nil
}

func (u *UseCase) captcha(ctx kratosx.Context, tp string) (*CaptchaResponse, error) {
	res, err := ctx.Captcha().Image(tp, ctx.ClientIP())
	if err != nil {
		return nil, errors.GenCaptchaFormat(err.Error())
	}
	return &CaptchaResponse{
		ID:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}

func (u *UseCase) email(ctx kratosx.Context, tp, email string) (*CaptchaResponse, error) {
	res, err := ctx.Captcha().Email(tp, ctx.ClientIP(), email)
	if err != nil {
		return nil, errors.GenCaptchaFormat(err.Error())
	}
	return &CaptchaResponse{
		ID:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}

// ParseToken 解析token数据
func (u *UseCase) ParseToken(ctx kratosx.Context) (*types.Auth, error) {
	data, err := md.Get(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GenToken 生成token数据
func (u *UseCase) GenToken(ctx kratosx.Context, user *User, app *app.App, channel *channel.Channel) (string, error) {
	// 如果应用不允许注册，则判断是否具有应用权限
	if app.AllowRegistry != nil && !*app.AllowRegistry {
		hasApp := false
		for _, item := range user.UserApps {
			if app.ID == item.AppID {
				hasApp = true
			}
		}
		if !hasApp && app.AllowRegistry != nil && !*app.AllowRegistry {
			return "", errors.NotUserApp()
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
		return "", errors.NotUserChannel()
	}

	// 生成登陆token
	token, err := ctx.JWT().NewToken(md.New(user.ID, app.ID, channel.ID, app.Keyword))
	if err != nil {
		return "", errors.GenTokenFormat(err.Error())
	}

	// 更新用户登录的token
	if err := u.repo.UpdateAuthByCU(ctx, &Auth{
		UserID:      user.ID,
		ChannelID:   channel.ID,
		JwtToken:    token,
		JwtExpireAt: time.Now().Add(ctx.Config().App().JWT.Expire).Unix(),
		LoginAt:     time.Now().Unix(),
	}); err != nil {
		return "", errors.Database()
	}

	// 更新用户的当前登录时间
	if err := u.repo.UpdateUserApp(ctx, &UserApp{
		UserID:  user.ID,
		AppID:   app.ID,
		LoginAt: time.Now().Unix(),
	}); err != nil {
		return "", errors.DatabaseFormat(err.Error())
	}

	return token, nil
}

// OAuthLogin 三方授权登录
func (u *UseCase) OAuthLogin(ctx kratosx.Context, in *OAuthLoginRequest) (string, error) {
	curApp, err := u.repo.GetApp(ctx, in.App)
	if err != nil {
		return "", errors.NotApp()
	}

	var ac *channel.Channel
	for _, item := range curApp.Channels {
		if item.Platform == in.Platform {
			ac = item
		}
	}
	if ac == nil {
		return "", errors.NotAppChannel()
	}

	// 获取授权信息
	af := authorizer.New(&authorizer.Config{
		Platform: ac.Platform,
		Ak:       ac.Ak,
		Sk:       ac.Sk,
		Code:     in.Code,
	})

	ti, err := af.GetToken(ctx)
	if err != nil {
		return "", errors.GetAuthInfoFormat(err.Error())
	}
	ai, err := af.GetAuthInfo(ctx, ti.Token)
	if err != nil {
		return "", errors.GetAuthInfoFormat(err.Error())
	}

	// 获取授权信息
	uc, err := u.repo.GetAuthByCA(ctx, ac.ID, ai.AuthId)
	if err != nil {
		return "", errors.UnBind()
	}

	// 获取用户信息
	user, err := u.repo.Get(ctx, uc.UserID)
	if err != nil {
		return "", errors.NotUser()
	}

	// 生成用户token
	token, err := u.GenToken(ctx, user, curApp, ac)
	if err != nil {
		return "", errors.GenToken()
	}
	return token, nil
}

// OAuthBindByPassword 通过密码三方账号绑定
func (u *UseCase) OAuthBindByPassword(ctx kratosx.Context, in *OAuthBindByPasswordRequest) (string, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyImage(bindImage, ctx.ClientIP(), in.CaptchaID, in.Captcha); err != nil {
		return "", errors.VerifyCaptcha()
	}

	passByte, _ := base64.StdEncoding.DecodeString(in.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		return "", errors.RsaDecodeFormat(err.Error())
	}

	var pw Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return "", errors.PasswordFormat()
	}

	// 判断当前时间戳是否过期,超过3s则拒绝
	if time.Now().UnixMilli()-pw.Time > 3*1000 {
		return "", errors.PasswordExpire()
	}

	// 获取用户信息
	var user *User
	if util.IsPhone(in.Username) {
		user, err = u.repo.GetByPhone(ctx, in.Username)
	} else if util.IsEmail(in.Username) {
		user, err = u.repo.GetByEmail(ctx, in.Username)
	} else {
		user, err = u.repo.GetByUsername(ctx, in.Username)
	}
	// 查询不到用户信息
	if err != nil {
		return "", errors.UsernameOrPassword()
	}

	// 用户被禁用则拒绝登陆
	if user.Status == nil && !*user.Status {
		if user.DisableDesc == nil {
			user.DisableDesc = proto.String("用户已被禁用")
		}
		return "", errors.UserDisableFormat(*user.DisableDesc)
	}

	// 对比用户密码，错误则拒绝登陆
	if !util.CompareHashPwd(user.Password, pw.Password) {
		return "", errors.UsernameOrPassword()
	}

	// 对比用户密码
	return u.bind(ctx, user, in.App, in.Platform, in.Code)
}

// OAuthBindCaptcha 绑定验证码
func (u *UseCase) OAuthBindCaptcha(ctx kratosx.Context) (*CaptchaResponse, error) {
	return u.captcha(ctx, bindImage)
}

// OAuthBindEmail 邮件绑定验证码
func (u *UseCase) OAuthBindEmail(ctx kratosx.Context, email string) (*CaptchaResponse, error) {
	if !u.repo.HasUserEmail(ctx, email) {
		return nil, errors.NotExistEmail()
	}
	return u.email(ctx, bindCaptcha, email)
}

// OAuthBindByCaptcha 通过验证码绑定三方账户
func (u *UseCase) OAuthBindByCaptcha(ctx kratosx.Context, in *OAuthBindByCaptchaRequest) (string, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyEmail(bindCaptcha, ctx.ClientIP(), in.CaptchaID, in.Captcha); err != nil {
		return "", errors.VerifyCaptcha()
	}

	// 获取用户信息
	var user *User
	var err error
	if util.IsPhone(in.Username) {
		user, err = u.repo.GetByPhone(ctx, in.Username)
	} else {
		user, err = u.repo.GetByEmail(ctx, in.Username)
	}
	if err != nil {
		return "", errors.UsernameOrPassword()
	}

	// 对比用户密码
	return u.bind(ctx, user, in.App, in.Platform, in.Code)
}

// PasswordLogin 通过账号密码登录
func (u *UseCase) PasswordLogin(ctx kratosx.Context, in *PasswordLoginRequest) (string, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyImage(loginImage, ctx.ClientIP(), in.CaptchaID, in.Captcha); err != nil {
		return "", errors.VerifyCaptcha()
	}

	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(in.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		return "", errors.RsaDecodeFormat(err.Error())
	}

	// 序列化密码
	var pw Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return "", errors.PasswordFormat()
	}

	// 判断当前时间戳是否过期,超过3s则拒绝
	if time.Now().UnixMilli()-pw.Time > 3*1000 {
		return "", errors.PasswordExpire()
	}

	// 获取用户信息
	var user *User
	if util.IsPhone(in.Username) {
		user, err = u.repo.GetByPhone(ctx, in.Username)
	} else if util.IsEmail(in.Username) {
		user, err = u.repo.GetByEmail(ctx, in.Username)
	} else {
		user, err = u.repo.GetByUsername(ctx, in.Username)
	}
	// 查询不到用户信息
	if err != nil {
		return "", errors.UsernameOrPassword()
	}

	// 用户被禁用则拒绝登陆
	if user.Status == nil && !*user.Status {
		if user.DisableDesc == nil {
			user.DisableDesc = proto.String("用户已被禁用")
		}
		return "", errors.UserDisableFormat(*user.DisableDesc)
	}

	// 对比用户密码，错误则拒绝登陆
	if !util.CompareHashPwd(user.Password, pw.Password) {
		return "", errors.UsernameOrPassword()
	}

	// 查询当前用户的所属应用
	curApp, err := u.repo.GetApp(ctx, in.App)
	if err != nil {
		return "", errors.NotApp()
	}

	var curChannel *channel.Channel
	for _, item := range curApp.Channels {
		if item.Platform == consts.PasswordChannel {
			curChannel = item
		}
	}
	if curChannel == nil {
		return "", errors.NotAppChannel()
	}

	return u.GenToken(ctx, user, curApp, curChannel)
}

// PasswordLoginCaptcha 登录验证码
func (u *UseCase) PasswordLoginCaptcha(ctx kratosx.Context) (*CaptchaResponse, error) {
	return u.captcha(ctx, loginImage)
}

// PasswordRegister 密码注册
func (u *UseCase) PasswordRegister(ctx kratosx.Context, in *PasswordRegisterRequest) (string, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyImage(registerImage, ctx.ClientIP(), in.CaptchaID, in.Captcha); err != nil {
		return "", errors.VerifyCaptcha()
	}

	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(in.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		return "", errors.RsaDecodeFormat(err.Error())
	}

	// 序列化密码
	var pw Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return "", errors.PasswordFormat()
	}

	// 判断当前时间戳是否过期,超过3s则拒绝
	if time.Now().UnixMilli()-pw.Time > 3*1000 {
		return "", errors.PasswordExpire()
	}

	// 判断是否存在这个用户
	if u.repo.HasUsername(ctx, in.Username) {
		return "", errors.AlreadyExistUsername()
	}

	// 获取应用信息
	curApp, err := u.repo.GetApp(ctx, in.App)
	if err != nil {
		return "", errors.NotApp()
	}

	// 判断app是否能注册
	if curApp.AllowRegistry != nil && !*curApp.AllowRegistry {
		return "", errors.NotUserApp()
	}

	// 判断app是否开启密码注册
	var curChannel *channel.Channel
	for _, item := range curApp.Channels {
		if item.Platform == consts.PasswordChannel {
			curChannel = item
		}
	}
	if curChannel == nil {
		return "", errors.NotAppChannel()
	}

	// 注册
	user := &User{
		Username: &in.Username,
		Password: pw.Password,
		Status:   proto.Bool(true),
		UserApps: []*UserApp{{
			AppID: curApp.ID,
		}},
		From:     curApp.Keyword,
		FromDesc: curApp.Name,
	}
	if _, err := u.repo.Add(ctx, user); err != nil {
		return "", errors.Register()
	}

	return u.GenToken(ctx, user, curApp, curChannel)
}

// PasswordRegisterCaptcha 注册验证码
func (u *UseCase) PasswordRegisterCaptcha(ctx kratosx.Context) (*CaptchaResponse, error) {
	return u.captcha(ctx, registerImage)
}

// PasswordRegisterCheck 密码注册用户名检测
func (u *UseCase) PasswordRegisterCheck(ctx kratosx.Context, username string) bool {
	return !u.repo.HasUsername(ctx, username)
}

func (u *UseCase) CaptchaLogin(ctx kratosx.Context, in *CaptchaLoginRequest) (string, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyEmail(loginCaptcha, ctx.ClientIP(), in.CaptchaID, in.Captcha); err != nil {
		return "", errors.VerifyCaptcha()
	}

	// 查询用户信息
	// 获取用户信息
	var user *User
	var err error
	if util.IsPhone(in.Username) {
		user, err = u.repo.GetByPhone(ctx, in.Username)
	} else {
		user, err = u.repo.GetByEmail(ctx, in.Username)
	}
	if err != nil {
		return "", errors.UsernameOrPassword()
	}

	if user.Status == nil || !*user.Status {
		if user.DisableDesc == nil {
			user.DisableDesc = proto.String("用户已被禁用")
		}
		return "", errors.UserDisableFormat(*user.DisableDesc)
	}

	// 查询当前用户的所属应用
	curApp, err := u.repo.GetApp(ctx, in.App)
	if err != nil {
		return "", errors.NotApp()
	}
	// 判断app是否开启密码注册
	var curChannel *channel.Channel
	for _, item := range curApp.Channels {
		if item.Platform == consts.PasswordChannel {
			curChannel = item
		}
	}
	if curChannel == nil {
		return "", errors.NotAppChannel()
	}

	return u.GenToken(ctx, user, curApp, curChannel)
}

// CaptchaLoginEmail 登录邮件验证码
func (u *UseCase) CaptchaLoginEmail(ctx kratosx.Context, email string) (*CaptchaResponse, error) {
	// 查询不到用户信息
	if !u.repo.HasUserEmail(ctx, email) {
		return nil, errors.NotExistEmail()
	}
	return u.email(ctx, loginCaptcha, email)
}

// CaptchaRegisterEmail 注册验证码
func (u *UseCase) CaptchaRegisterEmail(ctx kratosx.Context, email string) (*CaptchaResponse, error) {
	return u.email(ctx, registerCaptcha, email)
}

// CaptchaRegister 邮件注册
func (u *UseCase) CaptchaRegister(ctx kratosx.Context, in *CaptchaRegisterRequest) (string, error) {
	// 判断验证码是否正确
	if err := ctx.Captcha().VerifyEmail(registerCaptcha, ctx.ClientIP(), in.CaptchaID, in.Captcha); err != nil {
		return "", errors.VerifyCaptcha()
	}

	// 判断是否存在这个用户
	var phone *string
	var email *string
	if util.IsPhone(in.Username) {
		phone = &in.Username
		if u.repo.HasUserPhone(ctx, in.Username) {
			return "", errors.AlreadyExistEmail()
		}
	} else {
		email = &in.Username
		if u.repo.HasUserEmail(ctx, in.Username) {
			return "", errors.AlreadyExistEmail()
		}
	}

	// 获取应用信息
	curApp, err := u.repo.GetApp(ctx, in.App)
	if err != nil {
		return "", errors.NotApp()
	}
	// 判断app是否开启密码注册
	var curChannel *channel.Channel
	for _, item := range curApp.Channels {
		if item.Platform == consts.PasswordChannel {
			curChannel = item
		}
	}
	if curChannel == nil {
		return "", errors.NotAppChannel()
	}

	// 判断app是否能注册
	if curApp.AllowRegistry != nil && !*curApp.AllowRegistry {
		return "", errors.NotUserApp()
	}

	// 注册
	user := &User{
		Email:    email,
		Phone:    phone,
		Status:   proto.Bool(true),
		From:     curApp.Keyword,
		FromDesc: curApp.Name,
		UserApps: []*UserApp{{
			AppID: curApp.ID,
		}},
	}
	if _, err := u.repo.Add(ctx, user); err != nil {
		return "", errors.Register()
	}

	// 生成token
	return u.GenToken(ctx, user, curApp, curChannel)
}

// RefreshToken 刷新用户token
func (u *UseCase) RefreshToken(ctx kratosx.Context) (string, error) {
	token, err := ctx.JWT().Renewal(ctx)
	if err != nil {
		return "", errors.RefreshTokenFormat(err.Error())
	}
	return token, nil
}
