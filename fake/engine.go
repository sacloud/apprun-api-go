package fake

import "sync"

// Engine Fakeサーバであつかうダミーデータを表す
//
// Serverに渡した後はDataStore内のデータを外部から操作しないこと
// Note: 本来はサイトごとに保持するデータを分離すべきだが、現状だと実質1サイトのみのため分離していない。
type Engine struct {
	// User ユーザー
	User *User

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
