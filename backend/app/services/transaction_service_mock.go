package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

func NewTestTransaction(container *storage.Container) *Transaction {
	return &Transaction{
		container:     container,
		repo:          repository.NewTestTransaction(container),
		repoPortfolio: repository.NewTestPortfolio(container),
		repoAccount:   repository.NewTestAccount(container),
		repoCategory:  repository.NewTestCategory(container),
	}
}
