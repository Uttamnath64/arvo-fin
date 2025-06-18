package services

import (
	"testing"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/stretchr/testify/assert"
)

func NewTestAvatar() (*Avatar, bool) {
	container, ok := getTestContainer()
	if !ok {
		return nil, false
	}

	return &Avatar{
		container:     container,
		avatarService: services.NewTestAvatar(container),
	}, true
}

func TestGetAvatarsByType_Avatar(t *testing.T) {
	service, ok := NewTestAvatar()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		userId      uint
		avatarType  commonType.AvatarType
		expectError bool
	}{
		{
			name:        "Valid",
			avatarType:  commonType.AvatarTypeDefault,
			expectError: false,
		},
		{
			name:        "Valid",
			avatarType:  commonType.AvatarTypeUser,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.GetAvatarsByType(tt.avatarType)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestGet_Avatar(t *testing.T) {
	service, ok := NewTestAvatar()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		Id          uint
		expectError bool
	}{
		{
			name:        "Valid",
			Id:          1,
			expectError: false,
		},
		{
			name:        "Not Found",
			Id:          290,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.Get(tt.Id)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestCreate_Avatar(t *testing.T) {
	service, ok := NewTestAvatar()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		payload     requests.AvatarRequest
		expectError bool
	}{
		{
			name: "Valid",
			payload: requests.AvatarRequest{
				Name: "Test",
				Type: commonType.AvatarTypeDefault,
				Icon: "T",
			},
			expectError: false,
		},
		{
			name: "Valid",
			payload: requests.AvatarRequest{
				Name: "Test",
				Type: commonType.AvatarTypeDefault,
				Icon: "T",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.Create(tt.payload)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestUpdate_Avatar(t *testing.T) {
	service, ok := NewTestAvatar()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		Id          uint
		payload     requests.AvatarRequest
		userId      uint
		userType    commonType.UserType
		expectError bool
	}{
		{
			name:     "Valid",
			Id:       1,
			userId:   1,
			userType: commonType.UserTypeAdmin,
			payload: requests.AvatarRequest{
				Name: "Test",
				Type: commonType.AvatarTypeDefault,
				Icon: "T",
			},
			expectError: false,
		},
		{
			name:   "InValid",
			Id:     2,
			userId: 1,
			payload: requests.AvatarRequest{
				Name: "Test",
				Type: commonType.AvatarTypeDefault,
				Icon: "T",
			},
			userType:    commonType.UserTypeAdmin,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.Update(tt.Id, tt.payload)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}
