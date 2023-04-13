package fhirInterface

type IClient interface {
	LoadPage() struct {
		Next func(IResourceResult) IRequest
	}
	GetBaseUrl() string
	GetRaw(uri string, p UrlParameters) ([]byte, error)
	Get(uri string, p UrlParameters, resType ResourceType) (IResourceResult, error)
	Search(resourceName ResourceType) IResource
}
