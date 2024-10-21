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
