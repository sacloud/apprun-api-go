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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEngine_User(t *testing.T) {
	engine := &Engine{}

	t.Run("create user", func(t *testing.T) {
		err := engine.CreateUser()
		require.NoError(t, err)
		require.Equal(t, engine.User.ID, USER_ID)
		require.Equal(t, engine.User.NAME, USER_NAME)
	})

	t.Run("get user", func(t *testing.T) {
		err := engine.GetUser()
		require.NoError(t, err)
	})
}
