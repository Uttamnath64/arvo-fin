package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

func NewTestUser(container *storage.Container) *User {
	return &User{
		container:    container,
		repoUser:     repository.NewTestUser(container),
		repoAvatar:   repository.NewTestAvatar(container),
		repoCurrency: repository.NewTestCurrency(container),
	}
}
