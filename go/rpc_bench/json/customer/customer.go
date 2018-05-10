package customer

type CustomerRequest struct {
	Id        int32                      `json:"id,omitempty"`
	Name      string                     `json:"name,omitempty"`
	Email     string                     `json:"email,omitempty"`
	Phone     string                     `json:"phone,omitempty"`
	Addresses []*CustomerRequest_Address `json:"addresses,omitempty"`
}

type CustomerRequest_Address struct {
	Street            string `json:"street,omitempty"`
	City              string `json:"city,omitempty"`
	State             string `json:"state,omitempty"`
	Zip               string `json:"zip,omitempty"`
	IsShippingAddress bool   `json:"isShippingAddress,omitempty"`
}

type CustomerResponse struct {
	Id      int32 `json:"id,omitempty"`
	Success bool  `json:"success,omitempty"`
}

type CustomerFilter struct {
	Keyword string `json:"keyword,omitempty"`
}
