package fhirInterface

type IOrganization interface {
	GetName() (string, error)
	//Where(i interface{})
}
