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

package fake

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	v1 "github.com/sacloud/apprun-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

func TestEngine_Application(t *testing.T) {
	engine := &Engine{}

	t.Run("create application", func(t *testing.T) {
		server, userName, password := "apprun-example.sakuracr.jp", "apprun", "apprun" //nolint:gosec
		envKey, envValue := "envkey", "envvalue"
		headerName, headerValue := "Custom-Header", "Awesome"
		probe := v1.PostApplicationBodyComponentProbe{
			HttpGet: &struct {
				Headers *[]struct {
					Name  *string "json:\"name,omitempty\""
					Value *string "json:\"value,omitempty\""
				} "json:\"headers,omitempty\""
				Path string "json:\"path\""
				Port int    "json:\"port\""
			}{
				Headers: &[]struct {
					Name  *string `json:"name,omitempty"`
					Value *string `json:"value,omitempty"`
				}{
					{
						Name:  &headerName,
						Value: &headerValue,
					},
				},
				Path: "/healthz",
				Port: 8080,
			},
		}
		req := &v1.PostApplicationBody{
			Name:     "app1",
			Port:     8081,
			MinScale: 1,
			MaxScale: 10,
			Components: []v1.PostApplicationBodyComponent{
				{
					Datasource: v1.PostApplicationBodyComponentDataSource{
						ContainerRegistry: &v1.PostApplicationBodyComponentDataSourceContainerRegistry{
							Image:    "apprun-example.sakuracr.jp/helloworld:latest",
							Server:   &server,
							Username: &userName,
							Password: &password,
						},
					},
					Env: &[]v1.PostApplicationBodyComponentEnv{
						{
							Key:   &envKey,
							Value: &envValue,
						},
					},
					Probe: &probe,
				},
			},
			TimeoutSeconds: 20,
		}

		resp, err := engine.CreateApplication(req)
		require.NoError(t, err)

		respJson, err := json.Marshal(resp)
		require.NoError(t, err)

		expectedJSON := `
		{
			"id": "` + *resp.Id + `",
			"name": "app1",
			"timeout_seconds": 20,
			"port": 8081,
			"min_scale": 1,
			"max_scale": 10,
			"components": [
				{
					"name": "",
					"max_cpu": "",
					"max_memory": "",
					"datasource": {
						"container_registry": {
							"image": "apprun-example.sakuracr.jp/helloworld:latest",
							"server": "apprun-example.sakuracr.jp",
							"username": "apprun"
						}
					},
					"env": [
						{
							"key": "envkey",
							"value": "envvalue"
						}
					],
					"probe": {
						"http_get": {
							"path": "/healthz",
							"port": 8080,
							"headers": [
								{
									"name": "Custom-Header",
									"value": "Awesome"
								}
							]
						}
					}
				}
			],
			"status": "Success",
			"public_url": "` + *resp.PublicUrl + `",
			"created_at": "` + resp.CreatedAt.Format(time.RFC3339) + `"
		}`
		fmt.Print(string(respJson))
		require.JSONEq(t, expectedJSON, string(respJson))
	})
}
