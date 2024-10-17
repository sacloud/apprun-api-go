package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser ユーザー情報の取得
// (GET /user)
func (s *Server) GetUser(c *gin.Context) {
	// TODO: handle error
	s.Engine.GetUser()

	c.Status(http.StatusOK)
}

// PostUser ユーザーの作成
// (POST /user)
func (s *Server) PostUser(c *gin.Context) {
	// TODO: handle error
	s.Engine.CreateUser()

	c.Status(http.StatusCreated)
}
