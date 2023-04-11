package fhirInterface

type IParameters interface {
	//And() IRequest
	//Or() IRequest
	//ReturnBundle()
	ReturnBundle() IRequest
}

// IRequest
// Is the request that we have to execute
