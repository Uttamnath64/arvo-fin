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
	portfolioRepo *repository.Portfolio
	avatarRepo    *repository.Avatar
}

func NewPortfolio(container *storage.Container) *Portfolio {
	return &Portfolio{
		container:     container,
		portfolioRepo: repository.NewPortfolio(container),
		avatarRepo:    repository.NewAvatar(container),
	}
}

func (service *Portfolio) GetList(userId uint, userType commonType.UserType) responses.ServiceResponse {

	portfolioRes, err := service.portfolioRepo.GetList(userId, userType)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Record not found!",
			Error:      err,
		}
	}
	if err != nil {
		service.container.Logger.Error("portfolio.service.getList", err.Error(), userId, userType)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Portfolio records found!",
		Data:       portfolioRes,
	}
}

func (service *Portfolio) Get(id, userId uint, userType commonType.UserType) responses.ServiceResponse {

	portfolioRes, err := service.portfolioRepo.Get(id, userId, userType)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Record not found!",
			Error:      err,
		}
	}
	if err != nil {
		service.container.Logger.Error("portfolio.service.get", err.Error(), id, userId, userType)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Portfolio record found!",
		Data:       portfolioRes,
	}
}
