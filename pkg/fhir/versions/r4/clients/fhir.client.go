package clients_r4

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
	models_r4 "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/versions/r4/models"
)

// HTTP client for making requests to FHIR servers.
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type fhir struct {
	Client   httpClient
	BaseURL  string
	ApiKey   string
	ApiValue string
}

func NewFhirClient(baseURL, apiKey, apiValue string) fhirInterface.IClient {
	clientHttp := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxIdleConnsPerHost: 10,
		},
	}
	return &fhir{
		Client:   clientHttp,
		BaseURL:  baseURL,
		ApiKey:   apiKey,
		ApiValue: apiValue,
	}
}

type UrlParameters struct {
	Name string
}

/*
func init() {
	ClientHttp = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxIdleConnsPerHost: 10,
		},
	}
}*/
func (f *fhir) call(method string, path *url.URL, payload []byte, res interface{}) error {

	fmt.Println("\t\t\t\t\t", "-->", method, ":", f.BaseURL+path.String())

	req, err := http.NewRequest(method, f.BaseURL+path.String(), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(f.ApiKey, f.ApiValue)

	response, err := f.Client.Do(req)
	if err != nil {
		return err
	}
	if response != nil {
		defer response.Body.Close()
	}

	if response.StatusCode == 200 {
		return json.NewDecoder(response.Body).Decode(&res)
	}
	return nil

}
func (f *fhir) Get(uri string, p fhirInterface.UrlParameters, resType fhirInterface.ResourceType) (fhirInterface.IResource, error) {
	values := url.Values{}

	if p.Name != "" {
		values.Add("name:contains", p.Name)
	}

	path := &url.URL{
		Path:     uri,
		RawQuery: values.Encode(),
	}

	switch resType {
	case fhirInterface.BUNDLE:
		res := &models_r4.Bundle{}
		err := f.call("GET", path, nil, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, nil
	//return f.call("GET", path, nil, res)
}

func (f *fhir) Search(r fhirInterface.ResourceType) fhirInterface.IResource {
	switch r {
	case fhirInterface.ORGANIZATION:
		fmt.Println("\t--> Search(Organization)")
		return &models_r4.Organization{
			Client: f,
		}
	}
	return nil
}
