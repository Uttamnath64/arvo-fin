package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

func NewTestUser(container *storage.Container) *User {
	return &User{
		container: container,
		userRepo:  repository.NewTestUser(container),
	}
}
