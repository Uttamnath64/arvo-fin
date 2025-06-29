package requests

import (
	"errors"
	"strings"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

// Login payload
type LoginRequest struct {
	UsernameEmail string `json:"username_email" binding:"required"`
	Password      string `json:"password" binding:"required"`
}

func (r LoginRequest) IsValid() error {
	emailErr := Validate.IsValidEmail(r.UsernameEmail)
	usernameErr := Validate.IsValidUsername(r.UsernameEmail)
	if emailErr != nil && usernameErr != nil {
		if emailErr != nil {
			return emailErr
		} else {
			return usernameErr
		}
	}

	if passErr := Validate.IsValidPassword(r.Password); passErr != nil {
		return passErr
	}
	return nil
}

// Register payload
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	AvatarId uint   `json:"avatar_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	OTP      string `json:"otp" binding:"required"`
}

func (r RegisterRequest) IsValid() error {
	if err := Validate.IsValidName(r.Name); err != nil {
		return err
	}
	if err := Validate.IsValidUsername(r.Username); err != nil {
		return err
	}
	if err := Validate.IsValidEmail(r.Email); err != nil {
		return err
	}
	if !Validate.IsValidID(r.AvatarId) {
		return errors.New("Please select a valid avatar.")
	}
	if err := Validate.IsValidPassword(r.Password); err != nil {
		return err
	}
	if err := Validate.IsValidOTP(r.OTP); err != nil {
		return err
	}
	return nil
}

// Send OTP payload
type SentOTPRequest struct {
	Email string             `json:"email" binding:"required"`
	Type  commonType.OtpType `json:"type" binding:"required"`
}

func (r SentOTPRequest) IsValid() error {
	if err := Validate.IsValidEmail(r.Email); err != nil {
		return err
	}
	if !r.Type.IsValid() {
		return errors.New("Invalid OTP type.")
	}
	return nil
}

// Reset Password payload
type ResetPasswordRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	OTP      string `json:"otp" binding:"required"`
}

func (r ResetPasswordRequest) IsValid() error {
	if err := Validate.IsValidEmail(r.Email); err != nil {
		return err
	}
	if err := Validate.IsValidPassword(r.Password); err != nil {
		return err
	}
	if err := Validate.IsValidOTP(r.OTP); err != nil {
		return err
	}
	return nil
}

// Token payload
type TokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (r TokenRequest) IsValid() error {
	if strings.TrimSpace(r.RefreshToken) == "" {
		return errors.New("Refresh token is required.")
	}
	return nil
}
