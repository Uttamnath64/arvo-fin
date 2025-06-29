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

func (service *User) Get(rctx *requests.RequestContext, userId uint) responses.ServiceResponse {

	response, err := service.repoUser.Get(rctx, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Oops! That user doesn’t exist.", err)
		}
		service.container.Logger.Error("user.appService.get-Get", "error", err.Error(), "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("User details retrieved successfully!", response)
}

func (service *User) GetSettings(rctx *requests.RequestContext, userId uint) responses.ServiceResponse {
	response, err := service.repoUser.GetSettings(rctx, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Oops! That user doesn’t exist.", err)
		}

		service.container.Logger.Error("user.appService.getSettings-GetSettings", "error", err.Error(), "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("User settings retrieved successfully!", response)
}

func (service *User) Update(rctx *requests.RequestContext, payload requests.MeRequest, userId uint) responses.ServiceResponse {

	// Check avatar
	if err := service.repoAvatar.AvatarByTypeExists(rctx, payload.AvatarId, commonType.AvatarTypeUser); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Selected avatar not found. Please choose a valid one.", errors.New("avatar not found"))
		}
		service.container.Logger.Error("user.appService.update-GetSettings", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", commonType.AvatarTypeUser)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// update
	if err := service.repoUser.Update(rctx, userId, payload); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Oops! That user doesn’t exist.", err)
		}

		service.container.Logger.Error("user.appService.update-GetSettings", "error", err.Error(), "userId", userId, "payload", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("User profile has been updated!", nil)
}

func (service *User) UpdateSettings(rctx *requests.RequestContext, payload requests.SettingsRequest, userId uint) responses.ServiceResponse {

	// Check currencyCode
	if err := service.repoCurrency.CodeExists(rctx, payload.CurrencyCode); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "The specified currency code is invalid or not supported.", err)
		}

		service.container.Logger.Error("user.appService.updateSettings-CodeExists", "error", err.Error(), "userId", userId, "currencyCode", payload.CurrencyCode)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Update
	if err := service.repoUser.UpdateSettings(rctx, userId, payload); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Oops! That user doesn’t exist.!", err)
		}

		service.container.Logger.Error("user.appService.updateSettings-UpdateSettings", "error", err.Error(), "userId", userId, "payload", payload)
		return responses.ErrorResponse(common.StatusServerError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("Your settings have been updated!", nil)
}
