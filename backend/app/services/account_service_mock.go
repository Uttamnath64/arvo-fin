package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

func NewTestAccount(container *storage.Container) *Account {
	return &Account{
		container:     container,
		repoAccount:   repository.NewTestAccount(container),
		repoAvatar:    repository.NewTestAvatar(container),
		repoPortfolio: repository.NewTestPortfolio(container),
		repoCurrency:  repository.NewTestCurrency(container),
	}
}
