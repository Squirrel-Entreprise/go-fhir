package main

import (
	"fmt"
	"os"
	"time"

	fhir "github.com/Squirrel-Entreprise/go-fhir"
	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/interface"
	models_r4 "github.com/Squirrel-Entreprise/go-fhir/versions/r4/models"
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

	// The BundleResult returned by ReturnBundle() is a not complete prototype
	fmt.Printf("👨‍⚕️ PractitionerRole 0, details : \n")
	practitionerRoleRaw := clientFhir.
		Search(fhirInterface.PRACTITIONER_ROLE).
		ById(bundleRes.(*models_r4.BundleResult).Entry[0].Resource.Id).
		ReturnRaw().
		Execute()
	fmt.Println(string(practitionerRoleRaw.([]byte)))

	// NextPage of results
	bundleRes = clientFhir.LoadPage().Next(bundleRes.(*models_r4.BundleResult)).Execute()
	fmt.Println("👨‍⚕️ Next page of results : ", bundleRes.(*models_r4.BundleResult))

	fmt.Printf("👨‍⚕️ PractitionerRole 0 on next page, details : \n")
	practitionerRoleRaw = clientFhir.
		Search(fhirInterface.PRACTITIONER_ROLE).
		ById(bundleRes.(*models_r4.BundleResult).Entry[0].Resource.Id).
		ReturnRaw().
		Execute()
	fmt.Println(string(practitionerRoleRaw.([]byte)))

	// Get the practitioner with the Id 003-357936
	fmt.Println("👨‍⚕️ Practitioner with Id = 003-357936 : ")
	practitionerRaw := clientFhir.
		Search(fhirInterface.PRACTITIONER).
		ById("003-357936").
		ReturnRaw().
		Execute()
	fmt.Println(string(practitionerRaw.([]byte)))

	// Get the organization with the Id 001-01-702556
	fmt.Println("🏤 Organization with Id = 001-01-702556 : ")
	organizationRaw := clientFhir.
		Search(fhirInterface.ORGANIZATION).
		ById("001-01-702556").
		ReturnRaw().
		Execute()
	fmt.Println(string(organizationRaw.([]byte)))

	timeEnd := time.Now()
	fmt.Printf("🏁 Finished test of go-fhir in %v seconds ! 🎉\n", timeEnd.Sub(timeStart).Seconds())
}
