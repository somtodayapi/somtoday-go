package somtoday

// AuthResponse is the response Authorization routes return (from "productie.somtoday.nl/oauth2/token")
type AuthResponse struct {
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	SomtodayAPIURL string `json:"somtoday_api_url"`
	Scope          string `json:"scope"`
	SomtodayTenant string `json:"somtoday_tenant"`
	IDToken        string `json:"id_token"`
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
}
