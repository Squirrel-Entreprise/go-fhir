package parameters_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/interface"
	"github.com/Squirrel-Entreprise/go-fhir/versions/r4"
)

type OrganizationParameters struct {
	Client     fhirInterface.IClient
	Uri        string
	Parameters fhirInterface.UrlParameters
}

func (org *OrganizationParameters) ReturnBundle() fhirInterface.IRequest {
	fmt.Println("\t\t\t--> ReturnBundle()")
	return &r4.Request{
		Client:       org.Client,
		Uri:          org.Uri,
		Parameters:   org.Parameters,
		TypeReturned: fhirInterface.BUNDLE,
	}
}

func (org *OrganizationParameters) Return() fhirInterface.IRequest {
	return nil
}

func (org *OrganizationParameters) ReturnRaw() fhirInterface.IRequest {
	fmt.Println("\t\t\t--> ReturnRaw()")
	return &r4.Request{
		Client:       org.Client,
		Uri:          org.Uri,
		Parameters:   org.Parameters,
		TypeReturned: fhirInterface.RAW,
	}
}

func (org *OrganizationParameters) And(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	fmt.Println("\t\t--> And()")
	org.Parameters = org.Parameters.Intersection(option)
	return org
}

func (org *OrganizationParameters) Or(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	fmt.Println("\t\t--> Or()")
	org.Parameters = org.Parameters.Union(option)
	return org
}
