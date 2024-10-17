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
		engine.CreateUser()
		err := engine.GetUser()
		require.NoError(t, err)
	})
}
