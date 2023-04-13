package models_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
	parameters_r4 "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/versions/r4/parameters"
)

type Practitioner struct {
	Client  fhirInterface.IClient
	Address fhirInterface.FhirAddress
	Name    fhirInterface.FhirName
}

func (p *Practitioner) ById(id string) fhirInterface.IParameters {
	fmt.Printf("\t\t--> ById()\n")

	return &parameters_r4.PractitionerParameters{
		Client: p.Client,
		Uri:    "/Practitioner/" + id,
	}
}

// The where funciton is here to add parameters to the request
func (p *Practitioner) Where(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	fmt.Printf("\t\t--> Where()\n")

	return &parameters_r4.OrganizationParameters{
		Client:     p.Client,
		Uri:        "/Practitioner",
		Parameters: option,
	}
}
