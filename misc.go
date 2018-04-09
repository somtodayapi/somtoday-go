package somtoday

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// AppID is the ID somtoday uses to verify it's the app making reqs
const AppID = "D50E0C06-32D1-4B41-A137-A9A850C892C2"

// AppSecret is kinda the same as AppID
const AppSecret = "vDdWdKwPNaPCyhCDhaCnNeydyLxSGNJX"

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

// GetOrganisation Get an organisation from SOMtoday by name
func GetOrganisation(name string) (org Organisation, err error) {
	organisations, err := GetOrganisations()

	for _, organisation := range organisations {
		if organisation.Naam == name {
			org = organisation
			break
		}
	}
	return
}
