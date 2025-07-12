package query

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/gin-gonic/gin"
)

const (
	LayoutDate       = "2006-01-02"
	LayoutDateTime   = "2006-01-02 15:04:05"
	FriendlyDate     = "YYYY-MM-DD"
	FriendlyDateTime = "YYYY-MM-DD HH:MM:SS"
)

func QId(c *gin.Context, query string, required bool) (uint, bool) {
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
	} else if required {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: query + " query param is required!",
		})
		return 0, false
	}
	return valInt, true
}

// get query param date
func QDate(c *gin.Context, key string, required bool) (time.Time, bool) {
	return qTime(c, key, required, LayoutDate, FriendlyDate)
}

// get query param date range
func QDateRange(c *gin.Context, from, to string, required bool) (time.Time, time.Time, bool) {
	return qTimeRange(c, from, to, required, LayoutDate, FriendlyDate)
}

// get query param date time
func QDateTime(c *gin.Context, key string, required bool) (time.Time, bool) {
	return qTime(c, key, required, LayoutDateTime, FriendlyDateTime)
}

// get query param date time range
func QDateTimeRange(c *gin.Context, from, to string, required bool) (time.Time, time.Time, bool) {
	return qTimeRange(c, from, to, required, LayoutDateTime, FriendlyDateTime)
}

func qTime(c *gin.Context, key string, required bool, layout, friendly string) (time.Time, bool) {
	dateStr := c.Query(key)

	if required && dateStr == "" {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: fmt.Sprintf("'%s' is required and must be in %s format", key, friendly),
		})
		return time.Time{}, false
	}

	if dateStr == "" {
		return time.Time{}, true
	}

	parsed, err := time.Parse(layout, dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: fmt.Sprintf("Invalid '%s' format. Use %s", key, friendly),
		})
		return time.Time{}, false
	}

	return parsed, true
}

func qTimeRange(c *gin.Context, from, to string, required bool, layout, friendly string) (time.Time, time.Time, bool) {
	timeFrom, ok := qTime(c, from, required, layout, friendly)
	if !ok {
		return time.Time{}, time.Time{}, false
	}

	timeTo, ok := qTime(c, to, required, layout, friendly)
	if !ok {
		return time.Time{}, time.Time{}, false
	}

	// Check dateFrom <= dateTo (if both exist)
	if !timeFrom.IsZero() && !timeTo.IsZero() && timeFrom.After(timeTo) {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: fmt.Sprintf("'%s' cannot be after '%s'", from, to),
		})
		return time.Time{}, time.Time{}, false
	}

	return timeFrom, timeTo, true
}
