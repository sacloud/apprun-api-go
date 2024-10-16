package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	server     = &Server{}
	testServer *httptest.Server
)

func TestMain(m *testing.M) {
	testServer = httptest.NewServer(server.Handler())
	defer testServer.Close()

	os.Exit(m.Run())
}

func TestServer_ping(t *testing.T) {
	resp, err := http.Get(testServer.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, "pong", string(body))
}
