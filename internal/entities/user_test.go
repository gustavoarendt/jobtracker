package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Instance(t *testing.T) {
	user, err := NewUser("Test User", "user@mail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Test User", user.Name)
	assert.Equal(t, "user@mail.com", user.Email)
	assert.NotEmpty(t, user.Password)
	assert.NotNil(t, user.Created_at)
}

func TestUser_Password(t *testing.T) {
	user, err := NewUser("Test User", "user@mail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ComparePassword("123456"))
	assert.False(t, user.ComparePassword("654321"))
	assert.NotEqual(t, user.Password, "123456")
}
