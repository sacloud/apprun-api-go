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
	defer engine.rLock()()

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
