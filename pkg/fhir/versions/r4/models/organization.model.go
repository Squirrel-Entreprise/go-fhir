package models_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
	parameters_r4 "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/versions/r4/parameters"
)

type Organization struct {
	Client fhirInterface.IClient
	Entry  []struct {
		Resource struct {
			Name string `json:"name"`
		} `json:"resource"`
	} `json:"entry"`
}

func (org *Organization) Where(option string) fhirInterface.IParameters {
	fmt.Printf("\t\t--> Where(%s)\n", option)
	return &parameters_r4.OrganizationParameters{
		Client: org.Client,
		Query:  fmt.Sprintf("/Organization?%s", option),
	}
}
