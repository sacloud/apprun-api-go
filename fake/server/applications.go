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

// アプリケーション詳細を取得します。
// (GET /applications/{id})
func (s *Server) GetApplication(c *gin.Context, id string) {
	application, err := s.Engine.ReadApplication(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, &application)
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
