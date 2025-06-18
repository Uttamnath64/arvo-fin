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

	var avatar models.Avatar
	err := service.repoAvatar.Get(id, &avatar)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Avatar not found!",
			Error:      err,
		}
	}

	if err != nil {
		service.container.Logger.Error("auth.service.avatar-Get", err.Error(), id)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Avatar records found!",
		Data:       avatar,
	}
}

func (service *Avatar) GetAvatarsByType(avatarType commonType.AvatarType) responses.ServiceResponse {
	response, err := service.repoAvatar.GetAvatarsByType(avatarType)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Avatars not found by type!",
			Error:      err,
		}
	}

	if err != nil {
		service.container.Logger.Error("auth.service.avatar-GetByType", err.Error(), avatarType)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Avatars found by type!",
		Data:       response,
	}
}

func (service *Avatar) Creatre(payload requests.AvatarRequest) responses.ServiceResponse {

	err := service.repoAvatar.Create(models.Avatar{
		Name: payload.Name,
		Type: payload.Type,
		Icon: payload.Icon,
	})
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Avatar not found!",
			Error:      err,
		}
	}

	if err != nil {
		service.container.Logger.Error("auth.service.avater-Creatre", err.Error())
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Avatar is created!",
	}
}

func (service *Avatar) Update(id uint, payload requests.AvatarRequest) responses.ServiceResponse {

	err := service.repoAvatar.Update(id, payload)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Avatar not found!",
			Error:      err,
		}
	}

	if err != nil {
		service.container.Logger.Error("auth.service.avatar-Update", err.Error(), id)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Avatar is updated!",
	}
}
