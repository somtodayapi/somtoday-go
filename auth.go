package somtoday

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Login fetches the tokens and returns an AuthResponse and error
func Login(school Organisation, username, password string) (response AuthResponse, err error) {
	// Do http request
	client := &http.Client{}
	params := url.Values{}
	params.Set("grant_type", "password")
	params.Set("username", school.UUID+"\\"+username)
	params.Set("password", password)
	params.Set("scope", "openid")
	req, err := http.NewRequest("POST", "https://production.somtoday.nl/oauth2/token", strings.NewReader(params.Encode()))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(AppID+":"+AppSecret))) // I use base64 here so if somtoday changes the AppID and secret, I only have to change it in misc.go
	req.Header.Add("Accept", "application/json")                                                             // sometimes it returns xml

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return
}

// RefreshTokens refreshes the tokens
func RefreshTokens(refreshToken string) (response AuthResponse, err error) {
	client := &http.Client{}
	params := url.Values{}
	params.Set("grant_type", "refresh_token")
	params.Set("client_id", AppID)
	params.Set("client_secret", AppSecret)
	params.Set("refresh_token", refreshToken)

	req, err := http.NewRequest("POST", "https://production.somtoday.nl/oauth2/token", strings.NewReader(params.Encode()))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return
}
