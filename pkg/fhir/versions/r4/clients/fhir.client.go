package clients_r4

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/pkg/fhir/interface"
)

// HTTP client for making requests to FHIR servers.
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
type UrlParameters struct {
	Name string
}

var (
	ClientHttp httpClient
)

type fhir struct {
	BaseURL  string
	ApiKey   string
	ApiValue string
}

func New(baseURL, apiKey, apiValue string) fhirInterface.IFhirClient {
	return &fhir{
		BaseURL:  baseURL,
		ApiKey:   apiKey,
		ApiValue: apiValue,
	}
}
func init() {
	ClientHttp = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxIdleConnsPerHost: 10,
		},
	}
}
func (f *fhir) call(method string, path *url.URL, payload []byte, res interface{}) error {

	fmt.Println(method, "->", f.BaseURL+path.String())

	req, err := http.NewRequest(method, f.BaseURL+path.String(), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(f.ApiKey, f.ApiValue)

	response, err := ClientHttp.Do(req)
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

func (f *fhir) get(uri string, p UrlParameters, res interface{}) error {
	values := url.Values{}

	if p.Name != "" {
		values.Add("name:contains", p.Name)
	}

	path := &url.URL{
		Path:     uri,
		RawQuery: values.Encode(),
	}
	return f.call("GET", path, nil, res)
}
