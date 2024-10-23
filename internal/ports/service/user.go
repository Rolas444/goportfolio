package service

import "goportfolio/internal/domain"

type UserService interface {
	Create(user domain.User) (domain.User, error)
	Update(user domain.User) (string, error)
	getOne(id string) (domain.User, error)
	getMany() ([]domain.User, error)
	Delete(id string) (string, error)
}

type UserRepository interface {
	Insert(user domain.User) (string, error)
	Update(user domain.User) (string, error)
	getOne(id string) (domain.User, error)
	getMany() ([]domain.User, error)
	Delete(id string) (string, error)
}
