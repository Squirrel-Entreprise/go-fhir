package models_r4

type Entry struct {
	Resource struct {
		ResourceType string `json:"resourceType"`
		Id           string `json:"id"`
		Practitioner struct {
			Reference string `json:"reference"`
		} `json:"practitioner"`
		Organization struct {
			Reference string `json:"reference"`
		} `json:"organization"`
	} `json:"resource"`
}

func (e *Entry) GetId() string {
	return e.Resource.Id
}
func (e *Entry) GetPractitionerReference() string {
	if e.Resource.Practitioner.Reference == "" {
		return ""
	}
	return e.Resource.Practitioner.Reference[13:]
}

func (e *Entry) GetOrganizationReference() string {
	if e.Resource.Organization.Reference == "" {
		return ""
	}
	return e.Resource.Organization.Reference[13:]
}
