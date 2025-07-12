package handlers

import (
	"net/http"

	"github.com/Uttamnath64/arvo-fin/app/common"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/gin-gonic/gin"
)

func isErrorResponse(c *gin.Context, serviceResponse responses.ServiceResponse) bool {

	if serviceResponse.HasError() {
		apiResponse := responses.ApiResponse{
			Status:  false,
			Message: serviceResponse.Message,
			Details: serviceResponse.Error.Error(),
		}

		switch serviceResponse.StatusCode {
		case common.StatusNotFound:
			c.JSON(http.StatusNotFound, apiResponse)
		case common.StatusValidationError:
			c.JSON(http.StatusUnauthorized, apiResponse)
		case common.StatusServerError:
			c.JSON(http.StatusInternalServerError, apiResponse)
		}
		return true
	}
	return false
}

func bindAndValidateJson[T any](c *gin.Context, payload *T) bool {
	if err := c.ShouldBindJSON(payload); err != nil {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid request payload. Please check the input format.",
			Details: err.Error(),
		})
		return false
	}

	// Validate only if the struct has an IsValid() method
	if validatable, ok := interface{}(payload).(interface{ IsValid() error }); ok {
		if err := validatable.IsValid(); err != nil {
			c.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: err.Error(),
			})
			return false
		}
	}
	return true
}

func getRequestContext(c *gin.Context) (*requests.RequestContext, bool) {
	val, exists := c.Get("rctx")
	if !exists {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Authorization token is missing.",
		})
		return nil, false
	}
	rctx, ok := val.(*requests.RequestContext)
	if !ok {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid authorization token.",
		})
		return nil, false
	}
	return rctx, true
}
