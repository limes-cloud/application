package service

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/forgoer/openssl"
	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/library/db/gormtranserror"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"

	"github.com/limes-cloud/application/api/application/auth"
	"github.com/limes-cloud/application/api/application/errors"
	"github.com/limes-cloud/application/internal/conf"
	"github.com/limes-cloud/application/internal/domain/entity"
	"github.com/limes-cloud/application/internal/domain/repository"
	"github.com/limes-cloud/application/internal/pkg/authorizer"
	"github.com/limes-cloud/application/internal/pkg/md"
	"github.com/limes-cloud/application/internal/types"
)

const (
	LOGIN_IMAGE_CAPTCHA    = "loginImage"
	BIND_IMAGE_CAPTCHA     = "bindImage"
	REGISTER_IMAGE_CAPTCHA = "registerImage"

	LOGIN_EMAIL_CAPTCHA    = "loginEmail"
	BIND_EMAIL_CAPTCHA     = "bindEmail"
	REGISTER_EMAIL_CAPTCHA = "registerEmail"
)

const (
	PASSWORD_CERT = "password"
)

type Auth struct {
	conf       *conf.Config
	repo       repository.Auth
	user       repository.User
	app        repository.App
	channel    repository.Channel
	permission repository.Permission
}

func NewAuth(
	conf *conf.Config,
	repo repository.Auth,
	user repository.User,
	app repository.App,
	channel repository.Channel,
	permission repository.Permission,
) *Auth {
	return &Auth{
		conf:       conf,
		repo:       repo,
		user:       user,
		app:        app,
		channel:    channel,
		permission: permission,
	}
}

