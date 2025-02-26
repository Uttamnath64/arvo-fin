package responses

type MeResponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	AvatarID  uint   `json:"avatar_id"`
	AvatarURL string `json:"url"`
}

type SettingsResponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	AvatarID  uint   `json:"avatar_id"`
	AvatarURL string `json:"url"`
}
