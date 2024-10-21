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

var (
	USER_ID   = 111
	USER_NAME = "user_name"
)

// API上でUserの定義は存在しないが、ユーザーのサインアップを管理するために定義する
type User struct {
	ID   int
	NAME string
}

func (engine *Engine) CreateUser() error {
	defer engine.lock()()

	u := engine.User
	if u != nil {
		return newError(
			ErrorTypeConflict, "user", u.ID,
			"さくらのAppRunにユーザーが既に存在します。")
	}

	engine.User = &User{
		ID:   USER_ID,
		NAME: USER_NAME,
	}
	return nil
}

func (engine *Engine) GetUser() error {
	defer engine.rLock()()

	u := engine.User
	if u == nil {
		return newError(
			ErrorTypeNotFound, "user", nil,
			"さくらのAppRunにユーザーが存在しません。")
	}

	return nil
}
