package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	fhir "github.com/Squirrel-Entreprise/go-fhir"
	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/interface"
	models_r4 "github.com/Squirrel-Entreprise/go-fhir/versions/r4/models"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("FetchAllPhysiotherapists")

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("🤯 Error loading .env file")
	}
	apiKey := os.Getenv("ESANTE_API_KEY")

	clientFhir := fhir.New("https://gateway.api.esante.gouv.fr/fhir", "ESANTE-API-KEY", apiKey, fhir.R4)

	// LIMIT 500
	clientFhir.SetEntryLimit(500)
	clientFhir.SetTimeout(30)

	// Look for all physiotherapists
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

	res, ok := bundleRes.(*models_r4.BundleResult)
	if !ok {
		log.Println("error bundleRes type assertion")
		return
	}

	for {
		for _, e := range res.Entry {

			// Look for details of the organization
			if (e.GetOrganizationReference() == "") || (e.GetPractitionerReference() == "") {
				continue
			}
			organizationRaw := clientFhir.
				Search(fhirInterface.ORGANIZATION).
				ById(e.GetOrganizationReference()).
				ReturnRaw().
				Execute()

			// If the organization is in Mayotte or Reunion,
			// we look for the practitioner's details
			postalCodes := extractPostalCodesFromJson(organizationRaw.([]byte))
			if len(postalCodes) == 0 {
				continue
			}
			isMayotteOrReunion := false
			for _, postalCode := range postalCodes {
				if postalCode == "" {
					continue
				}
				if postalCode[:3] == "974" || postalCode[:3] == "976" {
					isMayotteOrReunion = true
					break
				}
			}
			if !isMayotteOrReunion {
				continue
			}

			log.Println("✅ Found a physiotherapist in Mayotte or Reunion")

			practitionerRaw := clientFhir.
				Search(fhirInterface.PRACTITIONER).
				ById(e.GetPractitionerReference()).
				ReturnRaw().
				Execute()

			practitionerRoleRaw := clientFhir.
				Search(fhirInterface.PRACTITIONER_ROLE).
				ById(e.GetId()).
				ReturnRaw().
				Execute()

			log.Print("Practitioner : \n\n", string(practitionerRaw.([]byte)), "\n\n")
			log.Print("PractitionerRoleRaw : \n\n", string(practitionerRoleRaw.([]byte)), "\n\n")
		}
		if res.GetNextLink() == "" {
			break
		}
		log.Println("next link", res.GetNextLink())
		res = clientFhir.LoadPage().Next(res).Execute().(*models_r4.BundleResult)
	}
}

func extractPostalCodesFromJson(jsonData []byte) []string {
	dec := json.NewDecoder(strings.NewReader(string(jsonData)))
	// max 2 postal codes
	postalCodes := make([]string, 2)
	for {
		t, err := dec.Token()
		if err != nil {
			break
		}
		if s, ok := t.(string); ok {
			if s == "postalCode" {
				t, err := dec.Token()
				if err != nil {
					fmt.Println("err", err)
					break
				}
				postalCodes = append(postalCodes, t.(string))
			}
		}
	}
	return postalCodes
}
