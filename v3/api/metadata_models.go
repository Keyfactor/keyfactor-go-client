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

// UpdateMetadataArgs holds the function arguments used for calling the UpdateMetadata method.
type UpdateMetadataArgs struct {
	CertID              int                    `json:"Id"`
	CertificateMetadata []StringTuple          `json:"-"`
	Metadata            map[string]interface{} `json:"Metadata"`
	CollectionId        int                    `json:"CollectionId"`
}

type MetadataField struct {
	Id           int    `json:"Id"`
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	DataType     int    `json:"DataType"`
	Hint         string `json:"Hint"`
	Validation   string `json:"Validation"`
	Enrollment   int    `json:"Enrollment"`
	Message      string `json:"Message"`
	Options      string `json:"Options"`
	DefaultValue string `json:"DefaultValue"`
	DisplayOrder int    `json:"DisplayOrder"`
}
