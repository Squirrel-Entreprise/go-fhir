package models_r4

import (
	"fmt"
	"net/url"

	fhirInterface "github.com/Squirrel-Entreprise/go-fhir/interface"
	"github.com/Squirrel-Entreprise/go-fhir/versions/r4"
)

// This BundleResult is not complete and is only used for testing
// TODO: Make it complete, and make it for the other models
// TODO: Make a parser for the BundleResult (and the other models)
type BundleResult struct {
	Client fhirInterface.IClient
	Id     string `json:"id"`
	Link   []struct {
		Relation string `json:"relation"`
		Url      string `json:"url"`
	} `json:"link"`
	Entry []struct {
		Resource struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"resource"`
	} `json:"entry"`
	// TODO: Add more fields
}

func (b *BundleResult) GetId() string {
	return b.Id
}

func (b *BundleResult) GetNextLink() string {
	for _, link := range b.Link {
		if link.Relation == "next" {
			return link.Url
		}
	}
	return ""
}

func (b *BundleResult) MakeRequestNextPage() (fhirInterface.IRequest, error) {
	nextLink := b.GetNextLink()
	if nextLink == "" {
		return nil, fmt.Errorf("No next link found")
	}
	url, err := url.Parse(nextLink)
	if err != nil {
		return nil, err
	}
	urlParams := url.Query()
	return &r4.Request{
		Client: b.Client,
		Uri:    "/",
		Parameters: fhirInterface.UrlParameters{
			GetPages:   urlParams.Get("_getpages"),
			PageId:     urlParams.Get("_pageId"),
			BundleType: urlParams.Get("_bundletype"),
		},
		TypeReturned: fhirInterface.BUNDLE,
	}, nil
}

type Bundle struct {
	Client fhirInterface.IClient
}

func (org *Bundle) ById(id string) fhirInterface.IParameters {
	fmt.Printf("\t\t--> ById()\n")
	return nil
}

func (org *Bundle) Where(option fhirInterface.UrlParameters) fhirInterface.IParameters {
	return nil
}
