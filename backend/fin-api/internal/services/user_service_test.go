package services

import (
	"testing"

	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/stretchr/testify/assert"
)

func NewTestUser() (*User, bool) {
	container, ok := getTestContainer()
	if !ok {
		return nil, false
	}
	return &User{
		container:   container,
		userService: services.NewTestUser(container),
	}, true
}

func TestGet_User(t *testing.T) {
	userService, ok := NewTestUser()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		userId      uint
		expectError bool
	}{
		{
			name:        "Valid",
			userId:      1,
			expectError: false,
		},
		{
			name:        "Not Found",
			userId:      10,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := userService.Get(tt.userId)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestGetSetting_User(t *testing.T) {
	userService, ok := NewTestUser()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		userId      uint
		expectError bool
	}{
		{
			name:        "Valid",
			userId:      1,
			expectError: false,
		},
		{
			name:        "Not Found",
			userId:      10,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := userService.GetSettings(tt.userId)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestUpdate_User(t *testing.T) {
	userService, ok := NewTestUser()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		Id          uint
		userId      uint
		payload     requests.MeRequest
		expectError bool
	}{
		{
			name:   "Valid",
			userId: 1,
			payload: requests.MeRequest{
				Name:     "Uttam Nath",
				Username: "uttam.nath",
				AvatarId: 1,
			},
			expectError: false,
		},
		{
			name:   "Not Found",
			userId: 10,
			payload: requests.MeRequest{
				Name:     "Uttam Nath",
				Username: "uttam.nath",
				AvatarId: 1,
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := userService.Update(tt.payload, tt.userId)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestUpdateSetting_User(t *testing.T) {
	userService, ok := NewTestUser()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		Id          uint
		userId      uint
		payload     requests.SettingsRequest
		expectError bool
	}{
		{
			name:   "Valid",
			userId: 1,
			payload: requests.SettingsRequest{
				CurrencyCode:       "INR",
				DecimalPlaces:      3,
				NumberFormat:       1,
				EmailNotifications: true,
			},
			expectError: false,
		},
		{
			name:   "Not Found",
			userId: 10,
			payload: requests.SettingsRequest{
				CurrencyCode:       "INR",
				DecimalPlaces:      3,
				NumberFormat:       1,
				EmailNotifications: true,
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := userService.UpdateSettings(tt.payload, tt.userId)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}
