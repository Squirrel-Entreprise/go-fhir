package clients_r4

import (
	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
	models_r4 "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/versions/r4/models"
)

func (f *fhir) GetOrganizationByName(name string) (fhirInterface.IOrganization, error) {
	uri := "/Organization"
	urlParams := UrlParameters{
		Name: name,
	}
	var org models_r4.Organization
	err := f.get(uri, urlParams, &org)
	if err != nil {
		return nil, err
	}
	return &org, nil
}
