package di

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	i := New()
	// if panic is not occurred, DI succeeds
	e := do.MustInvoke[*echo.Echo](i)

	assert.NotNil(t, e)
}
