package user

import "goportfolio/internal/ports"

type Handler struct {
	PlayerService ports.UserService
}
