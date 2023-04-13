package r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
)

type Request struct {
	Client       fhirInterface.IClient
	Uri          string
	Parameters   fhirInterface.UrlParameters
	TypeReturned fhirInterface.ResourceType
}

func (req *Request) Execute() interface{} {
	var err error
	if req.TypeReturned == fhirInterface.RAW {
		fmt.Println("\t\t\t\t--> ExecuteRaw()")
		var resRaw []byte
		resRaw, err = req.Client.GetRaw(req.Uri, req.Parameters)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return resRaw
	}
	fmt.Println("\t\t\t\t--> Execute()")
	var res fhirInterface.IResourceResult
	res, err = req.Client.Get(req.Uri, req.Parameters, req.TypeReturned)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res
}
