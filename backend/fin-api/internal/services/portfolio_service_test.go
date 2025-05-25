package services

import (
	"testing"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	commonServices "github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/stretchr/testify/assert"
)

func NewTestPortfolio() (*Portfolio, bool) {
	container, ok := getTestContainer()
	if !ok {
		return nil, false
	}

	return &Portfolio{
		container:        container,
		portfolioService: commonServices.NewTestPortfolio(container),
		portfolioRepo:    repository.NewTestPortfolio(container),
		avatarRepo:       repository.NewTestAvatar(container),
	}, true
}

func TestGetList(t *testing.T) {
	portfolioService, ok := NewTestPortfolio()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		userId      uint
		userType    commonType.UserType
		expectError bool
	}{
		{
			name:        "Valid",
			userId:      1,
			userType:    commonType.User,
			expectError: false,
		},
		{
			name:        "Not Found",
			userId:      1,
			userType:    commonType.Admin,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := portfolioService.GetList(tt.userId, tt.userType)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestGet(t *testing.T) {
	portfolioService, ok := NewTestPortfolio()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		Id          uint
		userId      uint
		userType    commonType.UserType
		expectError bool
	}{
		{
			name:        "Valid",
			Id:          1,
			userId:      1,
			userType:    commonType.User,
			expectError: false,
		},
		{
			name:        "Not Found",
			Id:          2,
			userId:      1,
			userType:    commonType.User,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := portfolioService.Get(tt.Id, tt.userId, tt.userType)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestAdd(t *testing.T) {
	portfolioService, ok := NewTestPortfolio()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		payload     requests.PortfolioRequest
		userId      uint
		userType    commonType.UserType
		expectError bool
	}{
		{
			name: "Valid",
			payload: requests.PortfolioRequest{
				Name:     "Test",
				AvatarId: 1,
			},
			userId:      1,
			expectError: false,
		},
		{
			name: "Not Found",
			payload: requests.PortfolioRequest{
				Name:     "Test",
				AvatarId: 2,
			},
			userId:      1,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := portfolioService.Add(tt.payload, tt.userId)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestUpdate(t *testing.T) {
	portfolioService, ok := NewTestPortfolio()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		Id          uint
		userId      uint
		payload     requests.PortfolioRequest
		expectError bool
	}{
		{
			name:   "Valid",
			Id:     1,
			userId: 1,
			payload: requests.PortfolioRequest{
				Name:     "Test2",
				AvatarId: 1,
			},
			expectError: false,
		},
		{
			name:   "Not Found",
			Id:     2,
			userId: 1,
			payload: requests.PortfolioRequest{
				Name:     "Test1",
				AvatarId: 1,
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := portfolioService.Update(tt.Id, tt.userId, tt.payload)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestDelete(t *testing.T) {
	portfolioService, ok := NewTestPortfolio()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		Id          uint
		userId      uint
		expectError bool
	}{
		{
			name:        "Valid",
			Id:          1,
			userId:      1,
			expectError: false,
		},
		{
			name:        "Not Found",
			Id:          2,
			userId:      1,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := portfolioService.Delete(tt.Id, tt.userId)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}
