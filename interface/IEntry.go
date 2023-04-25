package fhirInterface

type IEntry interface {
	GetId() string
	GetName() string
	GetPractitionerReference() string
	GetOrganizationReference() string
}
