package services

import (
	"github.com/Uttamnath64/arvo-fin/app/common"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type Avatar struct {
	container  *storage.Container
	repoAvatar repository.AvatarRepository
}

func NewAvatar(container *storage.Container) *Avatar {
	return &Avatar{
		container:  container,
		repoAvatar: repository.NewAvatar(container),
	}
}

func (service *Avatar) Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse {

	response, err := service.repoAvatar.Get(rctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatar not found. Please choose a valid one.", err)
		}

		service.container.Logger.Error("avatar.appService.get-Get", "error", err.Error(), "id", id)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("Avatar details retrieved successfully.", response)
}

func (service *Avatar) GetAvatarsByType(rctx *requests.RequestContext, avatarType commonType.AvatarType) responses.ServiceResponse {
	response, err := service.repoAvatar.GetAvatarsByType(rctx, avatarType)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatars not found by type!", err)
		}

		service.container.Logger.Error("avatar.appService.getAvatarsByType-GetAvatarsByType", "error", err.Error(), "avatarType", avatarType)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("Avatars retrieved successfully by type.", response)
}

func (service *Avatar) Create(rctx *requests.RequestContext, payload requests.AvatarRequest) responses.ServiceResponse {

	avatarId, err := service.repoAvatar.Create(rctx, models.Avatar{
		Name: payload.Name,
		Type: payload.Type,
		Icon: payload.Icon,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatar not found. Please choose a valid one.", err)
		}

		service.container.Logger.Error("avatar.appService.creatre-Creatre", "error", err.Error(), "payload", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	response, _ := service.repoAvatar.Get(rctx, avatarId)
	return responses.SuccessResponse("Success! The avatar was created.", response)
}

func (service *Avatar) Update(rctx *requests.RequestContext, id uint, payload requests.AvatarRequest) responses.ServiceResponse {

	err := service.repoAvatar.Update(rctx, id, payload)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatar not found. Please choose a valid one.", err)
		}

		service.container.Logger.Error("avatar.appService.update-Update", "error", err.Error(), "id", id, "payload", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	response, _ := service.repoAvatar.Get(rctx, id)
	return responses.SuccessResponse("Avatar information saved successfully.", response)
}
