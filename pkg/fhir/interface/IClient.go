package fhirInterface

type UrlParameters struct {
	Name string
}

type IClient interface {
	//GetOrganizationByName(name string) (IOrganization, error)
	Get(uri string, p UrlParameters, res interface{}) error
	Search(resourceName Resource) IResource
}

// IResources
// what is the resource we're looking for
