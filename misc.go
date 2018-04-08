package somtoday

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetOrganisations returns a list of organisations
func GetOrganisations() ([]Organisation, error) {
	resp, err := http.Get("https://servers.somtoday.nl/organisaties.json")
	if err != nil {
		return []Organisation{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Organisation{}, err
	}

	var response []OrganisationResponse // The serialization is retarded
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []Organisation{}, err
	}

	return response[0].Instellingen, err
}
