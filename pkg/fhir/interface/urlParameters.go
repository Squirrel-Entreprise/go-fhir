package fhirInterface

type UrlParameters struct {
	Name    string
	Address string
}

func (u UrlParameters) Intersection(u_cur UrlParameters) UrlParameters {
	/*return UrlParameters{
		Name: u.Name + "," + u_cur.Name,
	}*/
	return u
}

func (u UrlParameters) Union(u_cur UrlParameters) UrlParameters {
	if u_cur.Name != "" {
		u.Name = u.Name + "," + u_cur.Name
	}
	if u_cur.Address != "" {
		u.Address = u.Address + "," + u_cur.Address
	}
	return u
}

type FhirValue struct {
	Value string
}

func (f FhirValue) Contains() struct {
	Value func(v string) UrlParameters
} {
	return struct {
		Value func(v string) UrlParameters
	}{
		Value: func(v string) UrlParameters {
			return UrlParameters{
				Name: v,
			}
		},
	}
}
