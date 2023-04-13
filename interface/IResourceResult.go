package fhirInterface

type IResourceResult interface {
	GetId() string
	GetNextLink() string
	MakeRequestNextPage() (IRequest, error)
}
