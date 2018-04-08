package somtoday

import (
	"testing"
)

func TestGetOrganisations(t *testing.T) {
	organisations, err := GetOrganisations()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	// Check if my school is in there (in a clever hacky way), if it isn't, it probably didn't fetch it correctly
	inList := func() bool { // I know I only use inList once, but otherwise I'd have to place the function below in an if statement condition and that's worse IMO
		for _, organisation := range organisations {
			if organisation.Naam == "Bonhoeffer College" {
				return true
			}
		}
		return false
	}()
	if !inList {
		t.Log("Did not fetch correctly")
		t.Log(organisations)
		t.Fail()
	}
}
