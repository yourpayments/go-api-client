package ypmn

type Billing struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	CountryCode string `json:"countryCode"`
	City string `json:"city"`
	State string `json:"state"`
	CompanyName string `json:"companyName"`
	TaxId string `json:"taxId"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	ZipCode string `json:"zipCode"`
}

