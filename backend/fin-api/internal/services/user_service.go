package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type User struct {
	container *storage.Container
	userRepo  repository.UserRepository
}

func NewUser(container *storage.Container) *User {
	return &User{
		container: container,
		userRepo:  repository.NewUser(container),
	}
}
