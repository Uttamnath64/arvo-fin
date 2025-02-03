package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type UserService struct {
	container *storage.Container
	userRepo  *repository.UserRepository
}

func NewUserService(container *storage.Container) *UserService {
	return &UserService{
		container: container,
		userRepo:  repository.NewUserRepository(container),
	}
}
