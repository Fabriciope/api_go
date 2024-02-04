package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Fabrício", "fabricio@gmail.com", "password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Fabrício", user.Name)
	assert.Equal(t, "fabricio@gmail.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("Fabrício", "fabricio@gmail.com", "password")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("password"))
	assert.False(t, user.ValidatePassword("wrong-password"))
	assert.NotEqual(t, "password", user.Password)
}

func TestGetData(t *testing.T) {
	user, err := NewUser("Fabrício", "fabricio@gmail.com", "password")
	assert.Nil(t, err)

	expectedFields := map[string]interface{}{
		"id": user.ID, 
		"name": user.Name, 
		"email": user.Email, 
		"password": user.Password, 
		"created_at": user.CreatedAt,
	}
	data := user.DataForDB()
	for field, value := range expectedFields {
		valueData, ok := data[field]
		assert.True(t, ok)
		assert.Equal(t, value, valueData)
	}
}