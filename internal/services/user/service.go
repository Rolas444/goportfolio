package user

import "goportfolio/internal/ports"

type Service struct {
	Repo ports.UserRepository
}
