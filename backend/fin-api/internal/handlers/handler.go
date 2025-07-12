package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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

func getQueryParamId(c *gin.Context, query string, required bool) (uint, bool) {
	var valInt uint = 0
	valStr := c.Query(query)
	if valStr != "" {
		val, err := strconv.Atoi(valStr)
		if err != nil || val <= 0 {
			c.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: "Invalid " + query + "!",
			})
			return 0, false
		}
		valInt = uint(val)
	} else if required == true {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: query + " query param is required!",
		})
		return 0, false
	}
	return valInt, true
}

func getQueryParamDate(c *gin.Context, key string, required bool) (time.Time, bool) {
	layout := "2006-01-02"
	dateStr := c.Query(key)

	if required && dateStr == "" {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: fmt.Sprintf("'%s' is required and must be in YYYY-MM-DD format", key),
		})
		return time.Time{}, false
	}

	if dateStr == "" {
		return time.Time{}, true // Empty is allowed
	}

	parsed, err := time.Parse(layout, dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: fmt.Sprintf("'%s' is required and must be in YYYY-MM-DD format", key),
		})
		return time.Time{}, false
	}

	return parsed, true
}

func getQueryParamDateRange(c *gin.Context, from, to string, required bool) (time.Time, time.Time, bool) {
	dateFrom, ok := getQueryParamDate(c, from, required)
	if !ok {
		return time.Time{}, time.Time{}, false
	}

	dateTo, ok := getQueryParamDate(c, to, required)
	if !ok {
		return time.Time{}, time.Time{}, false
	}

	// Check dateFrom <= dateTo (if both exist)
	if !dateFrom.IsZero() && !dateTo.IsZero() && dateFrom.After(dateTo) {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: fmt.Sprintf("'%s' cannot be after '%s'", from, to),
		})
		return time.Time{}, time.Time{}, false
	}

	return dateFrom, dateTo, true
}

func getQueryParamDateTime(c *gin.Context, key string, required bool) (time.Time, bool) {
	layout := "2006-01-02"
	dateStr := c.Query(key)

	if required && dateStr == "" {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: fmt.Sprintf("'%s' is required and must be in YYYY-MM-DD HH:MM:SS format", key),
		})
		return time.Time{}, false
	}

	if dateStr == "" {
		return time.Time{}, true // Empty is allowed
	}

	parsed, err := time.Parse(layout, dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: fmt.Sprintf("'%s' is required and must be in YYYY-MM-DD HH:MM:SS format", key),
		})
		return time.Time{}, false
	}

	return parsed, true
}

func getQueryParamDateTimeRange(c *gin.Context, from, to string, required bool) (time.Time, time.Time, bool) {
	dateFrom, ok := getQueryParamDate(c, from, required)
	if !ok {
		return time.Time{}, time.Time{}, false
	}

	dateTo, ok := getQueryParamDate(c, to, required)
	if !ok {
		return time.Time{}, time.Time{}, false
	}

	// Check dateFrom <= dateTo (if both exist)
	if !dateFrom.IsZero() && !dateTo.IsZero() && dateFrom.After(dateTo) {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: fmt.Sprintf("'%s' cannot be after '%s'", from, to),
		})
		return time.Time{}, time.Time{}, false
	}

	return dateFrom, dateTo, true
}
