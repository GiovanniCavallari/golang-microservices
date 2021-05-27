package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization

	// Execution
	user, err := GetUser(0)

	// Validation
	assert.Nil(t, user, "We were not expecting an user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, "User with id 0 not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
}

func TestGetUserUserFound(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err, "We were not expecting an error")
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "John", user.Firstname)
	assert.EqualValues(t, "Doe", user.Lastname)
	assert.EqualValues(t, "email@email.com", user.Email)
}
