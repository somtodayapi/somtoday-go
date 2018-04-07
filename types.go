package somtoday

type AuthResponse struct {
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	somtodayApiUrl string `json:"somtoday_api_url"`
	Scope          string `json:"scope"`
	SomtodayTenant string `json:"somtoday_tenant"`
	IdToken        string `json:"id_token"`
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
}
