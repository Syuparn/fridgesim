package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		envs     map[string]string
		expected *Specification
	}{
		{
			"environment variables is set to specification",
			map[string]string{
				"FRIDGESIM_PORT": "8080",
			},
			&Specification{
				Port: 8080,
			},
		},
	}

	for _, tt := range tests {
		tt := tt // pin

		// NOTE: t.Run cannot be used because t.SetEnv cannot be used concurrently
		t.Logf("test: %s", tt.name)

		for k, v := range tt.envs {
			t.Setenv(k, v)
		}

		actual, err := New()

		assert.NoError(t, err)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestNewError(t *testing.T) {
	tests := []struct {
		name string
		envs map[string]string
	}{
		{
			"environment variable is invalid",
			map[string]string{
				"FRIDGESIM_PORT": "invalid",
			},
		},
		{
			"environment variable is empty",
			map[string]string{
				"FRIDGESIM_PORT": "",
			},
		},
	}

	for _, tt := range tests {
		tt := tt // pin

		// NOTE: t.Run cannot be used because t.SetEnv cannot be used concurrently
		t.Logf("test: %s", tt.name)

		for k, v := range tt.envs {
			t.Setenv(k, v)
		}

		c, err := New()

		assert.Error(t, err)
		assert.Nil(t, c)
	}
}
