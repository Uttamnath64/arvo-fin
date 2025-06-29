package services

import (
	"errors"
	"fmt"
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type OTP struct {
	RedisClient *storage.RedisClient
	TTL         int
}

func NewOTP(redisClient *storage.RedisClient, ttl int) *OTP {
	return &OTP{
		RedisClient: redisClient,
		TTL:         ttl,
	}
}

// GenerateOTP generates a random OTP (for simplicity, hardcoded here).
func (service *OTP) GenerateOTP() string {
	return fmt.Sprintf("%06d", time.Now().UnixNano()%1000000) // 6-digit OTP
}

// SaveOTP stores the OTP in Redis
func (service *OTP) SaveOTP(email string, otpType commonType.OtpType, otp string) error {
	key := fmt.Sprintf("OTP:email=%s&type=%d", email, otpType)
	err := service.RedisClient.SetValue(key, otp, service.TTL)
	if err != nil {
		return fmt.Errorf("failed to save OTP: %w", err)
	}
	return nil
}

// VerifyOTP verifies a user-provided OTP against the stored OTP
func (service *OTP) VerifyOTP(email string, otpType commonType.OtpType, providedOTP string) error {
	key := fmt.Sprintf("OTP:email=%s&type=%d", email, otpType)
	storedOTP, err := service.RedisClient.GetValue(key)
	if err != nil {
		return errors.New("OTP expired")
	}

	if storedOTP != providedOTP {
		return errors.New("invalid OTP")
	}

	// Delete the OTP after successful verification
	_ = service.RedisClient.DeleteKey(key)

	return nil
}
