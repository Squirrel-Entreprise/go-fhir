package models_r4

import fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"

type Bundle struct {
	Client fhirInterface.IClient
	Entry  []struct {
		Resource struct {
			Name string `json:"name"`
		} `json:"resource"`
	} `json:"entry"`
}

func (org *Bundle) Where(option string) fhirInterface.IParameters {
	return nil
}
