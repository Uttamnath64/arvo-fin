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

func (service *Avatar) Get(id uint) responses.ServiceResponse {

	response, err := service.repoAvatar.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatar not found!", err)
		}

		service.container.Logger.Error("avatar.appService.get-Get", "error", err.Error(), "id", id)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("Avatar records found!", response)
}

func (service *Avatar) GetAvatarsByType(avatarType commonType.AvatarType) responses.ServiceResponse {
	response, err := service.repoAvatar.GetAvatarsByType(avatarType)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatars not found by type!", err)
		}

		service.container.Logger.Error("avatar.appService.getAvatarsByType-GetAvatarsByType", "error", err.Error(), "avatarType", avatarType, "avatarTypeName", avatarType.String())
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("Avatars found by type!", response)
}

func (service *Avatar) Creatre(payload requests.AvatarRequest) responses.ServiceResponse {

	avatarId, err := service.repoAvatar.Create(models.Avatar{
		Name: payload.Name,
		Type: payload.Type,
		Icon: payload.Icon,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatar not found!", err)
		}

		service.container.Logger.Error("avatar.appService.creatre-Creatre", "error", err.Error(), "payload", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	response, _ := service.repoAvatar.Get(avatarId)
	return responses.SuccessResponse("Avatar is created!", response)
}

func (service *Avatar) Update(id uint, payload requests.AvatarRequest) responses.ServiceResponse {

	err := service.repoAvatar.Update(id, payload)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatar not found!", err)
		}

		service.container.Logger.Error("avatar.appService.update-Update", "error", err.Error(), "id", id, "payload", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	response, _ := service.repoAvatar.Get(id)
	return responses.SuccessResponse("Avatar is updated!", response)
}
