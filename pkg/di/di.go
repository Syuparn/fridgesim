package di

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"

	"github.com/syuparn/fridgesim/adapter"
	"github.com/syuparn/fridgesim/domain"
	"github.com/syuparn/fridgesim/ent"
	"github.com/syuparn/fridgesim/infrastructure"
	"github.com/syuparn/fridgesim/pkg/config"
	"github.com/syuparn/fridgesim/usecase"
)

func New() *do.Injector {
	injector := do.New()

	do.Provide(injector, newConfig)
	do.Provide(injector, NewEntClient)
	do.Provide(injector, newDB)
	do.Provide(injector, newIngredientFactory)
	do.Provide(injector, newIngredientRepository)
	do.Provide(injector, newCreateIngredientInputPort)
	do.Provide(injector, newListIngredientsInputPort)
	do.Provide(injector, newDeleteIngredientInputPort)
	do.Provide(injector, newController)
	do.Provide(injector, newServer)

	return injector
}

func newConfig(i *do.Injector) (*config.Specification, error) {
	return config.New()
}

func newDB(i *do.Injector) (*sql.DB, error) {
	cfg := do.MustInvoke[*config.Specification](i)

	source := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		"fridgesim",
		"disable",
	)

	return sql.Open("postgres", source)
}

func NewEntClient(i *do.Injector) (*ent.Client, error) {
	db := do.MustInvoke[*sql.DB](i)
	return infrastructure.NewClient(db), nil
}

func newIngredientFactory(i *do.Injector) (domain.IngredientFactory, error) {
	return domain.NewIngredientFacotry(nil), nil
}

func newIngredientRepository(i *do.Injector) (domain.IngredientRepository, error) {
	client := do.MustInvoke[*ent.Client](i)
	return infrastructure.NewIngredientRepository(client)
}

func newCreateIngredientInputPort(i *do.Injector) (usecase.CreateIngredientInputPort, error) {
	ingredientFactory := do.MustInvoke[domain.IngredientFactory](i)
	ingredientRepository := do.MustInvoke[domain.IngredientRepository](i)
	return usecase.NewCreateIngredientInputPort(ingredientFactory, ingredientRepository), nil
}

func newListIngredientsInputPort(i *do.Injector) (usecase.ListIngredientsInputPort, error) {
	ingredientRepository := do.MustInvoke[domain.IngredientRepository](i)
	return usecase.NewListIngredientsInputPort(ingredientRepository), nil
}

func newDeleteIngredientInputPort(i *do.Injector) (usecase.DeleteIngredientInputPort, error) {
	ingredientRepository := do.MustInvoke[domain.IngredientRepository](i)
	return usecase.NewDeleteIngredientInputPort(ingredientRepository), nil
}

func newController(i *do.Injector) (*adapter.Controller, error) {
	createIngredientsInputPort := do.MustInvoke[usecase.CreateIngredientInputPort](i)
	listIngredientsInputPort := do.MustInvoke[usecase.ListIngredientsInputPort](i)
	deleteIngredientInputPort := do.MustInvoke[usecase.DeleteIngredientInputPort](i)
	return adapter.NewController(
		createIngredientsInputPort,
		listIngredientsInputPort,
		deleteIngredientInputPort,
	), nil
}

func newServer(i *do.Injector) (*echo.Echo, error) {
	c := do.MustInvoke[*adapter.Controller](i)
	return adapter.NewServer(c), nil
}
