package models_r4

type Organization struct {
	Entry []struct {
		Resource struct {
			Name string `json:"name"`
		} `json:"resource"`
	} `json:"entry"`
}

func (org *Organization) GetName() (string, error) {
	return org.Entry[0].Resource.Name, nil
}
