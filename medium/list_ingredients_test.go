package medium

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"

	"github.com/syuparn/fridgesim/ent"
	"github.com/syuparn/fridgesim/pkg/config"
	"github.com/syuparn/fridgesim/pkg/di"
)

func TestListIngredients(t *testing.T) {
	setupConfig(t)

	tests := []struct {
		name           string
		dbRows         []*dbSchema
		expectedRes    string
		expectedStatus int
	}{
		{
			"list ingredients",
			[]*dbSchema{
				{"0123456789ABCDEFGHJKMNPQRS", "carrot", 2.0},
				{"1123456789ABCDEFGHJKMNPQRS", "cabbage", 0.5},
			},
			`
			{
				"ingredients": [
					{"id": "0123456789ABCDEFGHJKMNPQRS", "kind": "carrot", "amount": 2},
					{"id": "1123456789ABCDEFGHJKMNPQRS", "kind": "cabbage", "amount": 0.5}
				]
			}
			`,
			http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()

			// inject
			newDBClient, teardown := mockDBInjector(t)
			defer teardown()
			injector := di.New()
			do.Override(injector, newDBClient)

			// prepare rows
			client := do.MustInvoke[*ent.Client](injector)
			for _, row := range tt.dbRows {
				err := client.Ingredient.Create().
					SetID(row.ID).
					SetKind(row.Kind).
					SetAmount(row.Amount).
					Exec(ctx)
				if err != nil {
					t.Fatal(err)
				}
			}

			// run server
			cfg := do.MustInvoke[*config.Specification](injector)
			e := do.MustInvoke[*echo.Echo](injector)

			// Start server
			defer e.Shutdown(ctx)
			go func() {
				e.Start(fmt.Sprintf(":%d", cfg.Port))
			}()
			time.Sleep(1 * time.Second)

			// request
			res, err := http.Get(
				fmt.Sprintf("http://localhost:%d/ingredients", cfg.Port),
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
			assert.JSONEq(t, tt.expectedRes, string(resBody))
		})
	}
}
