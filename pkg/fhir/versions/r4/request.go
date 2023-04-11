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

func (req *Request) Execute() fhirInterface.IResource {
	fmt.Println("\t\t\t\t--> Execute()")
	var err error
	var res fhirInterface.IResource
	res, err = req.Client.Get(req.Uri, req.Parameters, req.TypeReturned)
	if err != nil {
		fmt.Println(err)
	}
	return res
}
