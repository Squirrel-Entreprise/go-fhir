package clients_r4

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
	models_r4 "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/versions/r4/models"
)

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
		err = json.NewDecoder(response.Body).Decode(res)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *fhir) GetRaw(uri string, p fhirInterface.UrlParameters) ([]byte, error) {
	values := p.BuildUrlValues()
	path := &url.URL{
		Path:     uri,
		RawQuery: values.Encode(),
	}

	fmt.Println("\t\t\t\t\t", "--> GetRAW:", f.BaseURL+path.String())
	req, err := http.NewRequest("GET", f.BaseURL+path.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(f.ApiKey, f.ApiValue)

	res, err := f.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if res != nil {
		defer res.Body.Close()
	}

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, nil
}

func (f *fhir) Get(uri string, p fhirInterface.UrlParameters, resType fhirInterface.ResourceType) (fhirInterface.IResourceResult, error) {
	values := p.BuildUrlValues()
	path := &url.URL{
		Path:     uri,
		RawQuery: values.Encode(),
	}

	switch resType {
	case fhirInterface.BUNDLE:
		res := &models_r4.BundleResult{}
		err := f.call("GET", path, nil, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	return nil, nil
}

func (f *fhir) Search(r fhirInterface.ResourceType) fhirInterface.IResource {
	switch r {
	case fhirInterface.ORGANIZATION:
		fmt.Println("\t--> Search(Organization)")
		return &models_r4.Organization{
			Client: f,
		}
	case fhirInterface.PRACTITIONER_ROLE:
		fmt.Println("\t--> Search(PractitionerRole)")
		return &models_r4.PractitionerRole{
			Client: f,
		}

	case fhirInterface.PRACTITIONER:
		fmt.Println("\t--> Search(Practitioner)")
		return &models_r4.Practitioner{
			Client: f,
		}
	}
	return nil
}
