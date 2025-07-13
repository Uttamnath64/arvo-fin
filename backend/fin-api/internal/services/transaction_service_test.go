package services

import (
	"context"
	"testing"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	commonServices "github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/pkg/pagination"
	"github.com/stretchr/testify/assert"
)

func NewTestTransaction() (*Transaction, *requests.RequestContext, bool) {
	container, ok := getTestContainer()
	if !ok {
		return nil, nil, false
	}

	return &Transaction{
			container: container,
			service:   commonServices.NewTestTransaction(container),
		}, &requests.RequestContext{
			Ctx:       context.Background(),
			UserID:    1,
			UserType:  commonType.UserTypeUser,
			SessionID: 1,
		}, true
}

func TestGetList_Transaction(t *testing.T) {
	service, rctx, ok := NewTestTransaction()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		expectError bool
	}{
		{
			name:        "Valid",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.GetList(rctx, requests.TransactionQuery{}, pagination.Pagination{Page: 1, Limit: 10})
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestGet_Transaction(t *testing.T) {
	service, rctx, ok := NewTestTransaction()
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

func TestCreate_Transaction(t *testing.T) {
	service, rctx, ok := NewTestTransaction()
	if !ok {
		return
	}

	transferAccountId := uint(1)
	tests := []struct {
		name        string
		payload     requests.TransactionRequest
		userType    commonType.UserType
		expectError bool
	}{
		{
			name: "Valid",
			payload: requests.TransactionRequest{
				TransferAccountId: &transferAccountId,
				AccountId:         1,
				CategoryId:        1,
				PortfolioId:       1,
				Amount:            100.0,
				Type:              commonType.TransactionType(1),
				Note:              "test",
			},
			expectError: false,
		},
		{
			name: "Not Found",
			payload: requests.TransactionRequest{
				TransferAccountId: &transferAccountId,
				AccountId:         1,
				CategoryId:        1,
				PortfolioId:       3,
				Amount:            100.0,
				Type:              commonType.TransactionType(1),
				Note:              "test",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.Create(rctx, tt.payload)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestUpdate_Transaction(t *testing.T) {
	service, rctx, ok := NewTestTransaction()
	if !ok {
		return
	}

	transferAccountId := uint(1)
	tests := []struct {
		name        string
		Id          uint
		payload     requests.TransactionRequest
		expectError bool
	}{
		{
			name: "Valid",
			Id:   1,
			payload: requests.TransactionRequest{
				TransferAccountId: &transferAccountId,
				AccountId:         1,
				CategoryId:        1,
				PortfolioId:       1,
				Amount:            100.0,
				Type:              commonType.TransactionType(1),
				Note:              "test",
			},
			expectError: false,
		},
		{
			name: "Not Found",
			Id:   2,
			payload: requests.TransactionRequest{
				TransferAccountId: &transferAccountId,
				AccountId:         1,
				CategoryId:        1,
				PortfolioId:       3,
				Amount:            100.0,
				Type:              commonType.TransactionType(1),
				Note:              "test",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.Update(rctx, tt.Id, tt.payload)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestDelete_Transaction(t *testing.T) {
	service, rctx, ok := NewTestTransaction()
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

			serviceResponse := service.Delete(rctx, tt.Id)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}
