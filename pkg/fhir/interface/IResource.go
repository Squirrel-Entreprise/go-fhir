package fhirInterface

type Resource string

const (
	BUNDLE       Resource = "Bundle"
	ORGANIZATION Resource = "Organization"
	PRACTITIONER Resource = "Practitioner"
	PATIENT      Resource = "Patient"
)

type IResource interface {
	//New(client IFhirClient) IResource
	Where(option string) IParameters
}

// IParameter
// what are the parameters we're looking for
