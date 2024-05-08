package api

// StringTuple is a struct holding two string elements used by the Keyfactor
// Go Client library for data types requiring a tuple of strings
type StringTuple struct {
	Elem1 string `json:"elem1,omitempty"`
	Elem2 string `json:"elem2,omitempty"`
}

// apiHeaders is a struct that holds an array of StringTuples used
// to modularize the passing of custom API headers.
type apiHeaders struct {
	Headers []StringTuple `json:"headers,omitempty"`
}

// apiQuery is a struct that holds an array of StringTuples used
// to modularize the passing of custom URL query parameters.
type apiQuery struct {
	Query []StringTuple `json:"query,omitempty"`
}

// request is a structure that holds required information for communicating with
// the Keyfactor API. Included inside this struct is a pointer to an APIHeaders struct, a payload as an
// interface, and other configuration information for the API call.
type request struct {
	Method   string      `json:"method"`
	Endpoint string      `json:"endpoint"`
	Headers  *apiHeaders `json:"headers,omitempty"`
	Query    *apiQuery   `json:"query,omitempty"`
	Payload  interface{} `json:"payload,omitempty"`
}
