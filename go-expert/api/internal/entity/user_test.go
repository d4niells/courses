package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Run("create user", func(t *testing.T) {
		user, err := NewUser("John Doe", "johndoe@email.com", "123456")
		assert.NoError(t, err)
		assert.NotEmpty(t, user.ID)
		assert.NotEmpty(t, user.Password)
		assert.Equal(t, "John Doe", user.Name)
		assert.Equal(t, "johndoe@email.com", user.Email)
	})

	t.Run("validate user password", func(t *testing.T) {
		user, err := NewUser("John Doe", "johndoe@email.com", "123456")
		assert.Nil(t, err)
		assert.True(t, user.validatePassword("123456"))
		assert.NotEqual(t, "123456", user.Password)
	})
}
