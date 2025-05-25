package services

import (
	"errors"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type TestOTP struct {
	RedisClient *storage.RedisClient
	TTL         int
}

func NewTestOTP(redisClient *storage.RedisClient, ttl int) *TestOTP {
	return &TestOTP{
		RedisClient: redisClient,
		TTL:         ttl,
	}
}

func (service *TestOTP) GenerateOTP() string {
	return "123456"
}

func (service *TestOTP) SaveOTP(email string, otpType commonType.OtpType, otp string) error {
	return nil
}

func (service *TestOTP) VerifyOTP(email string, otpType commonType.OtpType, providedOTP string) error {
	if (email == "uttam@example.com" || email == "uttam-new@example.com") && providedOTP == "123456" {
		return nil
	}
	return errors.New("Not found!")
}
