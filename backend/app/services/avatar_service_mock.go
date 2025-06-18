package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

func NewTestAvatar(container *storage.Container) *Avatar {
	return &Avatar{
		container:  container,
		repoAvatar: repository.NewTestAvatar(container),
	}
}
