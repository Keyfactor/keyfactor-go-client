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

import "time"

type WorkflowCertificate struct {
	Id                   int               `json:"Id"`
	CARequestId          string            `json:"CARequestId"`
	CommonName           string            `json:"CommonName"`
	DistinguishedName    string            `json:"DistinguishedName"`
	SubmissionDate       time.Time         `json:"SubmissionDate"`
	CertificateAuthority string            `json:"CertificateAuthority"`
	Template             string            `json:"Template"`
	Requester            string            `json:"Requester"`
	State                int               `json:"State"`
	StateString          string            `json:"StateString"`
	Metadata             map[string]string `json:"Metadata"`
}

type WorkflowActionResponse struct {
	Failures []struct {
		CARowId            int    `json:"CARowId"`
		CARequestId        string `json:"CARequestId"`
		CAHost             string `json:"CAHost"`
		CALogicalName      string `json:"CALogicalName"`
		KeyfactorRequestId int    `json:"KeyfactorRequestId"`
		Comment            string `json:"Comment"`
	} `json:"Failures"`
	Denials []struct {
		CARowId            int    `json:"CARowId"`
		CARequestId        string `json:"CARequestId"`
		CAHost             string `json:"CAHost"`
		CALogicalName      string `json:"CALogicalName"`
		KeyfactorRequestId int    `json:"KeyfactorRequestId"`
		Comment            string `json:"Comment"`
	} `json:"Denials"`
	Successes []struct {
		CARowId            int    `json:"CARowId"`
		CARequestId        string `json:"CARequestId"`
		CAHost             string `json:"CAHost"`
		CALogicalName      string `json:"CALogicalName"`
		KeyfactorRequestId int    `json:"KeyfactorRequestId"`
		Comment            string `json:"Comment"`
	} `json:"Successes"`
}

type WorkflowDenyCertificateRequest struct {
	Comment               string `json:"Comment"`
	CertificateRequestIds []int  `json:"CertificateRequestIds"`
}
