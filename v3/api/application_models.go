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

// ApplicationScheduleInterval defines an interval-based inventory schedule.
type ApplicationScheduleInterval struct {
	Minutes int `json:"Minutes"`
}

// ApplicationScheduleDaily defines a daily time-based inventory schedule.
type ApplicationScheduleDaily struct {
	Time string `json:"Time"` // ISO 8601 datetime string (e.g. "2023-11-25T23:30:00Z")
}

// ApplicationSchedule holds the schedule configuration for an application.
// Set exactly one of Interval or Daily; omit both to disable the schedule.
type ApplicationSchedule struct {
	Interval *ApplicationScheduleInterval `json:"Interval,omitempty"`
	Daily    *ApplicationScheduleDaily    `json:"Daily,omitempty"`
}

// ApplicationCertStore is a minimal certificate store reference within an application detail response.
type ApplicationCertStore struct {
	Id string `json:"Id"` // Store GUID (UUID)
}

// ApplicationListItem represents one entry returned by GET /Applications (list endpoint).
// The Schedule field is returned as a cron expression string by the list endpoint.
type ApplicationListItem struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	Schedule string `json:"Schedule"`
}

// ApplicationResponse is the full application detail returned by GET /Applications/{id}.
type ApplicationResponse struct {
	Id                 int                    `json:"Id"`
	Name               string                 `json:"Name"`
	OverwriteSchedules bool                   `json:"OverwriteSchedules"`
	Schedule           *ApplicationSchedule   `json:"Schedule,omitempty"`
	CertificateStores  []ApplicationCertStore `json:"CertificateStores,omitempty"`
}

// ApplicationCreateRequest is the request body for POST /Applications.
type ApplicationCreateRequest struct {
	Name               string               `json:"Name"`
	OverwriteSchedules bool                 `json:"OverwriteSchedules"`
	Schedule           *ApplicationSchedule `json:"Schedule,omitempty"`
}

// ApplicationUpdateRequest is the request body for PUT /Applications/{id}.
// The Id field is set automatically by UpdateApplication.
type ApplicationUpdateRequest struct {
	Id                 int                  `json:"Id"`
	Name               string               `json:"Name"`
	OverwriteSchedules bool                 `json:"OverwriteSchedules"`
	Schedule           *ApplicationSchedule `json:"Schedule,omitempty"`
}
