package responses

type PortfolioResponse struct {
	Id     uint           `json:"id"`
	Name   string         `json:"name"`
	Avatar AvatarResponse `json:"avatar"`
}
