package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

func NewTestPortfolio(container *storage.Container) *Portfolio {
	return &Portfolio{
		container:     container,
		portfolioRepo: repository.NewTestPortfolio(container),
		avatarRepo:    repository.NewTestAvatar(container),
	}
}
