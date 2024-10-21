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
	"os"

	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/apprun-api-go/apis/v1"
	"github.com/sacloud/apprun-api-go/fake"
)

var _ v1.ServerInterface = (*Server)(nil)

type Server struct {
	Engine *fake.Engine
}

func (s *Server) Handler() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	if os.Getenv("APPRUN_SERVER_DEBUG") != "" {
		gin.SetMode(gin.DebugMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())
	if os.Getenv("APPRUN_SERVER_LOGGING") != "" {
		engine.Use(gin.Logger())
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	v1.RegisterHandlers(engine, s)
	return engine
}
