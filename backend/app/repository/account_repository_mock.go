package repository

import (
	"errors"
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type TestAccount struct {
	container *storage.Container
}

func NewTestAccount(container *storage.Container) *TestAccount {
	return &TestAccount{
		container: container,
	}
}

func (repo *TestAccount) GetList(portfolioId, userId uint) (*[]models.Account, error) {
	if userId != 1 || portfolioId != 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return &[]models.Account{
		{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Add(-1 * time.Hour),
				UpdatedAt: time.Now(),
			},
			UserId:         1,
			PortfolioId:    1,
			Name:           "Test Account",
			Type:           commonType.AccountTypeBank,
			OpeningBalance: 100,
			Note:           "Testing.......",
			Avatar: models.Avatar{
				BaseModel: models.BaseModel{
					ID:        1,
					CreatedAt: time.Now().Add(-1 * time.Hour),
					UpdatedAt: time.Now(),
				},
				Name: "Test Avatar",
				Icon: "T",
				Type: commonType.AvatarTypeDefault,
			},
			Currency: models.Currency{
				Code:   "INR",
				Name:   "Indian Rupee",
				Symbol: "₹",
			},
		},
		{
			BaseModel: models.BaseModel{
				ID:        2,
				CreatedAt: time.Now().Add(-1 * time.Hour),
				UpdatedAt: time.Now(),
			},
			UserId:         1,
			PortfolioId:    1,
			Name:           "Test Account 2",
			Type:           commonType.AccountTypeCash,
			OpeningBalance: 200,
			Note:           "Testing.......",
			Avatar: models.Avatar{
				BaseModel: models.BaseModel{
					ID:        1,
					CreatedAt: time.Now().Add(-1 * time.Hour),
					UpdatedAt: time.Now(),
				},
				Name: "Test Avatar 2",
				Icon: "G",
				Type: commonType.AvatarTypeDefault,
			},
			Currency: models.Currency{
				Code:   "INR",
				Name:   "Indian Rupee",
				Symbol: "₹",
			},
		},
	}, nil
}

func (repo *TestAccount) Get(id uint) (*models.Account, error) {
	if id != 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return &models.Account{
		BaseModel: models.BaseModel{
			ID:        1,
			CreatedAt: time.Now().Add(-1 * time.Hour),
			UpdatedAt: time.Now(),
		},
		UserId:         1,
		PortfolioId:    1,
		Name:           "Test Account",
		Type:           commonType.AccountTypeBank,
		OpeningBalance: 100,
		Note:           "Testing.......",
		Avatar: models.Avatar{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Add(-1 * time.Hour),
				UpdatedAt: time.Now(),
			},
			Name: "Test Avatar",
			Icon: "T",
			Type: commonType.AvatarTypeDefault,
		},
		Currency: models.Currency{
			Code:   "INR",
			Name:   "Indian Rupee",
			Symbol: "₹",
		},
	}, nil
}

func (repo *TestAccount) Create(account models.Account) (uint, error) {
	return 1, nil
}

func (repo *TestAccount) Update(id, userId uint, payload requests.AccountUpdateRequest) (bool, error) {
	if id != 1 || userId != 1 {
		return false, errors.New("Error!")
	}

	return true, nil
}

func (repo *TestAccount) Delete(id, userId uint) (bool, error) {
	if id != 1 || userId != 1 {
		return false, errors.New("Error!")
	}

	return true, nil
}
