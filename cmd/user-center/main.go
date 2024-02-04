package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/configure/client"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	_ "go.uber.org/automaxprocs"

	v1 "github.com/limes-cloud/user-center/api/v1"
	systemConfig "github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/initiator"
	"github.com/limes-cloud/user-center/internal/service"
	_ "github.com/limes-cloud/user-center/pkg/field"
)

func main() {
	app := kratosx.New(
		kratosx.Config(client.NewFromEnv()),
		kratosx.RegistrarServer(RegisterServer),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	conf := &systemConfig.Config{}
	if err := c.Value("business").Scan(conf); err != nil {
		panic("business config format error:" + err.Error())
	}
	c.Watch("business", func(value config.Value) {
		if err := value.Scan(conf); err != nil {
			log.Error("business 配置变更失败")
		}
	})

	srv := service.New(conf)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)

	// 初始化逻辑
	ior := initiator.New(conf)
	if err := ior.Run(); err != nil {
		panic("initiator error:" + err.Error())
	}
}
