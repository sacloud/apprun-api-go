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

// アプリケーショントラフィック分散を取得します。
// (GET /applications/{id}/traffics)
func (s *Server) ListApplicationTraffics(c *gin.Context, id string) {
	ts, err := s.Engine.ListTraffics(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, ts)
}

// アプリケーショントラフィック分散を変更します。
// (PUT /applications/{id}/traffics)
func (s *Server) PutApplicationTraffic(c *gin.Context, id string) {
	paramJSON := &v1.PutTrafficsBody{}
	if err := c.ShouldBindJSON(paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ut, err := s.Engine.UpdateTraffic(id, paramJSON)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &ut)
}
