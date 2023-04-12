package fhirInterface

type UrlParameters struct {
	Name string
}

type FhirValue struct {
	Value string
}

func (f FhirValue) Matches() struct {
	Value func(v string) UrlParameters
} {
	return struct {
		Value func(v string) UrlParameters
	}{
		Value: func(v string) UrlParameters {
			return UrlParameters{
				Name: v,
			}
		},
	}
}

type IClient interface {
	//GetOrganizationByName(name string) (IOrganization, error)
	Get(uri string, p UrlParameters, resType ResourceType) (IResource, error)
	Search(resourceName ResourceType) IResource
}

// IResources
// what is the resource we're looking for
