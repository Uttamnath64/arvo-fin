package services

import (
	"github.com/Uttamnath64/arvo-fin/app/common"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type Portfolio struct {
	container     *storage.Container
	portfolioRepo repository.PortfolioRepository
	avatarRepo    repository.AvatarRepository
}

func NewPortfolio(container *storage.Container) *Portfolio {
	return &Portfolio{
		container:     container,
		portfolioRepo: repository.NewPortfolio(container),
		avatarRepo:    repository.NewAvatar(container),
	}
}

func (service *Portfolio) GetList(userId uint, userType commonType.UserType) responses.ServiceResponse {

	response, err := service.portfolioRepo.GetList(userId, userType)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Portfolios not found!", err)
		}

		service.container.Logger.Error("portfolio.appService.getList-GetList", "error", err.Error(), "userId", userId, "userType", userType, "userTypeName", userType.String())
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("Portfolios records found!", response)
}

func (service *Portfolio) Get(id, userId uint, userType commonType.UserType) responses.ServiceResponse {

	response, err := service.portfolioRepo.Get(id, userId, userType)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Portfolio not found!", err)
		}

		service.container.Logger.Error("portfolio.appService.get-Get", "error", err.Error(), "id", id, "userId", userId, "userType", userType, "userTypeName", userType.String())
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("Portfolio records found!", response)
}
