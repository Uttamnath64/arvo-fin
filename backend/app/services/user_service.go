package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type UserService struct {
	container *storage.Container
	repo      *repository.User
}

func NewUserService(container *storage.Container) *UserService {
	return &UserService{
		container: container,
		repo:      repository.NewUser(container),
	}
}
