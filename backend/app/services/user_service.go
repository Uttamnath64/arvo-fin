package services

import (
	"errors"

	"github.com/Uttamnath64/arvo-fin/app/common"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type User struct {
	container    *storage.Container
	repoUser     repository.UserRepository
	repoAvatar   repository.AvatarRepository
	repoCurrency repository.CurrencyRepository
}

func NewUser(container *storage.Container) *User {
	return &User{
		container:    container,
		repoUser:     repository.NewUser(container),
		repoAvatar:   repository.NewAvatar(container),
		repoCurrency: repository.NewCurrency(container),
	}
}

func (service *User) Get(userId uint) responses.ServiceResponse {
	response, err := service.repoUser.Get(userId)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "User not found!",
			Error:      err,
		}
	}

	if err != nil {
		service.container.Logger.Error("auth.service.user-Get", err.Error(), userId)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "User records found!",
		Data:       response,
	}
}

func (service *User) GetSettings(userId uint) responses.ServiceResponse {
	response, err := service.repoUser.GetSettings(userId)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "User not found!",
			Error:      err,
		}
	}

	if err != nil {
		service.container.Logger.Error("auth.service.user-GetSettings", err.Error(), userId)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "User records found!",
		Data:       response,
	}
}

func (service *User) Update(payload requests.MeRequest, userId uint) responses.ServiceResponse {

	ok, err := service.repoAvatar.AvatarByTypeExists(payload.AvatarId, commonType.AvatarTypeUser)
	if !ok {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Avatar not found!",
			Error:      errors.New("Avatar not found!"),
		}
	}
	if err != nil {
		service.container.Logger.Error("auth.service.user-Update", err.Error(), payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later!",
			Error:      err,
		}
	}

	err = service.repoUser.Update(userId, payload)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "User not found!",
			Error:      err,
		}
	}

	if err != nil {
		service.container.Logger.Error("auth.service.user-GetSettings", err.Error(), userId)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "User records updated!",
	}
}

func (service *User) UpdateSettings(payload requests.SettingsRequest, userId uint) responses.ServiceResponse {

	ok, err := service.repoCurrency.CodeExists(payload.CurrencyCode)
	if !ok || err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Currency not found!",
			Error:      err,
		}
	}

	err = service.repoUser.UpdateSettings(userId, payload)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "User not found!",
			Error:      err,
		}
	}

	if err != nil {
		service.container.Logger.Error("auth.service.user-GetSettings", err.Error(), userId)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "User setting updated!",
	}
}
