package services

import (
	"testing"

	"github.com/Uttamnath64/arvo-fin/app/auth"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	appService "github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/stretchr/testify/assert"
)

func NewTestAuth() (*Auth, bool) {
	container, ok := getTestContainer()
	if !ok {
		return nil, false
	}

	authRepo := repository.NewTestAuth(container)
	return &Auth{
		container:    container,
		userRepo:     repository.NewTestUser(container),
		authRepo:     authRepo,
		avatarRepo:   repository.NewTestAvatar(container),
		authHelper:   auth.New(container, authRepo),
		otpService:   appService.NewTestOTP(container.Redis, 300),
		emailService: appService.NewTestEmail(container),
	}, true
}

func TestLogin(t *testing.T) {
	authService, ok := NewTestAuth()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		payload     requests.LoginRequest
		expectError bool
	}{
		{
			name: "Faild To Login",
			payload: requests.LoginRequest{
				UsernameEmail: "uttam1@example.com",
				Password:      "Secret@123",
			},
			expectError: true,
		},
		{
			name: "Login",
			payload: requests.LoginRequest{
				UsernameEmail: "uttam@example.com",
				Password:      "Secret@123",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := authService.Login(tt.payload, "", "")
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestRegister(t *testing.T) {
	authService, ok := NewTestAuth()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		payload     requests.RegisterRequest
		expectError bool
	}{
		{
			name: "Faild To Register",
			payload: requests.RegisterRequest{
				Name:     "uttam nath",
				Email:    "uttam@example.com",
				Username: "uttam.nath",
				AvatarId: 1,
				Password: "Secret@123",
				OTP:      "123443",
			},
			expectError: true,
		},
		{
			name: "Register",
			payload: requests.RegisterRequest{
				Name:     "uttam nath",
				Email:    "uttam-new@example.com",
				Username: "uttam.nath1",
				AvatarId: 1,
				Password: "Secret@123",
				OTP:      "123456",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := authService.Register(tt.payload, "", "")
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestSentOTP(t *testing.T) {
	authService, ok := NewTestAuth()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		payload     requests.SentOTPRequest
		expectError bool
	}{
		{
			name: "Send OTP",
			payload: requests.SentOTPRequest{
				Email: "uttam@example.com",
				Type:  commonType.Register,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := authService.SentOTP(tt.payload)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestResetPassword(t *testing.T) {
	authService, ok := NewTestAuth()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		payload     requests.ResetPasswordRequest
		expectError bool
	}{
		{
			name: "Reset Password",
			payload: requests.ResetPasswordRequest{
				Email:    "uttam@example.com",
				Password: "Secret@1234",
				OTP:      "123456",
			},
			expectError: false,
		},
		{
			name: "Invalid Email",
			payload: requests.ResetPasswordRequest{
				Email:    "uttam1@example.com",
				Password: "Secret@1234",
				OTP:      "123456",
			},
			expectError: true,
		},
		{
			name: "Same Password",
			payload: requests.ResetPasswordRequest{
				Email:    "uttam@example.com",
				Password: "Secret@123",
				OTP:      "123456",
			},
			expectError: true,
		},
		{
			name: "Invalid OTP",
			payload: requests.ResetPasswordRequest{
				Email:    "uttam@example.com",
				Password: "Secret@1234",
				OTP:      "123406",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := authService.ResetPassword(tt.payload, "", "")
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestGetToken(t *testing.T) {
	authService, ok := NewTestAuth()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		payload     requests.TokenRequest
		expectError bool
	}{
		{
			name: "Valid",
			payload: requests.TokenRequest{
				RefreshToken: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX3R5cGUiOjEsInNlc3Npb25faWQiOjI2LCJleHAiOjczODI5NDUxMzN9.GFpuw1k2r65A2clj5D2PzXfaXq7t-GH0oWZlEtewZ_oDwykzbhzYGTTOlNPnenaFBltw-L3bJbpEpSbg0gdgBEOGQlYNaGNOl-6SJtCCXGQEQqcbWmYBMROVf39NFJAl3ZOrCQf8SSsLi7QjYKL9dFwKEpnMYXWCypf6HzjvWrqUo42O6GeXXxSFVSmQrUzB88WrkWCTFiOznmn_d8TkyyrEPGg6gfNk0Jkq0CcTNnMznxBdPNK8eWnfSzSsyWK-1NfujFL4ZYY7A6dX8SiJgtU8qZmsbnqAAm3RcJP2veOnc3YqdH6OzhCPBVzm_XX7C-faYNy6kopf45MDlrjkvyJnbESLuy7EMRf3q6xi4ozrDckfJrdybURGeXkACOvS0JpjSYjhWhHlfyjHwTFn7o--Vp6H1-RVYJTWQjqsEOGx9Ie6N3lr9Q24H1K-JHhaa2Hgp-SJcp4oEMK7vDYBetmfSglVKlDDjH_bN4W07O_9s4g5bSs5DFhnD8K_haCFUcwzk4qUOoKogSu8xW0UlPPJdxw0gCWKY8bFIVBaBtN1yoMn8iIyYPiz0fNjyi2IKt5FwigKewiId5TW77UZA4J4Zo4-F0Eke3eLKZ3986ouDRNI3T2Gf55HqRMnSU5w6naM8FL7w7l97V1Gn01BwzAZsmGesYarGP0WlhxxghM",
			},
			expectError: false,
		},
		{
			name: "Invalid",
			payload: requests.TokenRequest{
				RefreshToken: "123456765gefdfsersdfgfdfskjbhdfyuvdsfutwsicgu",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := authService.GetToken(tt.payload, "", "")
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}
