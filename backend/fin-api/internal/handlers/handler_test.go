package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Uttamnath64/arvo-fin/app/common"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIsErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name            string
		serviceResponse responses.ServiceResponse
		expectedStatus  int
		expectedBody    string
		expectError     bool
	}{
		{
			name: "Not Found Error",
			serviceResponse: responses.ServiceResponse{
				Error:      errors.New("not found"),
				Message:    "Resource not found",
				StatusCode: common.StatusNotFound,
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"status":false,"message":"Resource not found","metadata":null,"details":"not found"}`,
			expectError:    true,
		},
		{
			name: "Validation Error",
			serviceResponse: responses.ServiceResponse{
				Error:      errors.New("invalid input"),
				Message:    "Validation failed",
				StatusCode: common.StatusValidationError,
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"status":false,"message":"Validation failed","metadata":null,"details":"invalid input"}`,
			expectError:    true,
		},
		{
			name: "Server Error",
			serviceResponse: responses.ServiceResponse{
				Error:      errors.New("internal issue"),
				Message:    "Server crashed",
				StatusCode: common.StatusServerError,
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"status":false,"message":"Server crashed","metadata":null,"details":"internal issue"}`,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ok := isErrorResponse(ctx, tt.serviceResponse)
			assert.Equal(t, tt.expectError, ok)

			if tt.expectError {
				assert.Equal(t, tt.expectedStatus, w.Code)
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			} else {
				assert.Equal(t, 0, w.Code)
				assert.Empty(t, w.Body.String())
			}
		})
	}
}

func TestBindAndValidateJson(t *testing.T) {
	tests := []struct {
		name        string
		body        interface{}
		expectValid bool
	}{
		{
			name: "Valid LoginRequest",
			body: requests.LoginRequest{
				UsernameEmail: "uttam@example.com",
				Password:      "Secret@123",
			},
			expectValid: true,
		},
		{
			name: "InValid LoginRequest",
			body: requests.LoginRequest{
				UsernameEmail: "uttamexample.com",
				Password:      "Secret",
			},
			expectValid: false,
		},
		{
			name: "Valid RegisterRequest",
			body: requests.RegisterRequest{
				Name:     "uttam nath",
				Email:    "uttam@example.com",
				Username: "uttam.nath",
				AvatarId: 12,
				Password: "Secret@123",
				OTP:      "123443",
			},
			expectValid: true,
		},
		{
			name: "InValid RegisterRequest",
			body: requests.RegisterRequest{
				Name:     "uttam nath",
				Email:    "uttam!example.com",
				Username: "uttam.nath",
				AvatarId: 12,
				Password: "Secret@123",
				OTP:      "1234",
			},
			expectValid: false,
		},
		{
			name: "Invalid Password",
			body: requests.LoginRequest{
				UsernameEmail: "uttam@example.com",
				Password:      "1234567890",
			},
			expectValid: false,
		},
		{
			name: "Invalid UsernameEmail",
			body: requests.LoginRequest{
				UsernameEmail: "uttam-example",
				Password:      "Secret@123",
			},
			expectValid: false,
		},
	}

	requests.NewResponse()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			body, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			switch tt.body.(type) {
			case requests.LoginRequest:
				data := requests.LoginRequest{}
				ok := bindAndValidateJson(ctx, &data)
				assert.Equal(t, tt.expectValid, ok)
			case requests.RegisterRequest:
				data := requests.RegisterRequest{}
				ok := bindAndValidateJson(ctx, &data)
				assert.Equal(t, tt.expectValid, ok)
			default:
				t.Fatalf("unsupported type for test: %T", tt.body)
			}
		})
	}
}

func TestGetUserInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		userId     interface{}
		userType   interface{}
		exUserId   uint
		exUserType commonType.UserType
		isValid    bool
	}{
		{
			name:       "Invalid UserId",
			userId:     -1,
			userType:   1,
			exUserId:   uint(0),
			exUserType: commonType.UserTypeUser,
			isValid:    false,
		},
		{
			name:       "Invalid UserType",
			userId:     uint(10),
			userType:   10,
			exUserId:   uint(10),
			exUserType: commonType.UserTypeAdmin,
			isValid:    false,
		},
		{
			name:       "Valid User",
			userId:     uint(10),
			userType:   1,
			exUserId:   uint(10),
			exUserType: commonType.UserTypeUser,
			isValid:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Set("user_id", tt.userId)
			ctx.Set("user_type", tt.userType)

			user, ok := getUserInfo(ctx)
			assert.Equal(t, tt.isValid, ok && user.userType.IsValid())
			if tt.isValid {
				assert.Equal(t, tt.exUserId, user.userId)
				assert.Equal(t, tt.exUserType, user.userType)
			}

		})
	}
}
