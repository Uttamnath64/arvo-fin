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
		case common.StatusValidationError:
			ctx.JSON(http.StatusUnauthorized, apiResponse)
		case common.StatusServerError:
			ctx.JSON(http.StatusInternalServerError, apiResponse)
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
	userIdValue, exists := ctx.Get("user_id")
	userTypeValue, existsType := ctx.Get("user_type")

	if !exists || !existsType {
		ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
			Status:  false,
			Message: "Unauthorized",
		})
		ctx.Abort()
		return nil, false
	}

	userId, ok := userIdValue.(uint)
	if !ok {
		// Try to convert from float64 (common in JWT claims)
		if floatId, isFloat := userIdValue.(float64); isFloat {
			userId = uint(floatId)
		} else {
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid userId format",
			})
			ctx.Abort()
			return nil, false
		}
	}

	// Handle userType safely
	userTypeInt, ok := userTypeValue.(int)
	if !ok {
		// Handle case where it's a float64 (common in JWT claims)
		if userTypeFloat, isFloat := userTypeValue.(float64); isFloat {
			userTypeInt = int(userTypeFloat) // Convert float64 → int
		} else {
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid userType format",
			})
			ctx.Abort()
			return nil, false
		}
	}
	userType := commonType.UserType(userTypeInt)

	// Return parsed values
	return &userInfo{
		userId:   userId,
		userType: userType,
	}, true
}
