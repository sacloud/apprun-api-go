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
	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

// アプリケーション一覧を取得します。
// (GET /applications)
func (s *Server) ListApplications(c *gin.Context, params v1.ListApplicationsParams) {

}

// アプリケーションを削除します。
// (DELETE /applications/{id})
func (s *Server) DeleteApplication(c *gin.Context, id string) {

}

// アプリケーション詳細を取得します。
// (GET /applications/{id})
func (s *Server) GetApplication(c *gin.Context, id string) {

}

// アプリケーションを部分的に変更します。
// (PATCH /applications/{id})
func (s *Server) PatchApplication(c *gin.Context, id string) {

}

// アプリケーションステータスを取得します。
// (GET /applications/{id}/status)
func (s *Server) GetApplicationStatus(c *gin.Context, id string) {

}

// アプリケーショントラフィック分散を取得します。
// (GET /applications/{id}/traffics)
func (s *Server) ListApplicationTraffics(c *gin.Context, id string) {

}

// アプリケーショントラフィック分散を変更します。
// (PUT /applications/{id}/traffics)
func (s *Server) PutApplicationTraffic(c *gin.Context, id string) {

}

// アプリケーションバージョン一覧を取得します。
// (GET /applications/{id}/versions)
func (s *Server) ListApplicationVersions(c *gin.Context, id string, params v1.ListApplicationVersionsParams) {

}

// アプリケーションバージョンを削除します。
// (DELETE /applications/{id}/versions/{version_id})
func (s *Server) DeleteApplicationVersion(c *gin.Context, id string, versionId string) {

}

// アプリケーションバージョン詳細を取得します。
// (GET /applications/{id}/versions/{version_id})
func (s *Server) GetApplicationVersion(c *gin.Context, id string, versionId string) {

}
