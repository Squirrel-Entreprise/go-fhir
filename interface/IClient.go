package fhirInterface

type IClient interface {
	GetRaw(uri string, p UrlParameters) ([]byte, error)
	Get(uri string, p UrlParameters, resType ResourceType) (IResourceResult, error)
	Search(resourceName ResourceType) IResource
}
