package services

import (
	"errors"

	"github.com/Uttamnath64/arvo-fin/app/common"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type Category struct {
	container     *storage.Container
	repoCategory  repository.CategoryRepository
	repoAvatar    repository.AvatarRepository
	repoPortfolio repository.PortfolioRepository
}

func NewCategory(container *storage.Container) *Category {
	return &Category{
		container:     container,
		repoCategory:  repository.NewCategory(container),
		repoAvatar:    repository.NewAvatar(container),
		repoPortfolio: repository.NewPortfolio(container),
	}
}

func (service *Category) Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse {

	response, err := service.repoCategory.Get(rctx, id, rctx.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "We couldn’t find the category you were looking for.", err)
		}

		service.container.Logger.Error("category.appService.get-Get", "error", err.Error(), "id", id)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("Category details retrieved successfully.", response)
}
func (service *Category) GetList(rctx *requests.RequestContext, portfolioId, userId uint) responses.ServiceResponse {
	response, err := service.repoCategory.GetList(rctx, portfolioId, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "No category were found in this portfolio.", err)
		}

		service.container.Logger.Error("category.appService.getList-GetList", "error", err.Error(), "portfolioId", portfolioId, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("Categories details retrieved successfully.", response)
}

func (service *Category) Create(rctx *requests.RequestContext, payload requests.CategoryRequest) responses.ServiceResponse {
	sourceType := commonType.UserTypeAdmin

	if rctx.UserType == commonType.UserTypeUser {
		// Check portfolio
		if err := service.repoPortfolio.UserPortfolioExists(rctx, payload.PortfolioId, rctx.UserID); err != nil {
			if err == gorm.ErrRecordNotFound {
				return responses.ErrorResponse(common.StatusNotFound, "No portfolio found for this user.", errors.New("portfolio not found"))
			}
			service.container.Logger.Error("category.appService.create-UserPortfolioExists", "error", err.Error(), "payload", payload, "userId", rctx.UserID)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
		}

		// Set source type as user
		sourceType = commonType.UserTypeUser
	}

	// Check avatar
	if err := service.repoAvatar.AvatarByTypeExists(rctx, payload.AvatarId, commonType.AvatarTypeCategory); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Selected avatar not found. Please choose a valid one.", errors.New("avatar not found"))
		}

		service.container.Logger.Error("category.appService.create-AvatarByTypeExists", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", commonType.AvatarTypePortfolio)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	id, err := service.repoCategory.Create(rctx, models.Category{
		SourceId:    rctx.UserID,
		SourceType:  sourceType,
		PortfolioId: &payload.PortfolioId,
		AvatarId:    payload.AvatarId,
		Name:        payload.Name,
		Type:        payload.Type,
	})
	if err != nil {
		service.container.Logger.Error("category.appService.create-Create", "error", err.Error(), "payload", payload, "userId", rctx.UserID)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	response, _ := service.repoCategory.Get(rctx, id, rctx.UserID)
	return responses.SuccessResponse("Your category has been created successfully. 🎉", response)
}

func (service *Category) Update(rctx *requests.RequestContext, id uint, payload requests.CategoryRequest) responses.ServiceResponse {
	response, err := service.repoCategory.Get(rctx, id, rctx.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "We couldn’t find the category you were looking for.", err)
		}

		service.container.Logger.Error("category.appService.update-Get", "error", err.Error(), "id", id)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Check if user then check portfolio
	if rctx.UserType == commonType.UserTypeUser {
		if err := service.repoPortfolio.UserPortfolioExists(rctx, payload.PortfolioId, rctx.UserID); err != nil {
			if err == gorm.ErrRecordNotFound {
				return responses.ErrorResponse(common.StatusNotFound, "No portfolio found for this user.", errors.New("portfolio not found"))
			}

			service.container.Logger.Error("category.appService.create-UserPortfolioExists", "error", err.Error(), "payload", payload, "userId", rctx.UserID)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
		}
	}

	// Check avatar
	if err := service.repoAvatar.AvatarByTypeExists(rctx, payload.AvatarId, commonType.AvatarTypeCategory); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Selected avatar not found. Please choose a valid one.", errors.New("avatar not found"))
		}

		service.container.Logger.Error("category.appService.create-AvatarByTypeExists", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", commonType.AvatarTypePortfolio)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// User want to update shared category
	if response.SourceType == commonType.UserTypeAdmin && rctx.UserType != commonType.UserTypeAdmin {
		id, err = service.repoCategory.Create(rctx, models.Category{
			SourceId:     rctx.UserID,
			SourceType:   commonType.UserTypeUser,
			PortfolioId:  &payload.PortfolioId,
			AvatarId:     payload.AvatarId,
			Name:         payload.Name,
			Type:         payload.Type,
			CopiedFromId: &response.ID,
		})
		if err != nil {
			service.container.Logger.Error("category.appService.update-Create", "error", err.Error(), "payload", payload, "userId", rctx.UserID)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
		}
	} else {
		if err := service.repoCategory.Update(rctx, id, rctx.UserID, payload); err != nil {
			if err == gorm.ErrRecordNotFound {
				return responses.ErrorResponse(common.StatusNotFound, "Unable to update the account. Please try again.", errors.New("account not updated"))
			}
			service.container.Logger.Error("category.appService.update-Update", "error", err.Error(), "id", id, "userId", rctx.UserID, "id", payload)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
		}
	}

	response, _ = service.repoCategory.Get(rctx, id, rctx.UserID)
	return responses.SuccessResponse("Your category details have been updated successfully.", response)
}

func (service *Category) Delete(rctx *requests.RequestContext, portfolioId, id uint) responses.ServiceResponse {

	response, err := service.repoCategory.Get(rctx, id, rctx.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "We couldn’t find the category you were looking for.", err)
		}

		service.container.Logger.Error("category.appService.deleted-Get", "error", err.Error(), "id", id)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// User want to update shared category
	if response.SourceType == commonType.UserTypeAdmin && rctx.UserType != commonType.UserTypeAdmin {
		id, err = service.repoCategory.Create(rctx, models.Category{
			SourceId:     rctx.UserID,
			SourceType:   commonType.UserTypeUser,
			PortfolioId:  &portfolioId,
			AvatarId:     response.AvatarId,
			Name:         response.Name,
			Type:         response.Type,
			CopiedFromId: &response.ID,
		})
		if err != nil {
			service.container.Logger.Error("category.appService.deleted-Create", "error", err.Error(), "userId", rctx.UserID)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
		}
	}

	if err := service.repoCategory.Delete(rctx, id, rctx.UserID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "category deletion failed. Please try again.", errors.New("category not deleted"))
		}
		service.container.Logger.Error("category.appService.deleted-Delete", "error", err.Error(), "id", id, "userId", rctx.UserID)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	return responses.SuccessResponse("The category has been deleted successfully.", nil)
}
