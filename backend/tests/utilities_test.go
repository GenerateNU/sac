package tests

import (
	"github.com/go-playground/validator/v10"
	"testing"

	"github.com/GenerateNU/sac/backend/src/utilities"

	"github.com/huandu/go-assert"
)

func TestThatContainsWorks(t *testing.T) {
	assert := assert.New(t)

	slice := []string{"foo", "bar", "baz"}

	assert.Assert(utilities.Contains(slice, "foo"))
	assert.Assert(utilities.Contains(slice, "bar"))
	assert.Assert(utilities.Contains(slice, "baz"))
	assert.Assert(!utilities.Contains(slice, "qux"))
}

func TestPasswordValidationWorks(t *testing.T) {
	assert := assert.New(t)

	type Thing struct {
		password string `validate:"password"`
	}

	validate := validator.New()
	validate.RegisterValidation("password", utilities.ValidatePassword)

	assert.NilError(validate.Struct(Thing{password: "password!56"}))
	assert.NilError(validate.Struct(Thing{password: "cor+ect-h*rse-batte#ry-stap@le-100"}))
	assert.NilError(validate.Struct(Thing{password: "1!gooood"}))
	assert.Error(validate.Struct(Thing{password: "1!"}))
	assert.Error(validate.Struct(Thing{password: "tooshor"}))
	assert.Error(validate.Struct(Thing{password: "NoSpecialsOrNumbers"}))
}

func TestEmailValidationWorks(t *testing.T) {
	assert := assert.New(t)

	type Thing struct {
		email string `validate:"neu_email"`
	}

	validate := validator.New()
	validate.RegisterValidation("neu_email", utilities.ValidateEmail)

	assert.NilError(validate.Struct(Thing{email: "brennan.mic@northeastern.edu"}))
	assert.NilError(validate.Struct(Thing{email: "blerner@northeastern.edu"}))
	assert.NilError(validate.Struct(Thing{email: "validemail@northeastern.edu"}))
	assert.Error(validate.Struct(Thing{email: "notanortheasternemail@gmail.com"}))
	assert.Error(validate.Struct(Thing{email: "random123@_#!$string"}))
	assert.Error(validate.Struct(Thing{email: "local@mail"}))
	
}
