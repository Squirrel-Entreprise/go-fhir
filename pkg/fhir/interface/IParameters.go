package fhirInterface

type IParameters interface {
	And(up UrlParameters) IParameters
	Or(up UrlParameters) IParameters
	ReturnBundle() IRequest
	Return() IRequest
	ReturnRaw() IRequest
}
