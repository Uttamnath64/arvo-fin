package responses

type PortfolioResponse struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	AvatarID   uint   `json:"avatar_id"`
	AvatarIcon string `json:"avatar_icon"`
}
