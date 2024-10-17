package v1_test

import (
	"net/http/httptest"

	"github.com/sacloud/apprun-api-go/fake"
	"github.com/sacloud/apprun-api-go/fake/server"
)

func init() {
	initFakeServer()
}

func initFakeServer() {
	fakeServer := &server.Server{
		Engine: &fake.Engine{
			User: &fake.User{
				ID:   fake.USER_ID,
				NAME: fake.USER_NAME,
			},
		},
	}
	sv := httptest.NewServer(fakeServer.Handler())
	serverURL = sv.URL
}
