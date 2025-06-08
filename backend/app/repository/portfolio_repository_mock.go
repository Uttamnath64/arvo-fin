package repository

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type TestPortfolio struct {
	container *storage.Container
}

func NewTestPortfolio(container *storage.Container) *TestPortfolio {
	return &TestPortfolio{
		container: container,
	}
}

func (repo *TestPortfolio) GetList(userId uint, userType commonType.UserType) (*[]responses.PortfolioResponse, error) {
	if userId == 1 && userType == commonType.UserTypeUser {
		portfolios := []responses.PortfolioResponse{
			{
				Id:        1,
				Name:      "Retirement Fund",
				AvatarID:  101,
				AvatarURL: "https://example.com/avatars/retirement.png",
			},
			{
				Id:        2,
				Name:      "Crypto Portfolio",
				AvatarID:  102,
				AvatarURL: "https://example.com/avatars/crypto.png",
			},
		}
		return &portfolios, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (repo *TestPortfolio) Get(id, userId uint, userType commonType.UserType) (*responses.PortfolioResponse, error) {
	if id == 1 && userId == 1 && userType == commonType.UserTypeUser {
		portfolios := responses.PortfolioResponse{
			Id:        1,
			Name:      "Retirement Fund",
			AvatarID:  101,
			AvatarURL: "https://example.com/avatars/retirement.png",
		}
		return &portfolios, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (repo *TestPortfolio) Add(portfolio models.Portfolio) error {
	return nil
}

func (repo *TestPortfolio) Update(id, userId uint, payload requests.PortfolioRequest) error {
	if id == 1 && userId == 1 {
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (repo *TestPortfolio) Delete(id, userId uint) error {
	if id == 1 && userId == 1 {
		return nil
	}
	return gorm.ErrRecordNotFound
}
