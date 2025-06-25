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
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "User not found!", err)
		}
		service.container.Logger.Error("user.appService.get-Get", "error", err.Error(), "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("User records found!", response)
}

func (service *User) GetSettings(userId uint) responses.ServiceResponse {
	response, err := service.repoUser.GetSettings(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "User not found!", err)
		}

		service.container.Logger.Error("user.appService.getSettings-GetSettings", "error", err.Error(), "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("User settings found!", response)
}

func (service *User) Update(payload requests.MeRequest, userId uint) responses.ServiceResponse {

	// Check avatar
	if err := service.repoAvatar.AvatarByTypeExists(payload.AvatarId, commonType.AvatarTypeUser); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatar not found!", errors.New("avatar not found"))
		}
		service.container.Logger.Error("user.appService.update-GetSettings", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", commonType.AvatarTypeUser, "avatarTypeName", commonType.AvatarTypeUser.String())
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// update
	if err := service.repoUser.Update(userId, payload); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "User not found!", err)
		}

		service.container.Logger.Error("user.appService.update-GetSettings", "error", err.Error(), "userId", userId, "payload", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("User records updated!", nil)
}

func (service *User) UpdateSettings(payload requests.SettingsRequest, userId uint) responses.ServiceResponse {

	// Check currencyCode
	if err := service.repoCurrency.CodeExists(payload.CurrencyCode); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Currency not found!", err)
		}

		service.container.Logger.Error("user.appService.updateSettings-CodeExists", "error", err.Error(), "userId", userId, "currencyCode", payload.CurrencyCode)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Update
	if err := service.repoUser.UpdateSettings(userId, payload); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "User not found!", err)
		}

		service.container.Logger.Error("user.appService.updateSettings-UpdateSettings", "error", err.Error(), "userId", userId, "payload", payload)
		return responses.ErrorResponse(common.StatusServerError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("User setting updated!", nil)
}
