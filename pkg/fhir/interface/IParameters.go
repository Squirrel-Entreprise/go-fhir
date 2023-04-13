package fhirInterface

type IParameters interface {
	And(up UrlParameters) IParameters
	Or(up UrlParameters) IParameters
	//ReturnBundle()
	ReturnBundle() IRequest
	Return() IRequest
	ReturnRaw() IRequest
}

// IRequest
// Is the request that we have to execute
