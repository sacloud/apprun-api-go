package server

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

// アプリケーション一覧を取得します。
// (GET /applications)
func (s *Server) ListApplications(c *gin.Context, params v1.ListApplicationsParams) {

}

// アプリケーションを作成します。
// (POST /applications)
func (s *Server) PostApplication(c *gin.Context) {

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

// (GET /user)
func (s *Server) GetUser(c *gin.Context) {

}

// (POST /user)
func (s *Server) PostUser(c *gin.Context) {

}
