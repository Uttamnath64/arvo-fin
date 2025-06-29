package services

import (
	"errors"

	"github.com/Uttamnath64/arvo-fin/app/common"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	commonServices "github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type Portfolio struct {
	container        *storage.Container
	portfolioService commonServices.PortfolioService
	portfolioRepo    repository.PortfolioRepository
	avatarRepo       repository.AvatarRepository
}

func NewPortfolio(container *storage.Container) *Portfolio {
	return &Portfolio{
		container:        container,
		portfolioService: commonServices.NewPortfolio(container),
		portfolioRepo:    repository.NewPortfolio(container),
		avatarRepo:       repository.NewAvatar(container),
	}
}

func (service *Portfolio) GetList(rctx *requests.RequestContext, userId uint) responses.ServiceResponse {
	return service.portfolioService.GetList(rctx, userId)
}

func (service *Portfolio) Get(rctx *requests.RequestContext, id, userId uint, userType commonType.UserType) responses.ServiceResponse {
	return service.portfolioService.Get(rctx, id, userId, userType)
}

func (service *Portfolio) Create(rctx *requests.RequestContext, payload requests.PortfolioRequest, userId uint) responses.ServiceResponse {

	portfolio := models.Portfolio{
		Name:     payload.Name,
		AvatarId: payload.AvatarId,
		UserId:   userId,
	}

	// Verify avatar
	if err := service.avatarRepo.AvatarByTypeExists(rctx, payload.AvatarId, commonType.AvatarTypePortfolio); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatar not found!", errors.New("avatar not found"))
		}

		service.container.Logger.Error("portfolio.service.create-AvatarByTypeExists", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", commonType.AvatarTypePortfolio)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later", err)
	}

	// Create
	if err := service.portfolioRepo.Create(rctx, portfolio); err != nil {
		service.container.Logger.Error("portfolio.service.create-Create", "error", err.Error(), "portfolio", portfolio)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later", err)
	}

	// Response
	response, _ := service.portfolioRepo.Get(rctx, portfolio.ID, portfolio.UserId, commonType.UserTypeUser)
	return responses.SuccessResponse("Portfolio record added!", response)
}

func (service *Portfolio) Update(rctx *requests.RequestContext, id, userId uint, payload requests.PortfolioRequest) responses.ServiceResponse {

	// Verify avatar
	if err := service.avatarRepo.AvatarByTypeExists(rctx, payload.AvatarId, commonType.AvatarTypePortfolio); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatar not found!", errors.New("avatar not found"))
		}

		service.container.Logger.Error("portfolio.service.update-AvatarByTypeExists", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", commonType.AvatarTypePortfolio)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later", err)
	}

	if err := service.portfolioRepo.Update(rctx, id, userId, payload); err != nil {
		service.container.Logger.Error("portfolio.service.update-Update", "error", err.Error(), "id", id, "userId", userId, "payload", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later", err)
	}

	// Response
	response, _ := service.portfolioRepo.Get(rctx, id, userId, commonType.UserTypeUser)
	return responses.SuccessResponse("Portfolio record updated!", response)
}

func (service *Portfolio) Delete(rctx *requests.RequestContext, id, userId uint) responses.ServiceResponse {

	if err := service.portfolioRepo.Delete(rctx, id, userId); err != nil {
		service.container.Logger.Error("portfolio.service.delete-Delete", "error", err.Error(), "id", id, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later", err)
	}

	return responses.SuccessResponse("Portfolio record deleted!", nil)
}
