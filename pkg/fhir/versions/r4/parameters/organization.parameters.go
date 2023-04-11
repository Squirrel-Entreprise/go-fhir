package parameters_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
	"github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/versions/r4"
)

type OrganizationParameters struct {
	Client fhirInterface.IClient
	Query  string
}

func (org *OrganizationParameters) ReturnBundle() fhirInterface.IRequest {
	fmt.Println("\t\t\t--> ReturnBundle()")
	return &r4.Request{
		Client:       org.Client,
		Uri:          org.Query,
		TypeReturned: fhirInterface.BUNDLE,
	}
	//return nil
}
