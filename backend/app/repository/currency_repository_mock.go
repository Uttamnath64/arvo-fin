package repository

import (
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type TestCurrency struct {
	container *storage.Container
}

func NewTestCurrency(container *storage.Container) *TestCurrency {
	return &TestCurrency{
		container: container,
	}
}

func (repo *TestCurrency) CodeExists(code string) (bool, error) {
	if code != "INR" {
		return false, gorm.ErrRecordNotFound
	}
	return true, nil
}
