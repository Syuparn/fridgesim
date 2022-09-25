package medium

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"

	"github.com/syuparn/fridgesim/pkg/config"
	"github.com/syuparn/fridgesim/pkg/di"
)

func TestCreateIngredient(t *testing.T) {
	setupConfig(t)

	tests := []struct {
		name           string
		req            string
		expectedRes    string
		expectedStatus int
	}{
		{
			"create a new ingredient",
			`{"kind":"carrot","amount":2}`,
			`{"id":"0123456789ABCDEFGHJKMNPQRS","kind":"carrot","amount":2}` + "\n",
			http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()

			// inject
			newDBClient, teardown := mockDBInjector(t)
			newIDGenerator := mockIDInjector("0123456789ABCDEFGHJKMNPQRS")
			defer teardown()
			injector := di.New()
			do.Override(injector, newDBClient)
			do.Override(injector, newIDGenerator)

			// run server
			cfg := do.MustInvoke[*config.Specification](injector)
			e := do.MustInvoke[*echo.Echo](injector)

			// Start server
			defer e.Shutdown(ctx)
			go func() {
				e.Start(fmt.Sprintf(":%d", cfg.Port))
			}()

			// request
			res, err := http.Post(
				fmt.Sprintf("http://localhost:%d/ingredients", cfg.Port),
				"application/json",
				bytes.NewBuffer([]byte(tt.req)),
			)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			// assert
			assert.Equal(t, tt.expectedStatus, res.StatusCode)
			assert.Equal(t, tt.expectedRes, string(resBody))
		})
	}
}
