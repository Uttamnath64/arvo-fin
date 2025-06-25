package responses

import "github.com/Uttamnath64/arvo-fin/app/common"

type ServiceResponse struct {
	StatusCode int
	Data       interface{}
	Message    string
	Error      error
}

// Helper to check if the response contains an error
func (r *ServiceResponse) HasError() bool {
	return r.Error != nil
}

func ErrorResponse(status int, message string, err error) ServiceResponse {
	return ServiceResponse{
		StatusCode: status,
		Message:    message,
		Error:      err,
	}
}

func SuccessResponse(message string, data interface{}) ServiceResponse {
	return ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    message,
		Data:       data,
	}
}
