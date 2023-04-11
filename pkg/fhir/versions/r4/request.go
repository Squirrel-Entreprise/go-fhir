package r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
)

type Request struct {
	Client       fhirInterface.IClient
	Uri          string
	Parameters   fhirInterface.UrlParameters
	TypeReturned fhirInterface.Resource
}

func (req *Request) Execute() {
	req.Client.Get(req.Uri, req.Parameters, req.TypeReturned)
	fmt.Println("\t\t\t\t--> Execute()")
}
