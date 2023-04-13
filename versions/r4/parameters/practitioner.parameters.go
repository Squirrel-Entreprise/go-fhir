package parameters_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/interface"
	"github.com/Squirrel-Entreprise/go-fhir/versions/r4"
)

type PractitionerParameters struct {
	Client     fhirInterface.IClient
	Uri        string
	Parameters fhirInterface.UrlParameters
}

func (prac *PractitionerParameters) ReturnBundle() fhirInterface.IRequest {
	fmt.Println("\t\t\t--> ReturnBundle()")
	return &r4.Request{
		Client:       prac.Client,
		Uri:          prac.Uri,
		Parameters:   prac.Parameters,
		TypeReturned: fhirInterface.BUNDLE,
	}
}

func (prac *PractitionerParameters) Return() fhirInterface.IRequest {
	return nil
}

func (p *PractitionerParameters) ReturnRaw() fhirInterface.IRequest {
	fmt.Println("\t\t\t--> ReturnRaw()")
	return &r4.Request{
		Client:       p.Client,
		Uri:          p.Uri,
		Parameters:   p.Parameters,
		TypeReturned: fhirInterface.RAW,
	}
}

func (prac *PractitionerParameters) And(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	fmt.Println("\t\t\t--> And()")
	prac.Parameters = prac.Parameters.Intersection(option)
	return prac
}

func (prac *PractitionerParameters) Or(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	fmt.Println("\t\t\t--> Or()")
	prac.Parameters = prac.Parameters.Union(option)
	return prac
}
