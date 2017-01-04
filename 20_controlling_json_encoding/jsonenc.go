package jsonenc

type Person struct {
	Firstname  string `json:"first"` // should be encoded as 'first'
	Middlename string `json:"middle,omitempty"`// should be encoded as 'middle', and not present if blank
	Lastname   string `json:"last"` // should be encoded as 'last'

	SSID int64 `json:"-"` // should not be encoded

	City    string `json:"city,omitempty"` // should be encoded as 'city' and not present if missing
	Country string `json:"country"` // should be encoded as 'country'

	Telephone int64 `json:"tel,string"` // should be encoded as 'tel', the value should be a string, not a number
}

