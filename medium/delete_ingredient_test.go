package medium

import (
	"context"
	"fmt"
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

func TestDeleteIngredient(t *testing.T) {
	setupConfig(t)

	tests := []struct {
		name           string
		id             string
		dbRows         []*dbSchema
		expectedStatus int
	}{
		{
			"delete an ingredient",
			"0123456789ABCDEFGHJKMNPQRS",
			[]*dbSchema{
				{"0123456789ABCDEFGHJKMNPQRS", "carrot", 2.0},
			},
			http.StatusNoContent,
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
			c := &http.Client{}
			req, err := http.NewRequest(
				http.MethodDelete,
				fmt.Sprintf("http://localhost:%d/ingredients/%s", cfg.Port, tt.id),
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			res, err := c.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			// assert
			assert.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}
