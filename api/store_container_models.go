package api

// CertStoreContainer holds the function arguments used for calling the GetStoreContainers method.
type CertStoreContainer struct {
	Id                 *int   `json:"Id,omitempty"`
	Name               string `json:"Name"`
	OverwriteSchedules bool   `json:"OverwriteSchedules"`
	Schedule           string `json:"Schedule"`
	CertStoreType      int    `json:"CertStoreType"`
}
