package employers

type Employee struct {
	FirstName   string       `json:"firstName"`
	MiddleName  string       `json:"middleName"`
	LastName    string       `json:"lastName"`
	Inn         string       `json:"inn"`
	Position    string       `json:"position"`
	Phone       string       `json:"phone"`
	Description string       `json:"description"`
	Attributes  []Attributes `json:"attributes"`
}

type Attributes struct {
	Meta  []Meta `json:"meta"`
	Value string `json:"value"`
}

type Meta struct {
	Href      string `json:"href"`
	Type      string `json:"type"`
	MediaType string `json:"mediaType"`
}

type GetEmployers struct {
	Rows []Row
}

type Row struct {
	Meta      Meta2  `json:"meta"`
	ID        string `json:"id"`
	AccountID string `json:"accountId,omitempty"`
	Owner     struct {
		Meta struct {
			Href         string `json:"href"`
			MetadataHref string `json:"metadataHref"`
			Type         string `json:"type"`
			MediaType    string `json:"mediaType"`
			UUIDHref     string `json:"uuidHref"`
		} `json:"meta"`
	} `json:"owner,omitempty"`
	Shared bool `json:"shared,omitempty"`
	Group  struct {
		Meta struct {
			Href         string `json:"href"`
			MetadataHref string `json:"metadataHref"`
			Type         string `json:"type"`
			MediaType    string `json:"mediaType"`
		} `json:"meta"`
	} `json:"group,omitempty"`
	Updated      string `json:"updated,omitempty"`
	Name         string `json:"name,omitempty"`
	ExternalCode string `json:"externalCode,omitempty"`
	Archived     bool   `json:"archived,omitempty"`
	Created      string `json:"created,omitempty"`
	LastName     string `json:"lastName"`
	FullName     string `json:"fullName"`
	ShortFio     string `json:"shortFio"`
	Phone        string `json:"phone,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	Email        string `json:"email,omitempty"`
	Cashiers     []struct {
		Meta struct {
			Href      string `json:"href"`
			Type      string `json:"type"`
			MediaType string `json:"mediaType"`
		} `json:"meta"`
	} `json:"cashiers,omitempty"`
}

type Meta2 struct {
	Href         string `json:"href"`
	MetadataHref string `json:"metadataHref"`
	Type         string `json:"type"`
	MediaType    string `json:"mediaType"`
	UUIDHref     string `json:"uuidHref"`
}

type UpdateEmployee struct {
	Meta      Meta2  `json:"meta"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName"`
}

type Delete struct {
	Meta Meta2 `json:"meta"`
}

type tokenResponse struct {
	Token string `json:"access_token"`
}

func (r *Row) GetID() string {
	return r.ID
}
