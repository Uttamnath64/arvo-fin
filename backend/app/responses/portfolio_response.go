package responses

type PortfolioResponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	AvatarID  uint   `json:"avatar_id"`
	AvatarURL string `json:"url"`
}
