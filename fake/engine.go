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
	"sync"

	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

// Engine Fakeサーバであつかうダミーデータを表す
//
// Serverに渡した後はDataStore内のデータを外部から操作しないこと
// Note: 本来はサイトごとに保持するデータを分離すべきだが、現状だと実質1サイトのみのため分離していない。
type Engine struct {
	// User ユーザー
	User *User

	Applications []*v1.Application

	mu sync.RWMutex
}

func (engine *Engine) lock() func() {
	engine.mu.Lock()
	return engine.mu.Unlock
}

func (engine *Engine) rLock() func() {
	engine.mu.RLock()
	return engine.mu.RUnlock
}
