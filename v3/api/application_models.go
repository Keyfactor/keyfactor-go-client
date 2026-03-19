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
// Also reused as the shape for ExactlyOnce.
type ApplicationScheduleDaily struct {
	Time string `json:"Time"` // ISO 8601 datetime string (e.g. "2023-11-25T23:30:00Z")
}

// ApplicationScheduleWeekly defines a weekly inventory schedule.
// Days are weekday names ("Sunday"…"Saturday"); Time is an ISO 8601 UTC datetime.
type ApplicationScheduleWeekly struct {
	Days []string `json:"Days"` // e.g. ["Monday", "Wednesday"]
	Time string   `json:"Time"` // ISO 8601 datetime string
}

// ApplicationScheduleMonthly defines a monthly inventory schedule.
// Day is the day-of-month (1–31); Time is an ISO 8601 UTC datetime.
type ApplicationScheduleMonthly struct {
	Day  int    `json:"Day"`
	Time string `json:"Time"` // ISO 8601 datetime string
}

// ApplicationSchedule holds the schedule configuration for an application.
// Set exactly one field; omit all to disable the schedule (Off).
//
//   - Immediate:   run once immediately (server may convert to ExactlyOnce on next read)
//   - Interval:    run every N minutes
//   - Daily:       run at the same time each day
//   - Weekly:      run on specific weekdays at a given time
//   - Monthly:     run on a specific day of each month at a given time
//   - ExactlyOnce: run exactly once at the specified time
type ApplicationSchedule struct {
	Immediate   *bool                        `json:"Immediate,omitempty"`
	Interval    *ApplicationScheduleInterval `json:"Interval,omitempty"`
	Daily       *ApplicationScheduleDaily    `json:"Daily,omitempty"`
	Weekly      *ApplicationScheduleWeekly   `json:"Weekly,omitempty"`
	Monthly     *ApplicationScheduleMonthly  `json:"Monthly,omitempty"`
	ExactlyOnce *ApplicationScheduleDaily    `json:"ExactlyOnce,omitempty"`
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
