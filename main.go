package main

import (
	"fmt"
	"os"
	"time"

	fhir "github.com/Squirrel-Entreprise/go-fhir/cmd"
	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
)

func main() {
	timeStart := time.Now()
	fmt.Println("🏁 Starting test of go-fhir...")
	apiKey := os.Getenv("ESANTE_API_KEY")
	clientFhir := fhir.New("https://gateway.api.esante.gouv.fr/fhir", "ESANTE-API-KEY", apiKey, fhir.R4)

	// print the result
	/*organization, err := clientFhir.GetOrganizationByName("imagerie")
	if err != nil {
		fmt.Printf("🤯 Error: %v", err)
	} else {
		fmt.Println("🏤 Organisation (contenant 'imagerie') : ", organization)
	}*/

	// print the result
	var res fhirInterface.IResource = clientFhir.Search(fhirInterface.ORGANIZATION).Where("name:contains=imagerie").ReturnBundle().Execute()
	fmt.Println("🏤 Organisation (contenant 'imagerie') : ", res)

	timeEnd := time.Now()
	fmt.Printf("🏁 Finished test of go-fhir in %v seconds ! 🎉\n", timeEnd.Sub(timeStart).Seconds())
}
