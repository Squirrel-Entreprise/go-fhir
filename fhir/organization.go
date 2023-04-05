package fhir

import (
	"encoding/json"
)

type Organization struct {
	//ResourceType string `json:"resourceType"`
	Entry []struct {
		Resource struct {
			Name string `json:"name"`
		} `json:"resource"`
	} `json:"entry"`
}

func (f *Fhir) GetOrganizationByName(name string) (*Organization, error) {
	uri := "/Organization?name%3Acontains=" + name
	res, err := f.call("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var org Organization
	json.NewDecoder(res.Body).Decode(&org)
	return &org, nil
}
