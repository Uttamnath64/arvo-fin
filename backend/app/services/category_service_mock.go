package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

func NewTestCategory(container *storage.Container) *Category {
	return &Category{
		container:     container,
		repoCategory:  repository.NewTestCategory(container),
		repoAvatar:    repository.NewTestAvatar(container),
		repoPortfolio: repository.NewTestPortfolio(container),
	}
}
