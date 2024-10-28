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
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

var (
	defaultPageNum   = 1
	defaultPageSize  = 50
	defaultSortField = "created_at"
	defaultSortOrder = v1.ListApplicationsParamsSortOrderDesc
)

// アプリケーション詳細を取得します。
// (GET /applications/{id})
func (s *Server) GetApplication(c *gin.Context, id string) {
	application, err := s.Engine.ReadApplication(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, &application)
}

// アプリケーション一覧を取得します。
// (GET /applications)
func (s *Server) ListApplications(c *gin.Context, params v1.ListApplicationsParams) {
	// クエリパラメーターのデフォルト値のセット
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
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, applications)
}

// アプリケーションを削除します。
// (DELETE /applications/{id})
func (s *Server) DeleteApplication(c *gin.Context, id string) {
	err := s.Engine.DeleteApplication(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusNoContent, nil)
}

// アプリケーションを作成します。
// (POST /applications)
func (s *Server) PostApplication(c *gin.Context) {
	// デフォルト値のみ予めセットしておく
	paramJSON := &v1.PostApplicationBody{
		TimeoutSeconds: 60,
		MinScale:       0,
		MaxScale:       10,
	}
	if err := c.ShouldBindJSON(paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application, err := s.Engine.CreateApplication(paramJSON)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusCreated, &application)
}

// アプリケーションを部分的に変更します。
// (PATCH /applications/{id})
func (s *Server) PatchApplication(c *gin.Context, id string) {
	paramJSON := &v1.PatchApplicationBody{}
	if err := c.ShouldBindJSON(paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application, err := s.Engine.PatchApplication(id, paramJSON)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, &application)
}

// アプリケーションステータスを取得します。
// (GET /applications/{id}/status)
func (s *Server) GetApplicationStatus(c *gin.Context, id string) {
	application, err := s.Engine.ReadApplication(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	var status v1.HandlerGetApplicationStatusStatus
	message := ""
	switch *application.Status {
	case v1.ApplicationStatusSuccess:
		status = v1.HandlerGetApplicationStatusStatusSuccess
	case v1.ApplicationStatusFail:
		status = v1.HandlerGetApplicationStatusStatusFail
	case v1.ApplicationStatusUnknown:
		status = v1.HandlerGetApplicationStatusStatusFail
	}
	c.JSON(http.StatusOK, v1.HandlerGetApplicationStatusResponse{
		Status:  &status,
		Message: &message,
	})
}
