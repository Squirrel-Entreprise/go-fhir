package models_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
)

type BundleResult struct {
	Id    string `json:"id"`
	Entry []struct {
		Resource struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"resource"`
	} `json:"entry"`
}

func (b *BundleResult) GetId() string {
	return b.Id
}

type Bundle struct {
	Client fhirInterface.IClient
}

func (org *Bundle) ById(id string) fhirInterface.IParameters {
	fmt.Printf("\t\t--> ById()\n")
	return nil
}

func (org *Bundle) Where(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	return nil
}
