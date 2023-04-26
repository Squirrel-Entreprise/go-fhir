package fhirInterface

type IEntry interface {
	GetId() string
	GetPractitionerReference() string
	GetOrganizationReference() string
}
