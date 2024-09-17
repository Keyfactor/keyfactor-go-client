// Copyright 2024 Keyfactor
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

// StringTuple is a struct holding two string elements used by the Keyfactor
// Go Client library for data types requiring a tuple of strings
type StringTuple struct {
	Elem1 string
	Elem2 string
}

// apiHeaders is a struct that holds an array of StringTuples used
// to modularize the passing of custom API headers.
type apiHeaders struct {
	Headers []StringTuple
}

// apiQuery is a struct that holds an array of StringTuples used
// to modularize the passing of custom URL query parameters.
type apiQuery struct {
	Query []StringTuple
}

// request is a structure that holds required information for communicating with
// the Keyfactor API. Included inside this struct is a pointer to an APIHeaders struct, a payload as an
// interface, and other configuration information for the API call.
type request struct {
	Method   string
	Endpoint string
	Headers  *apiHeaders
	Query    *apiQuery
	Payload  interface{}
}
