package parameters_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/interface"
	"github.com/Squirrel-Entreprise/go-fhir/versions/r4"
)

type PractitionerRoleParameters struct {
	Client     fhirInterface.IClient
	Uri        string
	Parameters fhirInterface.UrlParameters
}

func (pr *PractitionerRoleParameters) ReturnBundle() fhirInterface.IRequest {
	fmt.Println("\t\t\t--> ReturnBundle()")
	return &r4.Request{
		Client:       pr.Client,
		Uri:          pr.Uri,
		Parameters:   pr.Parameters,
		TypeReturned: fhirInterface.BUNDLE,
	}
	//return nil
}

func (pr *PractitionerRoleParameters) Return() fhirInterface.IRequest {
	fmt.Println("\t\t\t--> Return()")
	return nil
}

func (pr *PractitionerRoleParameters) ReturnRaw() fhirInterface.IRequest {
	fmt.Println("\t\t\t--> ReturnRaw()")
	return &r4.Request{
		Client:       pr.Client,
		Uri:          pr.Uri,
		Parameters:   pr.Parameters,
		TypeReturned: fhirInterface.RAW,
	}
}

func (pr *PractitionerRoleParameters) And(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	fmt.Println("\t\t--> And()")
	pr.Parameters = pr.Parameters.Intersection(option)
	return pr
}

func (pr *PractitionerRoleParameters) Or(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	fmt.Println("\t\t--> Or()")
	pr.Parameters = pr.Parameters.Union(option)
	return pr
}
