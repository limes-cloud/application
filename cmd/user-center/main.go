package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/configure/api/client"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	_ "go.uber.org/automaxprocs"

	internalconfig "github.com/limes-cloud/user-center/internal/config"
	"github.com/limes-cloud/user-center/internal/initiator"
	"github.com/limes-cloud/user-center/internal/service"
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
	// 初始化并监听配置变更
	conf := &internalconfig.Config{}
	c.ScanWatch("business", func(value config.Value) {
		if err := value.Scan(conf); err != nil {
			panic("business config format error:" + err.Error())
		}
	})

	// 初始化逻辑
	ior := initiator.New(conf)
	if err := ior.Run(); err != nil {
		panic("initiator error:" + err.Error())
	}

	// 注册服务
	service.New(conf, hs, gs)
}
