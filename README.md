# Golang FHIR Client

[![Go Reference](https://pkg.go.dev/badge/github.com/Squirrel-Entreprise/go-fhir.svg)](https://pkg.go.dev/github.com/Squirrel-Entreprise/go-fhir)

## Sample
visible into `./cmd/sample/main.go`

```go
apiKey := os.Getenv("ESANTE_API_KEY")
clientFhir := fhir.New("https://gateway.api.esante.gouv.fr/fhir", "ESANTE-API-KEY", apiKey, fhir.R4)

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


fmt.Printf("👨‍⚕️ PractitionerRole 0, details : \n")

practitionerRoleRaw := clientFhir.
    Search(fhirInterface.PRACTITIONER_ROLE).
    ById(bundleRes.(*models_r4.BundleResult).Entry[0].Resource.Id).
    ReturnRaw().
    Execute()

fmt.Println(string(practitionerRoleRaw.([]byte)))


fmt.Println("👨‍⚕️ Practitioner with Id = 003-357936 : ")

practitionerRaw := clientFhir.
    Search(fhirInterface.PRACTITIONER).
    ById("003-357936").
    ReturnRaw().
    Execute()

fmt.Println(string(practitionerRaw.([]byte)))
```