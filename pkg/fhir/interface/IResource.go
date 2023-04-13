package fhirInterface

type ResourceType string

const (
	BUNDLE            ResourceType = "Bundle"
	ORGANIZATION      ResourceType = "Organization"
	PRACTITIONER      ResourceType = "Practitioner"
	PRACTITIONER_ROLE ResourceType = "PractitionerRole"
	RAW               ResourceType = "Raw"
)

type IResource interface {
	ById(id string) IParameters
	Where(option UrlParameters) IParameters
}
