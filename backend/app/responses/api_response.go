package responses

type ApiResponse struct {
	Status   bool        `json:"status"`
	Message  string      `json:"message"`
	Details  string      `json:"details"`
	Metadata interface{} `json:"metadata"`
}
