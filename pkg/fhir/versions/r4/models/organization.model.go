package models_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
	parameters_r4 "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/versions/r4/parameters"
)

type Organization struct {
	Client  fhirInterface.IClient
	Address fhirInterface.FhirValue
	Name    fhirInterface.FhirValue
	/*Entry  []struct {
		Resource struct {
			Name string `json:"name"`
		} `json:"resource"`
	} `json:"entry"`*/
}

// The where funciton is here to add parameters to the request
func (org *Organization) Where(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	fmt.Printf("\t\t--> Where(%s)\n", option)

	return &parameters_r4.OrganizationParameters{
		Client:     org.Client,
		Uri:        "/Organization",
		Parameters: option,
	}
}
