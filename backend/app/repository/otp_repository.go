package repository

import (
	"strings"

	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type OTPRepository struct {
	container *storage.Container
}

func NewOTPRepository(container *storage.Container) *OTPRepository {
	return &OTPRepository{
		container: container,
	}
}

func (repo *OTPRepository) Exists(email string, otp int) (bool, error) {
	var count int64

	err := repo.container.Config.ReadOnlyDB.Model(&models.OTP{}).
		Where("email = ? and otp = ?", strings.ToLower(email), otp).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
