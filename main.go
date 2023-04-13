package main

import (
	"fmt"
	"os"
	"time"

	fhir "github.com/Squirrel-Entreprise/go-fhir/cmd"
	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
	models_r4 "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/versions/r4/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("🤯 Error loading .env file")
	}
	apiKey := os.Getenv("ESANTE_API_KEY")
	timeStart := time.Now()
	fmt.Println("🏁 Starting test of go-fhir...")
	clientFhir := fhir.New("https://gateway.api.esante.gouv.fr/fhir", "ESANTE-API-KEY", apiKey, fhir.R4)

	// print the result
	/*organization, err := clientFhir.GetOrganizationByName("imagerie")
	if err != nil {
		fmt.Printf("🤯 Error: %v", err)
	} else {
		fmt.Println("🏤 Organisation (contenant 'imagerie') : ", organization)
	}*/

	// print the result
	/*var res fhirInterface.IResource = clientFhir.
		Search(fhirInterface.ORGANIZATION).
		Where(models_r4.
			Organization{}.
			Name.
			Contains().
			Value("imagerie")).
		Or(models_r4.
			Organization{}.
			Name.
			Contains().
			Value("centre")).
		ReturnBundle().Execute()
	fmt.Println("🏤 Organisation (contenant 'imagerie') : ", res)*/

	// print the result of sample 1
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

	// details about the first entry
	fmt.Printf("👨‍⚕️ PractitionerRole 1 details : \n")
	practitionerRoleRaw := clientFhir.
		Search(fhirInterface.PRACTITIONER_ROLE).
		ById(bundleRes.(*models_r4.Bundle).Entry[0].Resource.Id).
		ReturnRaw().
		Execute()

	// print the raw result in a string
	fmt.Println(string(practitionerRoleRaw.([]byte)))

	// get the practitioner with the Id 003-357936
	fmt.Println("👨‍⚕️ Practitioner with Id = 003-357936 : ")
	practitionerRaw := clientFhir.
		Search(fhirInterface.PRACTITIONER).
		ById("003-357936").
		ReturnRaw().
		Execute()

	// print the raw result in a string
	fmt.Println(string(practitionerRaw.([]byte)))

	// get the organization with the Id 001-01-702556
	fmt.Println("🏤 Organization with Id = 001-01-702556 : ")
	organizationRaw := clientFhir.
		Search(fhirInterface.ORGANIZATION).
		ById("001-01-702556").
		ReturnRaw().
		Execute()

	// print the raw result in a string
	fmt.Println(string(organizationRaw.([]byte)))

	timeEnd := time.Now()
	fmt.Printf("🏁 Finished test of go-fhir in %v seconds ! 🎉\n", timeEnd.Sub(timeStart).Seconds())
}
