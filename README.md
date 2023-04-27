# Golang FHIR Client Esante

[![Go Reference](https://pkg.go.dev/badge/github.com/Squirrel-Entreprise/go-fhir.svg)](https://pkg.go.dev/github.com/Squirrel-Entreprise/go-fhir)

## Introduction 🇫🇷
Bienvenue sur la librairie [Go FHIR](https://github.com/Squirrel-Entreprise/go-fhir), une bibliothèque open-source développée par [Squirrel](https://www.squirrel.fr) pour faciliter la manipulation des données de santé en utilisant le format [FHIR](https://www.hl7.org/fhir/) (Fast Healthcare Interoperability Resources). Cette librairie est conçue pour être performante et facile à utiliser, offrant une interface simple et intuitive pour interagir avec des serveurs FHIR.

## Introduction 🇬🇧
Welcome to the [Go FHIR](https://github.com/Squirrel-Entreprise/go-fhir) library, an open-source library developed by [Squirrel](https://www.squirrel.fr) to facilitate the manipulation of health data using the [FHIR](https://www.hl7.org/fhir/) (Fast Healthcare Interoperability Resources) format. This library is designed to be performant and easy to use, providing a simple and intuitive interface for interacting with FHIR servers.

## Sample

visible into `./cmd/usecases/physioReunionMayotte/main.go`

### Initialization

```go
apiKey := os.Getenv("ESANTE_API_KEY")
clientFhir := fhir.New("https://gateway.api.esante.gouv.fr/fhir", "ESANTE-API-KEY", apiKey, fhir.R4)
clientFhir.SetEntryLimit(500)
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

### Load the next page

```go
res = clientFhir.LoadPage().Next(res).Execute()
```

### Searching Organization by Id

```go
organizationRaw := clientFhir.
    Search(fhirInterface.ORGANIZATION).
    ById(e[0].GetOrganizationReference()).
    ReturnRaw().
    Execute()
```

## Credits

This package was inspired by the excellent HAPI FHIR Java library,
