package services

import (
	"github.com/Uttamnath64/arvo-fin/app/common"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
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

func (service *Portfolio) GetList(rctx *requests.RequestContext, userId uint) responses.ServiceResponse {

	response, err := service.portfolioRepo.GetList(rctx, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "You haven’t created any portfolios yet.", err)
		}

		service.container.Logger.Error("portfolio.appService.getList-GetList", "error", err.Error(), "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("Portfolios retrieved successfully!", response)
}

func (service *Portfolio) Get(rctx *requests.RequestContext, id, userId uint, userType commonType.UserType) responses.ServiceResponse {

	response, err := service.portfolioRepo.Get(rctx, id, userId, userType)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "We couldn't find the requested portfolio.", err)
		}

		service.container.Logger.Error("portfolio.appService.get-Get", "error", err.Error(), "id", id, "userId", userId, "userType", userType)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("Great! We found your portfolios.", response)
}
