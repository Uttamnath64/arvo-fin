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

func (service *Portfolio) GetList(userId uint, userType commonType.UserType) responses.ServiceResponse {
	return service.portfolioService.GetList(userId, userType)
}

func (service *Portfolio) Get(id, userId uint, userType commonType.UserType) responses.ServiceResponse {
	return service.portfolioService.Get(id, userId, userType)
}

func (service *Portfolio) Add(payload requests.PortfolioRequest, userId uint) responses.ServiceResponse {

	portfolio := models.Portfolio{
		Name:     payload.Name,
		AvatarId: payload.AvatarId,
		UserId:   userId,
	}

	if err := service.avatarRepo.GetAvatarByType(payload.AvatarId, commonType.PortfolioAvatar, &models.Avatar{}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.ServiceResponse{
				StatusCode: common.StatusValidationError,
				Message:    "Avatar not found!",
				Error:      err,
			}
		}

		// Other database errors
		service.container.Logger.Error("portfolio.service.add", err.Error(), portfolio, userId)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later!",
			Error:      err,
		}
	}

	err := service.portfolioRepo.Add(portfolio)
	if err != nil {
		service.container.Logger.Error("portfolio.service.add", err.Error(), portfolio, userId)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Portfolio record added!",
	}
}

func (service *Portfolio) Update(id, userId uint, payload requests.PortfolioRequest) responses.ServiceResponse {

	if err := service.avatarRepo.GetAvatarByType(payload.AvatarId, commonType.PortfolioAvatar, &models.Avatar{}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.ServiceResponse{
				StatusCode: common.StatusValidationError,
				Message:    "Avatar not found!",
				Error:      err,
			}
		}

		// Other database errors
		service.container.Logger.Error("portfolio.service.update", err.Error(), id, userId, payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later!",
			Error:      err,
		}
	}

	err := service.portfolioRepo.Update(id, userId, payload)
	if err != nil {
		service.container.Logger.Error("portfolio.service.update", err.Error(), id, userId, payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    err.Error(),
			Error:      err,
		}
	}

	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Portfolio record updated!",
	}
}

func (service *Portfolio) Delete(id, userId uint) responses.ServiceResponse {

	err := service.portfolioRepo.Delete(id, userId)
	if err != nil {
		service.container.Logger.Error("portfolio.service.delete", err.Error(), id, userId)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    err.Error(),
			Error:      err,
		}
	}

	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Portfolio record deleted!",
	}
}
