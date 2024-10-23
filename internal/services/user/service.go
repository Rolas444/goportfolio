package user

import (
	"goportfolio/internal/ports/service"
)

type Service struct {
}

func NewUserService() service.UserService {
	return &Service{}
}
