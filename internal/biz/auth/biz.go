package auth

import (
	"encoding/base64"
	"time"

	"github.com/forgoer/openssl"
	"github.com/google/uuid"
	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/usercenter/api/usercenter/auth"
	"github.com/limes-cloud/usercenter/api/usercenter/errors"
	"github.com/limes-cloud/usercenter/internal/conf"
	"github.com/limes-cloud/usercenter/internal/consts"
	"github.com/limes-cloud/usercenter/internal/pkg/authorizer"
	"github.com/limes-cloud/usercenter/internal/pkg/md"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// Auth 解析token数据
func (u *UseCase) Auth(ctx kratosx.Context) (*auth.Auth, error) {
	data, err := md.Get(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ListAuth 获取应用授权信息列表
func (u *UseCase) ListAuth(ctx kratosx.Context, req *ListAuthRequest) ([]*Auth, uint32, error) {
	list, total, err := u.repo.ListAuth(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateAuth 创建应用授权信息
func (u *UseCase) CreateAuth(ctx kratosx.Context, req *Auth) (uint32, error) {
	req.Status = proto.Bool(true)
	id, err := u.repo.CreateAuth(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateAuthStatus 更新应用授权信息状态
func (u *UseCase) UpdateAuthStatus(ctx kratosx.Context, req *UpdateAuthStatusRequest) error {
	if err := u.repo.UpdateAuthStatus(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteAuth 删除应用授权信息
func (u *UseCase) DeleteAuth(ctx kratosx.Context, userId uint32, appId uint32) error {
	if err := u.repo.DeleteAuth(ctx, userId, appId); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// GenAuthCaptcha 生成验证码
func (u *UseCase) GenAuthCaptcha(ctx kratosx.Context, req *GenAuthCaptchaRequest) (*GenAuthCaptchaResponse, error) {
	switch req.Type {
	case LOGIN_IMAGE_CAPTCHA, BIND_IMAGE_CAPTCHA, REGISTER_IMAGE_CAPTCHA:
		return u.captcha(ctx, req.Type)
	case LOGIN_EMAIL_CAPTCHA, BIND_EMAIL_CAPTCHA:
		if req.Email == "" {
			return nil, errors.ParamsError()
		}
		if !u.repo.HasUserByEmail(ctx, req.Email) {
			return nil, errors.NotExistEmailError()
		}
		return u.email(ctx, req.Type, req.Email)
	case REGISTER_EMAIL_CAPTCHA:
		if req.Email == "" {
			return nil, errors.ParamsError()
		}
		return u.email(ctx, req.Type, req.Email)
	default:
		return nil, errors.GenCaptchaTypeError()
	}
}

// ListOAuth 获取应用授权信息列表
func (u *UseCase) ListOAuth(ctx kratosx.Context, req *ListOAuthRequest) ([]*OAuth, uint32, error) {
	list, total, err := u.repo.ListOAuth(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// DeleteOAuth 删除应用授权信息
func (u *UseCase) DeleteOAuth(ctx kratosx.Context, userId uint32, channelId uint32) error {
	if err := u.repo.DeleteOAuth(ctx, userId, channelId); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

func (u *UseCase) GenToken(ctx kratosx.Context, app *App, user *User) (string, error) {
	// 如果应用不允许注册，则判断是否具有应用权限
	if err := u.repo.HasAppScope(ctx, app.Id, user.Id); err != nil {
		return "", errors.NotAppScopeError(err.Error())
	}

	// 生成登陆token
	token, err := ctx.JWT().NewToken(md.New(auth.Auth{
		AppKeyword: app.Keyword,
		UserId:     user.Id,
	}))
	if err != nil {
		return "", errors.GenTokenError(err.Error())
	}

	if err = u.repo.UpdateAuth(ctx, &Auth{
		UserId:    user.Id,
		AppId:     app.Id,
		LoggedAt:  time.Now().Unix(),
		Token:     token,
		ExpiredAt: time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()),
	}); err != nil {
		return "", errors.DatabaseError(err.Error())
	}

	return token, nil
}

func (u *UseCase) OAuthLogin(ctx kratosx.Context, req *OAuthLoginRequest) (*OAuthLoginReply, error) {
	app, err := u.repo.GetApp(ctx, req.App)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	channel, err := u.repo.GetAppChannel(ctx, req.App, req.Channel)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	// 获取授权信息
	author := authorizer.New(&authorizer.Config{
		Platform: channel.Keyword,
		Ak:       channel.Ak,
		Sk:       channel.Sk,
		Extra:    channel.Extra,
		Code:     req.Code,
	})

	ti, err := author.GetToken(ctx)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}
	ai, err := author.GetAuthInfo(ctx, ti.Token)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	// 判断是否绑定用户
	if !u.repo.IsBindUser(ctx, channel.Id, ai.AuthId) {
		id, err := u.repo.CreateOAuth(ctx, &OAuth{
			ChannelId: channel.Id,
			Token:     ti.Token,
			ExpiredAt: int64(ti.Expire.Seconds()),
			LoggedAt:  time.Now().Unix(),
			AuthId:    ai.AuthId,
			UnionId:   ai.UnionId,
		})
		if err != nil {
			return nil, errors.OAuthLoginError(err.Error())
		}
		return &OAuthLoginReply{
			OAuthUid: &id,
			IsBind:   false,
		}, nil
	}

	// 获取用户信息
	user, err := u.repo.GetUserByCA(ctx, channel.Id, ai.AuthId)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	reply := OAuthLoginReply{IsBind: true, Expire: proto.Uint32(uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds())))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		token, err := u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.UpdateOAuth(ctx, &OAuth{
			UserId:    &user.Id,
			ChannelId: channel.Id,
			Token:     ti.Token,
			ExpiredAt: int64(ti.Expire.Seconds()),
			LoggedAt:  time.Now().Unix(),
			AuthId:    ai.AuthId,
			UnionId:   ai.UnionId,
		}); err != nil {
			return errors.DatabaseError(err.Error())
		}
		reply.Token = &token
		return nil
	}); err != nil {
		return nil, err
	}

	return &reply, nil
}

// EmailLogin 邮箱登陆
func (u *UseCase) EmailLogin(ctx kratosx.Context, req *EmailLoginRequest) (*EmailLoginReply, error) {
	app, err := u.repo.GetApp(ctx, req.App)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	channel, err := u.repo.GetAppChannel(ctx, req.App, consts.EmailChannel)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	if err := u.verifyCaptcha(ctx, LOGIN_EMAIL_CAPTCHA, req.CaptchaId, req.Captcha, req.Email); err != nil {
		return nil, err
	}

	// 获取用户信息
	user, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.NotUserError()
	}

	reply := EmailLoginReply{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.UpdateOAuth(ctx, &OAuth{
			UserId:    &user.Id,
			ChannelId: channel.Id,
			ExpiredAt: int64(reply.Expire),
			LoggedAt:  time.Now().Unix(),
		}); err != nil {
			return errors.DatabaseError(err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &reply, nil
}

// PasswordLogin 密码登陆
func (u *UseCase) PasswordLogin(ctx kratosx.Context, req *PasswordLoginRequest) (*PasswordLoginReply, error) {
	app, err := u.repo.GetApp(ctx, req.App)
	if err != nil {
		return nil, errors.LoginError(err.Error())
	}

	channel, err := u.repo.GetAppChannel(ctx, req.App, consts.PasswordChannel)
	if err != nil {
		return nil, errors.LoginError(err.Error())
	}

	if err := u.verifyCaptcha(ctx, LOGIN_IMAGE_CAPTCHA, req.CaptchaId, req.Captcha, ""); err != nil {
		return nil, err
	}

	// 获取用户信息
	user, err := u.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.NotUserError()
	}

	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(req.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(PASSWORD_CERT))
	if err != nil {
		return nil, errors.RsaDecodeError(err.Error())
	}

	// 序列化密码
	var pw Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return nil, errors.PasswordFormatError()
	}

	// 判断当前时间戳是否过期,超过3s则拒绝
	if time.Now().UnixMilli()-pw.Time > u.conf.DefaultPasswordExpire.Milliseconds() {
		return nil, errors.PasswordExpireError()
	}

	// 对比用户密码，错误则拒绝登陆
	pwd := ""
	if user.Password != nil {
		pwd = *user.Password
	}
	if !crypto.CompareHashPwd(pwd, pw.Password) {
		return nil, errors.PasswordError()
	}

	reply := PasswordLoginReply{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.UpdateOAuth(ctx, &OAuth{
			UserId:    &user.Id,
			ChannelId: channel.Id,
			ExpiredAt: int64(reply.Expire),
			LoggedAt:  time.Now().Unix(),
		}); err != nil {
			return errors.DatabaseError(err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &reply, nil
}

// EmailRegister 邮箱注册
func (u *UseCase) EmailRegister(ctx kratosx.Context, req *EmailRegisterRequest) (*EmailRegisterReply, error) {
	app, err := u.repo.GetApp(ctx, req.App)
	if err != nil {
		return nil, errors.RegisterError(err.Error())
	}
	if app.AllowRegister == nil || !*app.AllowRegister {
		return nil, errors.DisableRegisterError()
	}

	channel, err := u.repo.GetAppChannel(ctx, req.App, consts.EmailChannel)
	if err != nil {
		return nil, errors.RegisterError(err.Error())
	}

	if err := u.verifyCaptcha(ctx, REGISTER_EMAIL_CAPTCHA, req.CaptchaId, req.Captcha, req.Email); err != nil {
		return nil, err
	}

	// 获取用户信息
	if u.repo.HasUserByEmail(ctx, req.Email) {
		return nil, errors.AlreadyExistEmailError()
	}

	// 创建账户
	md5 := crypto.MD5([]byte(uuid.NewString()))
	if err = u.repo.Register(ctx, &User{
		Email:    proto.String(req.Email),
		From:     app.Keyword,
		FromDesc: app.Name,
		AppIds:   []uint32{app.Id},
		Nickname: u.conf.DefaultNickName + md5[:8],
		Status:   proto.Bool(true),
	}); err != nil {
		return nil, errors.RegisterError(err.Error())
	}

	user, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.RegisterError(err.Error())
	}

	reply := EmailRegisterReply{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.UpdateOAuth(ctx, &OAuth{
			UserId:    &user.Id,
			ChannelId: channel.Id,
			ExpiredAt: int64(reply.Expire),
			LoggedAt:  time.Now().Unix(),
		}); err != nil {
			return errors.DatabaseError(err.Error())
		}

		if req.OAuthUid != nil {
			if err := u.repo.BindOAuthByUid(ctx, user.Id, *req.OAuthUid); err != nil {
				return errors.DatabaseError(err.Error())
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &reply, nil
}

// PasswordRegister 密码注册
func (u *UseCase) PasswordRegister(ctx kratosx.Context, req *PasswordRegisterRequest) (*PasswordRegisterReply, error) {
	app, err := u.repo.GetApp(ctx, req.App)
	if err != nil {
		return nil, errors.RegisterError(err.Error())
	}
	if app.AllowRegister == nil || !*app.AllowRegister {
		return nil, errors.DisableRegisterError()
	}

	channel, err := u.repo.GetAppChannel(ctx, req.App, consts.PasswordChannel)
	if err != nil {
		return nil, errors.RegisterError(err.Error())
	}

	if err := u.verifyCaptcha(ctx, REGISTER_IMAGE_CAPTCHA, req.CaptchaId, req.Captcha, ""); err != nil {
		return nil, err
	}

	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(req.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(PASSWORD_CERT))
	if err != nil {
		return nil, errors.RsaDecodeError(err.Error())
	}

	// 序列化密码
	var pw Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return nil, errors.PasswordFormatError()
	}

	// 判断当前时间戳是否过期,超过则拒绝
	if time.Now().UnixMilli()-pw.Time > u.conf.DefaultPasswordExpire.Milliseconds() {
		return nil, errors.PasswordExpireError()
	}

	// 判断用户是否已经存在
	if u.repo.HasUserByUsername(ctx, req.Username) {
		return nil, errors.AlreadyExistUsernameError()
	}

	// 创建账户
	md5 := crypto.MD5([]byte(uuid.NewString()))
	if err = u.repo.Register(ctx, &User{
		Username: proto.String(req.Username),
		Password: proto.String(crypto.EncodePwd(pw.Password)),
		From:     app.Keyword,
		FromDesc: app.Name,
		AppIds:   []uint32{app.Id},
		Nickname: u.conf.DefaultNickName + md5[:8],
		Status:   proto.Bool(true),
	}); err != nil {
		return nil, errors.RegisterError(err.Error())
	}

	user, err := u.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.RegisterError(err.Error())
	}

	reply := PasswordRegisterReply{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.UpdateOAuth(ctx, &OAuth{
			UserId:    &user.Id,
			ChannelId: channel.Id,
			ExpiredAt: int64(reply.Expire),
			LoggedAt:  time.Now().Unix(),
		}); err != nil {
			return errors.DatabaseError(err.Error())
		}

		if req.OAuthUid != nil {
			if err := u.repo.BindOAuthByUid(ctx, user.Id, *req.OAuthUid); err != nil {
				return errors.DatabaseError(err.Error())
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &reply, nil
}

// EmailBind 邮箱注册
func (u *UseCase) EmailBind(ctx kratosx.Context, req *EmailBindRequest) (*EmailBindReply, error) {
	app, err := u.repo.GetApp(ctx, req.App)
	if err != nil {
		return nil, errors.BindError(err.Error())
	}

	if _, err = u.repo.GetAppChannel(ctx, req.App, consts.EmailChannel); err != nil {
		return nil, errors.BindError(err.Error())
	}

	if err := u.verifyCaptcha(ctx, BIND_EMAIL_CAPTCHA, req.CaptchaId, req.Captcha, req.Email); err != nil {
		return nil, err
	}

	user, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.NotUserError()
	}

	reply := EmailBindReply{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.BindOAuthByUid(ctx, user.Id, req.OAuthUid); err != nil {
			return errors.DatabaseError(err.Error())
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &reply, nil
}

// PasswordBind 密码注册
func (u *UseCase) PasswordBind(ctx kratosx.Context, req *PasswordBindRequest) (*PasswordBindReply, error) {
	app, err := u.repo.GetApp(ctx, req.App)
	if err != nil {
		return nil, errors.BindError(err.Error())
	}

	if _, err := u.repo.GetAppChannel(ctx, req.App, consts.PasswordChannel); err != nil {
		return nil, errors.BindError(err.Error())
	}

	if err := u.verifyCaptcha(ctx, BIND_IMAGE_CAPTCHA, req.CaptchaId, req.Captcha, ""); err != nil {
		return nil, err
	}

	// 获取用户信息
	user, err := u.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.BindError(err.Error())
	}

	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(req.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(PASSWORD_CERT))
	if err != nil {
		return nil, errors.RsaDecodeError(err.Error())
	}

	// 序列化密码
	var pw Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return nil, errors.PasswordFormatError()
	}

	// 判断当前时间戳是否过期,超过则拒绝
	if time.Now().UnixMilli()-pw.Time > u.conf.DefaultPasswordExpire.Milliseconds() {
		return nil, errors.PasswordExpireError()
	}

	// 对比用户密码，错误则拒绝登陆
	pwd := ""
	if user.Password != nil {
		pwd = *user.Password
	}
	if !crypto.CompareHashPwd(pwd, pw.Password) {
		return nil, errors.PasswordError()
	}

	reply := PasswordBindReply{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.BindOAuthByUid(ctx, user.Id, req.OAuthUid); err != nil {
			return errors.DatabaseError(err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &reply, nil
}

// verifyCaptcha 验证验证码
func (u *UseCase) verifyCaptcha(ctx kratosx.Context, tp string, id string, captcha, sender string) error {
	switch tp {
	case LOGIN_IMAGE_CAPTCHA, BIND_IMAGE_CAPTCHA, REGISTER_IMAGE_CAPTCHA:
		if err := ctx.Captcha().VerifyImage(tp, ctx.ClientIP(), id, captcha); err != nil {
			return errors.VerifyCaptchaError()
		}
	case LOGIN_EMAIL_CAPTCHA, BIND_EMAIL_CAPTCHA, REGISTER_EMAIL_CAPTCHA:
		if err := ctx.Captcha().VerifyEmail(tp, ctx.ClientIP(), id, captcha, sender); err != nil {
			return errors.VerifyCaptchaError()
		}
	default:
		return errors.GenCaptchaTypeError()
	}
	return nil
}

func (u *UseCase) captcha(ctx kratosx.Context, tp string) (*GenAuthCaptchaResponse, error) {
	res, err := ctx.Captcha().Image(tp, ctx.ClientIP())
	if err != nil {
		return nil, errors.GenCaptchaError(err.Error())
	}
	return &GenAuthCaptchaResponse{
		Id:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}

func (u *UseCase) email(ctx kratosx.Context, tp, email string) (*GenAuthCaptchaResponse, error) {
	res, err := ctx.Captcha().Email(tp, ctx.ClientIP(), email)
	if err != nil {
		return nil, errors.GenCaptchaError(err.Error())
	}
	return &GenAuthCaptchaResponse{
		Id:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}
