package ports

import "goportfolio/internal/domain"

type UserService interface {
	Create(user domain.User) (id uint, err error)
}

type UserRepository interface {
}
