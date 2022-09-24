package infrastructure

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	// use mock DB
	"github.com/DATA-DOG/go-sqlmock"

	"github.com/syuparn/fridgesim/domain"
)

func TestIngredientRepositoryList(t *testing.T) {
	columns := []string{"id", "kind", "amount"}

	tests := []struct {
		name     string
		query    string
		mockRows [][]driver.Value
		expected []*domain.Ingredient
	}{
		{
			"obtain all ingredients",
			`SELECT DISTINCT "ingredients"."id", "ingredients"."kind", "ingredients"."amount" FROM "ingredients"`,
			[][]driver.Value{
				{"0123456789ABCDEFGHJKMNPQRS", "carrot", 2.0},
				{"1123456789ABCDEFGHJKMNPQRS", "cabbage", 0.5},
			},
			[]*domain.Ingredient{
				{
					ID:     "0123456789ABCDEFGHJKMNPQRS",
					Kind:   "carrot",
					Amount: 2.0,
				},
				{
					ID:     "1123456789ABCDEFGHJKMNPQRS",
					Kind:   "cabbage",
					Amount: 0.5,
				},
			},
		},
		{
			"return empty slice if no ingredients found",
			`SELECT DISTINCT "ingredients"."id", "ingredients"."kind", "ingredients"."amount" FROM "ingredients"`,
			[][]driver.Value{},
			[]*domain.Ingredient{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, teardown := prepareDB(t)
			defer teardown()

			r, err := NewIngredientRepository(db)
			if err != nil {
				t.Fatal(err)
			}

			// mock
			rows := sqlmock.NewRows(columns)
			lo.ForEach(tt.mockRows, func(row []driver.Value, _ int) {
				rows.AddRow(row...)
			})
			mock.ExpectQuery(regexp.QuoteMeta(tt.query)).
				WillReturnRows(rows)

			// exec
			actual, err := r.List(context.TODO())

			// check
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func prepareDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	teardown := func() {
		db.Close()
	}

	return db, mock, teardown
}
