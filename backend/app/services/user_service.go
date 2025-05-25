package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type User struct {
	container *storage.Container
	repo      repository.UserRepository
}

func NewUserService(container *storage.Container) *User {
	return &User{
		container: container,
		repo:      repository.NewUser(container),
	}
}
