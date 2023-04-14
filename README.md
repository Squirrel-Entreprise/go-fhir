# Golang FHIR Client Esante

[![Go Reference](https://pkg.go.dev/badge/github.com/Squirrel-Entreprise/go-fhir.svg)](https://pkg.go.dev/github.com/Squirrel-Entreprise/go-fhir)

## Sample
visible into `./cmd/sample/main.go`

### Initialization
```go
apiKey := os.Getenv("ESANTE_API_KEY")
clientFhir := fhir.New("https://gateway.api.esante.gouv.fr/fhir", "ESANTE-API-KEY", apiKey, fhir.R4)
```

### Searching PractitionerRole by Role and Active Status
Please note that we receive a prototype of the BundleResult struct, which is not yet complete, after executing the request.
```go
bundleRes := clientFhir.
    Search(fhirInterface.PRACTITIONER_ROLE).
    Where(models_r4.PractitionerRole{}.
        Role.
        Contains().
        Value("70")).
    And(models_r4.PractitionerRole{}.
        Active.
        IsActive()).
    ReturnBundle().Execute()
```

### Searching PractitionerRole by ID
```go
practitionerRoleRaw := clientFhir.
    Search(fhirInterface.PRACTITIONER_ROLE).
    ById(bundleRes.(*models_r4.BundleResult).Entry[0].Resource.Id).
    ReturnRaw().
    Execute()
```

## Credits
This package was inspired by the excellent HAPI FHIR Java library,
