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
	var res fhirInterface.IResource = clientFhir.
		Search(fhirInterface.PRACTITIONER_ROLE).
		Where(models_r4.PractitionerRole{}.
			Role.
			Contains().
			Value("70")).
		And(models_r4.PractitionerRole{}.
			Active.
			IsActive()).
		ReturnBundle().Execute()

	fmt.Println(" 🏥 PractitionerRole (role '70' and active = true) : ", res)

	timeEnd := time.Now()
	fmt.Printf("🏁 Finished test of go-fhir in %v seconds ! 🎉\n", timeEnd.Sub(timeStart).Seconds())
}
