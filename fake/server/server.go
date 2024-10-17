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
	if os.Getenv("APPRUN_SERVER_DEBUG") != "" {
		engine.Use(gin.Logger())
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	v1.RegisterHandlers(engine, s)
	return engine
}
