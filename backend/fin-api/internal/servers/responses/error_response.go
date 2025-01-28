package response

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Details string `json:"details"`
}
