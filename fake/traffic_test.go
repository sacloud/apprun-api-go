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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEngine_Traffic(t *testing.T) {
	t.Run("list traffics", func(t *testing.T) {
		engine := NewEngine()
		req := postApplicationBody()
		createdApp, err := engine.CreateApplication(req)
		require.NoError(t, err)

		resp, err := engine.ListTraffics(*createdApp.Id)
		require.NoError(t, err)

		respJson, err := json.Marshal(resp)
		require.NoError(t, err)

		expectedJSON := `
		{
			"meta": null,
			"data": [
				{
					"version_name": "",
					"is_latest_version": true,
					"percent": 100
				}
			]
		}`
		require.JSONEq(t, expectedJSON, string(respJson))
	})
}
