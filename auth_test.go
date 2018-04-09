package somtoday

import (
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	organisation, err := GetOrganisation("Bonhoeffer College")
	t.Log(organisation)
	t.Log(err)
	if err != nil {
		t.Fail()
	}
	auth, err := Login(organisation, os.Getenv("SOM_USERNAME"), os.Getenv("SOM_PASSWORD")) // safe
	t.Log(auth)
	t.Log(err)
	if err != nil {
		t.Fail()
	}

	// Check if it fetched correctly
	if auth.SomtodayTenant != "bonhoeffer" {
		t.Log("Did not fetch correctly")
		t.Fail()
	}
}

func TestRefreshTokens(t *testing.T) {
	// Get auth tokens
	organisation, err := GetOrganisation("Bonhoeffer College")
	t.Log(organisation)
	t.Log(err)
	if err != nil {
		t.Fail()
	}
	auth, err := Login(organisation, os.Getenv("SOM_USERNAME"), os.Getenv("SOM_PASSWORD")) // safe
	t.Log(auth)
	t.Log(err)
	if err != nil {
		t.Fail()
	}

	// Check if it fetched correctly
	if auth.SomtodayTenant != "bonhoeffer" {
		t.Log("Did not fetch correctly")
		t.Fail()
	}

	// Refresh them
	newAuth, err := RefreshTokens(auth.RefreshToken)
	t.Log(newAuth)
	t.Log(err)
	if err != nil {
		t.Fail()
	}

	// Check if the tokens from old and new are the same
	if auth.AccessToken == newAuth.AccessToken {
		t.Log("Tokens are the same, it did not refresh")
		t.Fail()
	}
}
