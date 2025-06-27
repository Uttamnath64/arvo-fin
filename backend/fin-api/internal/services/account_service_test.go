package services

import (
	"context"
	"testing"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	commonServices "github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/stretchr/testify/assert"
)

func NewTestAccount() (*Account, *requests.RequestContext, bool) {
	container, ok := getTestContainer()
	if !ok {
		return nil, nil, false
	}

	return &Account{
			container:      container,
			accountService: commonServices.NewTestAccount(container),
		}, &requests.RequestContext{
			Ctx:       context.Background(),
			UserID:    1,
			UserType:  commonType.UserTypeUser,
			SessionID: 1,
		}, true
}

func TestGetList_Account(t *testing.T) {
	service, rctx, ok := NewTestAccount()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		userId      uint
		portfolioId uint
		expectError bool
	}{
		{
			name:        "Valid",
			userId:      1,
			portfolioId: 1,
			expectError: false,
		},
		{
			name:        "Not Found",
			userId:      1,
			portfolioId: 3,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.GetList(rctx, tt.userId, tt.portfolioId)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestGet_Account(t *testing.T) {
	service, rctx, ok := NewTestAccount()
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
			Id:          2,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.Get(rctx, tt.Id)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestCreate_Account(t *testing.T) {
	service, rctx, ok := NewTestAccount()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		payload     requests.AccountRequest
		userId      uint
		userType    commonType.UserType
		expectError bool
	}{
		{
			name: "Valid",
			payload: requests.AccountRequest{
				AvatarId:       1,
				PortfolioId:    1,
				Name:           "Test Account",
				Type:           commonType.AccountTypeBank,
				CurrencyCode:   "INR",
				OpeningBalance: 100,
				Note:           "Testing.......",
			},
			userId:      1,
			expectError: false,
		},
		{
			name: "Not Found",
			payload: requests.AccountRequest{
				AvatarId:       1,
				PortfolioId:    1,
				Name:           "Test Account",
				Type:           commonType.AccountTypeBank,
				CurrencyCode:   "INR",
				OpeningBalance: 100,
				Note:           "Testing.......",
			},
			userId:      2,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.Create(rctx, tt.userId, tt.payload)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestUpdate_Account(t *testing.T) {
	service, rctx, ok := NewTestAccount()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		Id          uint
		userId      uint
		payload     requests.AccountUpdateRequest
		expectError bool
	}{
		{
			name:   "Valid",
			Id:     1,
			userId: 1,
			payload: requests.AccountUpdateRequest{
				AvatarId:     1,
				Name:         "Test Account",
				Type:         commonType.AccountTypeBank,
				CurrencyCode: "INR",
				Note:         "Testing.......",
			},
			expectError: false,
		},
		{
			name:   "Not Found",
			Id:     2,
			userId: 1,
			payload: requests.AccountUpdateRequest{
				AvatarId:     2,
				Name:         "Test Account",
				Type:         commonType.AccountTypeBank,
				CurrencyCode: "INR",
				Note:         "Testing.......",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.Update(rctx, tt.Id, tt.userId, tt.payload)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestDelete_Account(t *testing.T) {
	service, rctx, ok := NewTestAccount()
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

			serviceResponse := service.Delete(rctx, tt.Id, tt.userId)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}
