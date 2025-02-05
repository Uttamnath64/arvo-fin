package responses

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
