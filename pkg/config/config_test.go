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
				"FRIDGESIM_PORT":       "8080",
				"FRIDGESIM_DBHOST":     "db",
				"FRIDGESIM_DBPORT":     "5432",
				"FRIDGESIM_DBUSER":     "postgres",
				"FRIDGESIM_DBPASSWORD": "pass",
			},
			&Specification{
				Port:       8080,
				DBHost:     "db",
				DBPort:     5432,
				DBUser:     "postgres",
				DBPassword: "pass",
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
			"an environment variable is invalid",
			map[string]string{
				"FRIDGESIM_PORT":       "invalid",
				"FRIDGESIM_DBHOST":     "db",
				"FRIDGESIM_DBPORT":     "5432",
				"FRIDGESIM_DBUSER":     "postgres",
				"FRIDGESIM_DBPASSWORD": "pass",
			},
		},
		{
			"an environment variable is empty",
			map[string]string{
				"FRIDGESIM_PORT":       "",
				"FRIDGESIM_DBHOST":     "db",
				"FRIDGESIM_DBPORT":     "5432",
				"FRIDGESIM_DBUSER":     "postgres",
				"FRIDGESIM_DBPASSWORD": "pass",
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
