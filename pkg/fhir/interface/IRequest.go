package fhirInterface

type IRequest interface {
	Execute() IResource // Maybe interface{} after ?
}
