package somtoday

// OrganisationResponse is how the server returns organisations
type OrganisationResponse struct {
	Instellingen []Organisation `json:"instellingen"`
}

// An Organisation is an organisation (from "https://servers.somtoday.nl/organisaties.json")
type Organisation struct {
	Naam   string `json:"naam"`
	Plaats string `json:"plaats"`
	UUID   string `json:"uuid"`
}

// AuthResponse is the response Authorization routes return (from "https://productie.somtoday.nl/oauth2/token")
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
