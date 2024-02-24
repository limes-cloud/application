package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	agreementpb "github.com/limes-cloud/user-center/api/agreement/v1"
	apppb "github.com/limes-cloud/user-center/api/app/v1"
	channelpb "github.com/limes-cloud/user-center/api/channel/v1"
	fieldpb "github.com/limes-cloud/user-center/api/field/v1"
	userpb "github.com/limes-cloud/user-center/api/user/v1"
	"github.com/limes-cloud/user-center/internal/config"
)

func New(c *config.Config, hs *http.Server, gs *grpc.Server) {
	agreementSrv := NewAgreement(c)
	agreementpb.RegisterServiceHTTPServer(hs, agreementSrv)
	agreementpb.RegisterServiceServer(gs, agreementSrv)

	channelSrv := NewChannel(c)
	channelpb.RegisterServiceHTTPServer(hs, channelSrv)
	channelpb.RegisterServiceServer(gs, channelSrv)

	appSrv := NewApp(c)
	apppb.RegisterServiceHTTPServer(hs, appSrv)
	apppb.RegisterServiceServer(gs, appSrv)

	fieldSrv := NewField(c)
	fieldpb.RegisterServiceHTTPServer(hs, fieldSrv)
	fieldpb.RegisterServiceServer(gs, fieldSrv)

	userSrv := NewUser(c)
	userpb.RegisterServiceHTTPServer(hs, userSrv)
	userpb.RegisterServiceServer(gs, userSrv)
}
