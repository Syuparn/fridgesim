package config

import (
	"golang.org/x/xerrors"

	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	Port int
}

func New() (*Specification, error) {
	var s Specification
	// NOTE: env FRIDGESIM_FOO is set to s.Foo
	if err := envconfig.Process("fridgesim", &s); err != nil {
		return nil, xerrors.Errorf("failed to read config from environment variable: %w", err)
	}

	return &s, nil
}
