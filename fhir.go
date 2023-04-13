package fhir

import (
	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/interface"
	clients_r4 "github.com/Squirrel-Entreprise/go-fhir/versions/r4/clients"
)

type FhirVersion string

const (
	R4 FhirVersion = "r4"
)

func New(baseUrl string, apiKey string, apiValue string, version FhirVersion) fhirInterface.IClient {
	switch version {
	case R4:
		return clients_r4.NewFhirClient(baseUrl, apiKey, apiValue)
	default:
		return nil
	}
}

// New creates a new FHIR client for the given version.
