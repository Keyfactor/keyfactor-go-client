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
