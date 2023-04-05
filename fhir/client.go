package fhir

// HTTP client for making requests to FHIR servers.

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// HTTPClient
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	ClientHttp HTTPClient
)

// Initializations

func init() {
	ClientHttp = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        0,
			MaxIdleConnsPerHost: 10,
		},
	}
}

// Client is a FHIR client.
type Fhir struct {
	BaseURL  string
	ApiKey   string
	ApiValue string
	//VersionFhir string
}

// NewClient creates a new FHIR client.
func New(baseURL, apiKey, apiValue string) *Fhir {
	return &Fhir{
		BaseURL:  baseURL,
		ApiKey:   apiKey,
		ApiValue: apiValue,
		//Version: version,
	}
}

func (f *Fhir) call(method, path string, body interface{}) (*http.Response, error) {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, f.BaseURL+path, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(f.ApiKey, f.ApiValue)

	return ClientHttp.Do(req)
}

func (f *Fhir) Get(uri string) (*http.Response, error) {
	values := url.Values{}
	path := (&url.URL{
		Path:     uri,
		RawQuery: values.Encode(),
	}).String()
	return f.call("GET", path, nil)
}
