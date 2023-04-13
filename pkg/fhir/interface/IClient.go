package fhirInterface

type IClient interface {
	//GetOrganizationByName(name string) (IOrganization, error)
	GetRaw(uri string, p UrlParameters) ([]byte, error)
	Get(uri string, p UrlParameters, resType ResourceType) (IResource, error)
	Search(resourceName ResourceType) IResource
}

// IResources
// what is the resource we're looking for
