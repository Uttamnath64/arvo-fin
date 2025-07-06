package services

import (
	"context"
	"testing"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	commonServices "github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/stretchr/testify/assert"
)

func NewTestCategory() (*Category, *requests.RequestContext, bool) {
	container, ok := getTestContainer()
	if !ok {
		return nil, nil, false
	}

	return &Category{
			container: container,
			service:   commonServices.NewTestCategory(container),
		}, &requests.RequestContext{
			Ctx:       context.Background(),
			UserID:    1,
			UserType:  commonType.UserTypeUser,
			SessionID: 1,
		}, true
}

func TestGetList_Category(t *testing.T) {
	service, rctx, ok := NewTestCategory()
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

			serviceResponse := service.GetList(rctx, tt.portfolioId, tt.userId)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}

func TestGet_Category(t *testing.T) {
	service, rctx, ok := NewTestCategory()
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

func TestCreate_Category(t *testing.T) {
	service, rctx, ok := NewTestCategory()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		payload     requests.CategoryRequest
		userId      uint
		userType    commonType.UserType
		expectError bool
	}{
		{
			name: "Valid",
			payload: requests.CategoryRequest{
				AvatarId:    1,
				PortfolioId: 1,
				Type:        commonType.TransactionTypeExpense,
				Name:        "Testing",
			},
			userId:      1,
			expectError: false,
		},
		{
			name: "Not Found",
			payload: requests.CategoryRequest{
				AvatarId:    2,
				PortfolioId: 2,
				Type:        commonType.TransactionTypeExpense,
				Name:        "Testing",
			},
			userId:      2,
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

func TestUpdate_Category(t *testing.T) {
	service, rctx, ok := NewTestCategory()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		Id          uint
		userId      uint
		payload     requests.CategoryRequest
		expectError bool
	}{
		{
			name:   "Valid",
			Id:     1,
			userId: 1,
			payload: requests.CategoryRequest{
				AvatarId:    1,
				PortfolioId: 1,
				Type:        commonType.TransactionTypeExpense,
				Name:        "Testing",
			},
			expectError: false,
		},
		{
			name:   "Not Found",
			Id:     2,
			userId: 1,
			payload: requests.CategoryRequest{
				AvatarId:    1,
				PortfolioId: 1,
				Type:        commonType.TransactionTypeExpense,
				Name:        "Testing",
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

func TestDelete_Category(t *testing.T) {
	service, rctx, ok := NewTestCategory()
	if !ok {
		return
	}

	tests := []struct {
		name        string
		portfolioId uint
		id          uint
		expectError bool
	}{
		{
			name:        "Valid",
			portfolioId: 1,
			id:          1,
			expectError: false,
		},
		{
			name:        "Not Found",
			portfolioId: 2,
			id:          21,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceResponse := service.Delete(rctx, tt.portfolioId, tt.id)
			assert.Equal(t, tt.expectError, serviceResponse.HasError())
		})
	}
}
