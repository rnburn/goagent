package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInt32Env(t *testing.T) {
	intVal, ok := getInt32Env("my_int")
	assert.False(t, ok)

	os.Setenv("my_int", "1")
	intVal, ok = getInt32Env("my_int")
	assert.True(t, ok)
	assert.Equal(t, int32(1), intVal)

	os.Setenv("my_int", "a")
	intVal, ok = getInt32Env("my_int")
	assert.False(t, ok)
}

func TestGetBoolEnv(t *testing.T) {
	boolVal, ok := getBoolEnv("my_bool")
	assert.False(t, ok)

	os.Setenv("my_bool", "true")
	boolVal, ok = getBoolEnv("my_bool")
	assert.True(t, ok)
	assert.Equal(t, true, boolVal)

	os.Setenv("my_bool", "false")
	boolVal, ok = getBoolEnv("my_bool")
	assert.True(t, ok)
	assert.Equal(t, false, boolVal)

	os.Setenv("my_bool", "x")
	boolVal, ok = getBoolEnv("my_bool")
	assert.False(t, ok)
}

func TestGetStringEnv(t *testing.T) {
	stringVal, ok := getStringEnv("my_string")
	assert.False(t, ok)

	os.Setenv("my_string", "my_value")
	stringVal, ok = getStringEnv("my_string")
	assert.True(t, ok)
	assert.Equal(t, "my_value", stringVal)
}
