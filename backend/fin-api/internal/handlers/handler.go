package handlers

import (
	"net/http"

	"github.com/Uttamnath64/arvo-fin/app/common"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/gin-gonic/gin"
)

type userInfo struct {
	userId   uint
	userType commonType.UserType
}

func isErrorResponse(ctx *gin.Context, serviceResponse responses.ServiceResponse) bool {

	if serviceResponse.HasError() {
		apiResponse := responses.ApiResponse{
			Status:  false,
			Message: serviceResponse.Message,
			Details: serviceResponse.Error.Error(),
		}

		switch serviceResponse.StatusCode {
		case common.StatusNotFound:
			ctx.JSON(http.StatusBadRequest, apiResponse)
			break
		case common.StatusValidationError:
			ctx.JSON(http.StatusUnauthorized, apiResponse)
			break
		case common.StatusServerError:
			ctx.JSON(http.StatusInternalServerError, apiResponse)
			break
		}
		return true
	}
	return false
}

func bindAndValidateJson[T any](ctx *gin.Context, payload *T) bool {
	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid request payload. Please check the input data format!",
			Details: err.Error(),
		})
		return false
	}

	// Validate only if the struct has an IsValid() method
	if validatable, ok := interface{}(payload).(interface{ IsValid() error }); ok {
		if err := validatable.IsValid(); err != nil {
			ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: err.Error(),
			})
			return false
		}
	}
	return true
}

func getUserInfo(ctx *gin.Context) (*userInfo, bool) {
	userId, exists := ctx.Get("userId")
	userType, existsType := ctx.Get("userType")

	// Check if values exist
	if !exists || !existsType {
		ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
		})
		return nil, false
	}
	return &userInfo{
		userId:   userId.(uint),
		userType: userType.(commonType.UserType),
	}, true
}
