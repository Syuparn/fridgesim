package infrastructure

import (
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/syuparn/fridgesim/ent"
)

func NewClient(db *sql.DB) *ent.Client {
	drv := entsql.OpenDB("postgres", db)
	return ent.NewClient(ent.Driver(drv))
}
