package models_r4

import (
	"fmt"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/interface"
)

// This BundleResult is not complete and is only used for testing
// TODO: Make it complete, and make it for the other models
// TODO: Make a parser for the BundleResult (and the other models)
type BundleResult struct {
	Id    string `json:"id"`
	Entry []struct {
		Resource struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"resource"`
	} `json:"entry"`
	// TODO: Add more fields
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
