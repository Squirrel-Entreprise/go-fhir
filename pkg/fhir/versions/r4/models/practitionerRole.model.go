package models_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
	parameters_r4 "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/versions/r4/parameters"
)

type PractitionerRole struct {
	Client  fhirInterface.IClient
	Address fhirInterface.FhirAddress
	Name    fhirInterface.FhirName
	Role    fhirInterface.FhirRole
	Active  fhirInterface.FhirActive
	/*Entry  []struct {
		Resource struct {
			Name string `json:"name"`
		} `json:"resource"`
	} `json:"entry"`*/
}

// The where funciton is here to add parameters to the request
func (pr *PractitionerRole) Where(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	fmt.Printf("\t\t--> Where()\n")

	return &parameters_r4.PractitionerRoleParameters{
		Client:     pr.Client,
		Uri:        "/PractitionerRole",
		Parameters: option,
	}
}
