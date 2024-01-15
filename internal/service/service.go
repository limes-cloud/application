package service

import (
	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/logic"
)

type Service struct {
	v1.UnimplementedServiceServer
	user *logic.User
}

func New(conf *config.Config) *Service {
	return &Service{
		user: logic.NewLogic(conf),
	}
}
