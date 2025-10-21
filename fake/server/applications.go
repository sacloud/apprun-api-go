// Copyright 2021-2024 The sacloud/apprun-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"encoding/json"
	"net/http"

	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

var (
	defaultPageNum   = 1
	defaultPageSize  = 50
	defaultSortField = "created_at"
	defaultSortOrder = v1.ListApplicationsParamsSortOrderDesc
)

// ListApplications returns the list of applications.
// (GET /applications)
func (s *Server) ListApplications(w http.ResponseWriter, r *http.Request, params v1.ListApplicationsParams) {
	if params.PageNum == nil {
		params.PageNum = &defaultPageNum
	}
	if params.PageSize == nil {
		params.PageSize = &defaultPageSize
	}
	if params.SortField == nil {
		params.SortField = &defaultSortField
	}
	if params.SortOrder == nil {
		params.SortOrder = &defaultSortOrder
	}

	applications, err := s.Engine.ListApplications(params)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, applications)
}

// PostApplication creates an application.
// (POST /applications)
func (s *Server) PostApplication(w http.ResponseWriter, r *http.Request) {
	paramJSON := &v1.PostApplicationBody{
		TimeoutSeconds: 60,
		MinScale:       0,
		MaxScale:       10,
	}
	if err := json.NewDecoder(r.Body).Decode(paramJSON); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	application, err := s.Engine.CreateApplication(paramJSON)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusCreated, &application)
}

// GetApplication returns application details.
// (GET /applications/{id})
func (s *Server) GetApplication(w http.ResponseWriter, r *http.Request, id string) {
	application, err := s.Engine.ReadApplication(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, &application)
}

// PatchApplication partially updates an application.
// (PATCH /applications/{id})
func (s *Server) PatchApplication(w http.ResponseWriter, r *http.Request, id string) {
	paramJSON := &v1.PatchApplicationBody{}
	if err := json.NewDecoder(r.Body).Decode(paramJSON); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	application, err := s.Engine.UpdateApplication(id, paramJSON)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, &application)
}

// DeleteApplication deletes an application.
// (DELETE /applications/{id})
func (s *Server) DeleteApplication(w http.ResponseWriter, r *http.Request, id string) {
	if err := s.Engine.DeleteApplication(id); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetApplicationStatus returns application status.
// (GET /applications/{id}/status)
func (s *Server) GetApplicationStatus(w http.ResponseWriter, r *http.Request, id string) {
	application, err := s.Engine.ReadApplication(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	var status v1.HandlerGetApplicationStatusStatus
	message := ""
	switch application.Status {
	case v1.ApplicationStatusHealthy:
		status = v1.HandlerGetApplicationStatusStatusHealthy
	case v1.ApplicationStatusDeploying:
		status = v1.HandlerGetApplicationStatusStatusDeploying
	case v1.ApplicationStatusUnHealthy:
		status = v1.HandlerGetApplicationStatusStatusUnHealthy
	}
	writeJSON(w, http.StatusOK, v1.HandlerGetApplicationStatusResponse{
		Status:  status,
		Message: message,
	})
}
