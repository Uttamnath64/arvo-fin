package responses

type MeResponse struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Email      string `json:"emil"`
	AvatarID   uint   `json:"avatar_id"`
	AvatarIcon string `json:"avatar_icon"`
}

type SettingsResponse struct {
	Id                 uint   `json:"id"`
	DecimalPlaces      int    `json:"decimal_places"`
	NumberFormat       int    `json:"number_format"`
	EmailNotifications bool   `json:"email_notifications"`
	CurrencyCode       string `json:"currency_code"`
	CurrencySymbol     string `json:"currency_symbol"`
	CurrencyName       string `json:"currency_name"`
}
