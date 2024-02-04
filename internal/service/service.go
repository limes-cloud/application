package service

import (
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/biz"
	"github.com/limes-cloud/user-center/internal/data"
)

type Service struct {
	v1.UnimplementedServiceServer
	conf         *config.Config
	agreement    *biz.AgreementUseCase
	scene        *biz.SceneUseCase
	channel      *biz.ChannelUseCase
	app          *biz.AppUseCase
	appInterface *biz.AppInterfaceUseCase
	extraField   *biz.ExtraFieldUseCase
	user         *biz.UserUseCase
	userApp      *biz.UserAppUseCase
	userChannel  *biz.UserChannelUseCase
	auth         *biz.AuthUseCase
}

func New(conf *config.Config) *Service {
	return &Service{
		conf:         conf,
		channel:      biz.NewChannelUseCase(conf, data.NewChannelRepo()),
		app:          biz.NewAppUseCase(conf, data.NewAppRepo()),
		appInterface: biz.NewAppInterfaceUseCase(conf, data.NewAppInterfaceRepo()),
		extraField:   biz.NewExtraFieldUseCase(conf, data.NewExtraFieldRepo()),
		user:         biz.NewUserUseCase(conf, data.NewUserRepo()),
		userApp:      biz.NewUserAppUseCase(conf, data.NewUserAppRepo()),
		userChannel:  biz.NewUserChannelUseCase(conf, data.NewUserChannelRepo()),
		agreement:    biz.NewAgreementUseCase(conf, data.NewAgreementRepo()),
		scene:        biz.NewSceneUseCase(conf, data.NewSceneRepo()),
		auth:         biz.NewAuthUseCase(conf, data.NewAuthRepo()),
	}
}
