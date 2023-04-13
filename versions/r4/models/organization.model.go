package models_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/interface"
	parameters_r4 "github.com/Squirrel-Entreprise/go-fhir/versions/r4/parameters"
)

type Organization struct {
	Client  fhirInterface.IClient
	Address fhirInterface.FhirAddress
	Name    fhirInterface.FhirName
}

func (org *Organization) ById(id string) fhirInterface.IParameters {
	fmt.Printf("\t\t--> ById()\n")

	return &parameters_r4.OrganizationParameters{
		Client: org.Client,
		Uri:    "/Organization/" + id,
	}
}

func (org *Organization) Where(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	fmt.Printf("\t\t--> Where()\n")

	return &parameters_r4.OrganizationParameters{
		Client:     org.Client,
		Uri:        "/Organization",
		Parameters: option,
	}
}
