package fhirInterface

type IRequest interface {
	Execute() interface{}
}
