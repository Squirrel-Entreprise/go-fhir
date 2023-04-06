package fhirInterface

type IFhirClient interface {
	GetOrganizationByName(name string) (IOrganization, error)
}

// interface organization

// interface patient

// ...
