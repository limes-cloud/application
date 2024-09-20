package app

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	pb "github.com/limes-cloud/application/api/application/auth/v1"
	"github.com/limes-cloud/application/internal/conf"
	"github.com/limes-cloud/application/internal/domain/service"
	"github.com/limes-cloud/application/internal/infra/dbs"
	"github.com/limes-cloud/application/internal/infra/rpc"
)

type Auth struct {
	pb.UnimplementedAuthServer
	srv *service.Auth
}

func NewAuth(conf *conf.Config) *Auth {
	return &Auth{
		srv: service.NewAuth(
			conf,
			dbs.NewAuth(),
			dbs.NewUser(),
			dbs.NewApp(),
			dbs.NewChannel(),
			rpc.NewPermission(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewAuth(c)
		pb.RegisterAuthHTTPServer(hs, srv)
		pb.RegisterAuthServer(gs, srv)
	})
}