// Auth 解析token数据
func (u *Auth) Auth(ctx kratosx.Context) (*auth.Auth, error) {
	data, err := md.Get(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ListAuth 获取应用授权信息列表
func (u *Auth) ListAuth(ctx kratosx.Context, req *types.ListAuthRequest) ([]*entity.Auth, uint32, error) {
	list, total, err := u.repo.ListAuth(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateAuth 创建应用授权信息
func (u *Auth) CreateAuth(ctx kratosx.Context, auth *entity.Auth) (uint32, error) {
	auth.Status = proto.Bool(true)
	id, err := u.repo.CreateAuth(ctx, auth)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateAuthStatus 更新应用授权信息状态
func (u *Auth) UpdateAuthStatus(ctx kratosx.Context, req *types.UpdateAuthStatusRequest) error {
	if err := u.repo.UpdateAuthStatus(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteAuth 删除应用授权信息
func (u *Auth) DeleteAuth(ctx kratosx.Context, userId uint32, appId uint32) error {
	if err := u.repo.DeleteAuth(ctx, userId, appId); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// GenAuthCaptcha 生成验证码
func (u *Auth) GenAuthCaptcha(ctx kratosx.Context, req *types.GenAuthCaptchaRequest) (*types.GenAuthCaptchaResponse, error) {
	switch req.Type {
	case LOGIN_IMAGE_CAPTCHA, BIND_IMAGE_CAPTCHA, REGISTER_IMAGE_CAPTCHA:
		return u.captcha(ctx, req.Type)
	case LOGIN_EMAIL_CAPTCHA, BIND_EMAIL_CAPTCHA:
		if req.Email == "" {
			return nil, errors.ParamsError()
		}
		if !u.user.HasUserByEmail(ctx, req.Email) {
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
func (u *Auth) ListOAuth(ctx kratosx.Context, req *types.ListOAuthRequest) ([]*entity.OAuth, uint32, error) {
	list, total, err := u.repo.ListOAuth(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// DeleteOAuth 删除应用授权信息
func (u *Auth) DeleteOAuth(ctx kratosx.Context, userId uint32, channelId uint32) error {
	if err := u.repo.DeleteOAuth(ctx, userId, channelId); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

func (u *Auth) GenToken(ctx kratosx.Context, app *entity.App, user *entity.User) (string, error) {
	// 如果应用不允许注册，则判断是否具有应用权限
	ua, err := u.repo.GetAuthByUA(ctx, user.Id, app.Id)
	if err != nil {
		if !gormtranserror.Is(err, gorm.ErrRecordNotFound) {
			ctx.Logger().Warnw("msg", "get auth by ua error", "err", err.Error())
			return "", errors.SystemError()
		}

		// 不具有权限，则判断是否允许注册
		if app.AllowRegistry == nil || !*app.AllowRegistry {
			return "", errors.NotAppScopeError()
		}
	}

	// 判断用户是否被禁止登陆
	if ua.Status != nil && !*ua.Status {
		return "", errors.UserDisableError(ua.GetDisableDescString())
	}

	// 生成登陆token
	token, err := ctx.JWT().NewToken(md.New(auth.Auth{
		AppKeyword: app.Keyword,
		UserId:     user.Id,
	}))
	if err != nil {
		return "", errors.GenTokenError(err.Error())
	}

	if err = u.repo.UpsertAuth(ctx, &entity.Auth{
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

func (u *Auth) OAuthLogin(ctx kratosx.Context, req *types.OAuthLoginRequest) (*types.OAuthLoginReply, error) {
	app, err := u.app.GetAppByKeyword(ctx, req.App)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	channel, err := u.getAppChannel(ctx, req.App, req.Channel)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	// 获取授权信息
	author := authorizer.New(&authorizer.Config{
		Platform: channel.Keyword,
		Ak:       channel.GetAkString(),
		Sk:       channel.GetSkString(),
		Extra:    channel.GetExtraString(),
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

	// 判断是否绑定三方
	if !u.repo.IsBindOAuth(ctx, channel.Id, ai.AuthId) {
		id, err := u.repo.CreateOAuth(ctx, &entity.OAuth{
			ChannelId: channel.Id,
			Token:     ti.Token,
			ExpiredAt: ti.Expire,
			LoggedAt:  time.Now().Unix(),
			AuthId:    ai.AuthId,
			UnionId:   ai.UnionId,
		})
		if err != nil {
			return nil, errors.OAuthLoginError(err.Error())
		}
		return &types.OAuthLoginReply{
			OAuthUid: &id,
			IsBind:   false,
		}, nil
	}

	// 获取用户信息
	user, err := u.getUserByCA(ctx, channel.Id, ai.AuthId)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	reply := types.OAuthLoginReply{IsBind: true, Expire: proto.Uint32(uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds())))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		token, err := u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.UpdateOAuth(ctx, &entity.OAuth{
			UserId:    &user.Id,
			ChannelId: channel.Id,
			Token:     ti.Token,
			ExpiredAt: ti.Expire,
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
func (u *Auth) EmailLogin(ctx kratosx.Context, req *types.EmailLoginRequest) (*types.TokenInfo, error) {
	app, err := u.app.GetAppByKeyword(ctx, req.App)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	channel, err := u.getAppChannel(ctx, req.App, EmailChannel)
	if err != nil {
		return nil, errors.OAuthLoginError(err.Error())
	}

	if err := u.verifyCaptcha(ctx, LOGIN_EMAIL_CAPTCHA, req.CaptchaId, req.Captcha, req.Email); err != nil {
		return nil, err
	}

	// 获取用户信息
	user, err := u.user.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.NotUserError()
	}

	reply := types.TokenInfo{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.UpdateOAuth(ctx, &entity.OAuth{
			UserId:    &user.Id,
			ChannelId: channel.Id,
			AuthId:    req.Email,
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
func (u *Auth) PasswordLogin(ctx kratosx.Context, req *types.PasswordLoginRequest) (*types.PasswordLoginReply, error) {
	app, err := u.app.GetAppByKeyword(ctx, req.App)
	if err != nil {
		return nil, errors.LoginError(err.Error())
	}

	channel, err := u.getAppChannel(ctx, req.App, PasswordChannel)
	if err != nil {
		return nil, errors.LoginError(err.Error())
	}

	if err := u.verifyCaptcha(ctx, LOGIN_IMAGE_CAPTCHA, req.CaptchaId, req.Captcha, ""); err != nil {
		return nil, err
	}

	// 获取用户信息
	user, err := u.user.GetUserByUsername(ctx, req.Username)
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
	var pw types.Password
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

	reply := types.PasswordLoginReply{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.UpdateOAuth(ctx, &entity.OAuth{
			UserId:    &user.Id,
			ChannelId: channel.Id,
			AuthId:    req.Username,
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
func (u *Auth) EmailRegister(ctx kratosx.Context, req *types.EmailRegisterRequest) (*types.TokenInfo, error) {
	app, err := u.app.GetAppByKeyword(ctx, req.App)
	if err != nil {
		return nil, errors.RegisterError(err.Error())
	}
	if app.AllowRegistry == nil || !*app.AllowRegistry {
		return nil, errors.DisableRegisterError()
	}

	channel, err := u.getAppChannel(ctx, req.App, EmailChannel)
	if err != nil {
		return nil, errors.RegisterError(err.Error())
	}

	if err := u.verifyCaptcha(ctx, REGISTER_EMAIL_CAPTCHA, req.CaptchaId, req.Captcha, req.Email); err != nil {
		return nil, err
	}

	// 获取用户信息
	if u.user.HasUserByEmail(ctx, req.Email) {
		return nil, errors.AlreadyExistEmailError()
	}

	// 创建账户
	md5 := crypto.MD5([]byte(uuid.NewString()))
	if err = u.register(ctx, app.Id, &entity.User{
		Email:    proto.String(req.Email),
		From:     app.Keyword,
		FromDesc: app.Name,
		NickName: u.conf.DefaultNickName + md5[:8],
		Status:   proto.Bool(true),
	}); err != nil {
		return nil, errors.RegisterError(err.Error())
	}

	user, err := u.user.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.RegisterError(err.Error())
	}

	reply := types.TokenInfo{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.UpdateOAuth(ctx, &entity.OAuth{
			UserId:    &user.Id,
			ChannelId: channel.Id,
			ExpiredAt: int64(reply.Expire),
			LoggedAt:  time.Now().Unix(),
		}); err != nil {
			return errors.DatabaseError(err.Error())
		}

		if req.OAuthUid != nil {
			if err := u.repo.BindOAuthByUid(ctx, user.Id, *req.OAuthUid); err != nil {
				if gormtranserror.Is(err, gorm.ErrDuplicatedKey) {
					return errors.AlreadyBindError()
				}
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
func (u *Auth) PasswordRegister(ctx kratosx.Context, req *types.PasswordRegisterRequest) (*types.TokenInfo, error) {
	app, err := u.app.GetAppByKeyword(ctx, req.App)
	if err != nil {
		return nil, errors.RegisterError(err.Error())
	}
	if app.AllowRegistry == nil || !*app.AllowRegistry {
		return nil, errors.DisableRegisterError()
	}

	channel, err := u.getAppChannel(ctx, req.App, PasswordChannel)
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
	var pw types.Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return nil, errors.PasswordFormatError()
	}

	// 判断当前时间戳是否过期,超过则拒绝
	if time.Now().UnixMilli()-pw.Time > u.conf.DefaultPasswordExpire.Milliseconds() {
		return nil, errors.PasswordExpireError()
	}

	// 判断用户是否已经存在
	if u.user.HasUserByUsername(ctx, req.Username) {
		return nil, errors.AlreadyExistUsernameError()
	}

	// 创建账户
	md5 := crypto.MD5([]byte(uuid.NewString()))
	if err = u.register(ctx, app.Id, &entity.User{
		Username: proto.String(req.Username),
		Password: proto.String(crypto.EncodePwd(pw.Password)),
		From:     app.Keyword,
		FromDesc: app.Name,
		NickName: u.conf.DefaultNickName + md5[:8],
		Status:   proto.Bool(true),
	}); err != nil {
		return nil, errors.RegisterError(err.Error())
	}

	user, err := u.user.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.RegisterError(err.Error())
	}

	reply := types.TokenInfo{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.UpdateOAuth(ctx, &entity.OAuth{
			UserId:    &user.Id,
			ChannelId: channel.Id,
			ExpiredAt: int64(reply.Expire),
			LoggedAt:  time.Now().Unix(),
		}); err != nil {
			return errors.DatabaseError(err.Error())
		}

		if req.OAuthUid != nil {
			if err := u.repo.BindOAuthByUid(ctx, user.Id, *req.OAuthUid); err != nil {
				if gormtranserror.Is(err, gorm.ErrDuplicatedKey) {
					return errors.AlreadyBindError()
				}
				return errors.DatabaseError(err.Error())
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &reply, nil
}

// EmailBind 邮箱绑定
func (u *Auth) EmailBind(ctx kratosx.Context, req *types.EmailBindRequest) (*types.TokenInfo, error) {
	app, err := u.app.GetAppByKeyword(ctx, req.App)
	if err != nil {
		return nil, errors.BindError(err.Error())
	}

	if _, err = u.getAppChannel(ctx, req.App, EmailChannel); err != nil {
		return nil, errors.BindError(err.Error())
	}

	if err := u.verifyCaptcha(ctx, BIND_EMAIL_CAPTCHA, req.CaptchaId, req.Captcha, req.Email); err != nil {
		return nil, err
	}

	user, err := u.user.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.NotUserError()
	}

	reply := types.TokenInfo{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.BindOAuthByUid(ctx, user.Id, req.OAuthUid); err != nil {
			if gormtranserror.Is(err, gorm.ErrDuplicatedKey) {
				return errors.AlreadyBindError()
			}
			return errors.DatabaseError(err.Error())
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &reply, nil
}

// PasswordBind 密码注册
func (u *Auth) PasswordBind(ctx kratosx.Context, req *types.PasswordBindRequest) (*types.TokenInfo, error) {
	app, err := u.app.GetAppByKeyword(ctx, req.App)
	if err != nil {
		return nil, errors.BindError(err.Error())
	}

	if _, err := u.getAppChannel(ctx, req.App, PasswordChannel); err != nil {
		return nil, errors.BindError(err.Error())
	}

	// if err := u.verifyCaptcha(ctx, BIND_IMAGE_CAPTCHA, req.CaptchaId, req.Captcha, ""); err != nil {
	//	return nil, err
	// }

	// 获取用户信息
	user, err := u.user.GetUserByUsername(ctx, req.Username)
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
	var pw types.Password
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

	reply := types.TokenInfo{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		reply.Token, err = u.GenToken(ctx, app, user)
		if err != nil {
			return err
		}

		if err := u.repo.BindOAuthByUid(ctx, user.Id, req.OAuthUid); err != nil {
			if gormtranserror.Is(err, gorm.ErrDuplicatedKey) {
				return errors.AlreadyBindError()
			}
			return errors.DatabaseError(err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &reply, nil
}

// RefreshToken 刷新用户token
func (u *Auth) RefreshToken(ctx kratosx.Context) (*types.TokenInfo, error) {
	var err error
	reply := types.TokenInfo{Expire: uint32(time.Now().Unix() + int64(ctx.Config().App().JWT.Expire.Seconds()))}
	reply.Token, err = ctx.JWT().Renewal(ctx)
	if err != nil {
		return nil, errors.RefreshTokenError(err.Error())
	}
	return &reply, nil
}

// Logout 用户退出登陆
func (u *Auth) Logout(ctx kratosx.Context) error {
	token := ctx.Token()
	if token == "" {
		return nil
	}

	if ctx.JWT().IsBlacklist(token) {
		return errors.RefreshTokenError()
	}

	ctx.JWT().AddBlacklist(token)
	return nil
}

// verifyCaptcha 验证验证码
func (u *Auth) verifyCaptcha(ctx kratosx.Context, tp string, id string, captcha, sender string) error {
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

func (u *Auth) captcha(ctx kratosx.Context, tp string) (*types.GenAuthCaptchaResponse, error) {
	res, err := ctx.Captcha().Image(tp, ctx.ClientIP())
	if err != nil {
		return nil, errors.GenCaptchaError(err.Error())
	}
	return &types.GenAuthCaptchaResponse{
		Id:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}

func (u *Auth) email(ctx kratosx.Context, tp, email string) (*types.GenAuthCaptchaResponse, error) {
	res, err := ctx.Captcha().Email(tp, ctx.ClientIP(), email)
	if err != nil {
		return nil, errors.GenCaptchaError(err.Error())
	}
	return &types.GenAuthCaptchaResponse{
		Id:     res.ID(),
		Base64: res.Base64String(),
		Expire: int(res.Expire().Seconds()),
	}, nil
}

func (u *Auth) getAppChannel(ctx kratosx.Context, ak, ck string) (*entity.Channel, error) {
	// 获取应用
	app, err := u.app.GetAppByKeyword(ctx, ak)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.SystemError()
	}
	if app.Status != nil && !*app.Status {
		return nil, errors.AppMaintenanceError()
	}

	// 获取channel
	var channel *entity.Channel
	for _, ch := range app.Channels {
		if ch.Keyword == ck {
			channel = ch
		}
	}
	if channel == nil {
		return nil, errors.AppNotBindChannelError()
	}

	return channel, nil
}

func (u *Auth) getUserByCA(ctx kratosx.Context, cid uint32, aid string) (*entity.User, error) {
	oauth, err := u.repo.GetOAuthByCA(ctx, cid, aid)
	if err != nil {
		ctx.Logger().Warnw("msg", "get oauth by ca error", "err", err.Error())
		return nil, errors.SystemError()
	}
	if oauth.UserId == nil {
		return nil, errors.ChannelNotBindUserError()
	}

	user, err := u.user.GetUser(ctx, *oauth.UserId)
	if err != nil {
		ctx.Logger().Warnw("msg", "get user  error", "err", err.Error())
		return nil, errors.NotUserError()
	}

	if user.Status != nil && !*user.Status {
		return nil, errors.UserDisableError(user.DisableDesc)
	}
	return user, nil
}

func (u *Auth) register(ctx kratosx.Context, aid uint32, user *entity.User) error {
	return ctx.Transaction(func(ctx kratosx.Context) error {
		uid, err := u.user.CreateUser(ctx, user)
		if err != nil {
			ctx.Logger().Warnw("msg", "register user error", "err", err.Error())
			return errors.RegisterError()
		}

		if _, err := u.repo.CreateAuth(ctx, &entity.Auth{
			UserId: uid,
			AppId:  aid,
			Status: proto.Bool(true),
		}); err != nil {
			ctx.Logger().Warnw("msg", "create auth error", "err", err.Error())
			return errors.RegisterError()
		}
		return nil
	})
}
