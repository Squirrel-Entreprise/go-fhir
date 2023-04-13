package models_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
)

type Bundle struct {
	Client fhirInterface.IClient
	Entry  []struct {
		Resource struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"resource"`
	} `json:"entry"`
}

func (org *Bundle) ById(id string) fhirInterface.IParameters {
	fmt.Printf("\t\t--> ById()\n")
	return nil
}

func (org *Bundle) Where(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	return nil
}
