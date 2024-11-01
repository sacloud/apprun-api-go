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
	versionDefaultPageNum   = 1
	versionDefaultPageSize  = 50
	versionDefaultSortField = "created_at"
	versionDefaultSortOrder = v1.ListApplicationVersionsParamsSortOrderDesc
)

// アプリケーションバージョン一覧を取得します。
// (GET /applications/{id}/versions)
func (s *Server) ListApplicationVersions(c *gin.Context, id string, params v1.ListApplicationVersionsParams) {
	// クエリパラメーターのデフォルト値のセット
	if params.PageNum == nil {
		params.PageNum = &versionDefaultPageNum
	}
	if params.PageSize == nil {
		params.PageSize = &versionDefaultPageSize
	}
	if params.SortField == nil {
		params.SortField = &versionDefaultSortField
	}
	if params.SortOrder == nil {
		params.SortOrder = &versionDefaultSortOrder
	}

	versions, err := s.Engine.ListVersions(id, params)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, versions)
}

// アプリケーションバージョン詳細を取得します。
// (GET /applications/{id}/versions/{version_id})
func (s *Server) GetApplicationVersion(c *gin.Context, id string, versionId string) {
	v, err := s.Engine.ReadVersion(id, versionId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, v)
}

// アプリケーションバージョンを削除します。
// (DELETE /applications/{id}/versions/{version_id})
func (s *Server) DeleteApplicationVersion(c *gin.Context, id string, versionId string) {
	err := s.Engine.DeleteVersion(id, versionId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
