package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser ユーザー情報の取得
// (GET /user)
func (s *Server) GetUser(c *gin.Context) {
	// TODO: handle error
	err := s.Engine.GetUser()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)
}

// PostUser ユーザーの作成
// (POST /user)
func (s *Server) PostUser(c *gin.Context) {
	// TODO: handle error
	err := s.Engine.CreateUser()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.Status(http.StatusCreated)
}
