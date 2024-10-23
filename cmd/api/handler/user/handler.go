package user

import "goportfolio/internal/ports/service"

type Handler struct {
	PlayerService service.UserService
}
