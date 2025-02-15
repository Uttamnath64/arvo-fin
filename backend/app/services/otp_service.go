package services

import (
	"errors"
	"fmt"
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type OTPService struct {
	RedisClient *storage.RedisClient
	TTL         int
}

func NewOTPService(redisClient *storage.RedisClient, ttl int) *OTPService {
	return &OTPService{
		RedisClient: redisClient,
		TTL:         ttl,
	}
}

// GenerateOTP generates a random OTP (for simplicity, hardcoded here).
func (service *OTPService) GenerateOTP() string {
	return fmt.Sprintf("%06d", time.Now().UnixNano()%1000000) // 6-digit OTP
}

// SaveOTP stores the OTP in Redis
func (service *OTPService) SaveOTP(email string, otpType commonType.OtpType, otp string) error {
	key := fmt.Sprintf("OTP:email=%s&type=%d", email, otpType)
	err := service.RedisClient.SetValue(key, otp, service.TTL)
	if err != nil {
		return fmt.Errorf("failed to save OTP: %v", err)
	}
	return nil
}

// VerifyOTP verifies a user-provided OTP against the stored OTP
func (service *OTPService) VerifyOTP(email string, otpType commonType.OtpType, providedOTP string) error {
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
