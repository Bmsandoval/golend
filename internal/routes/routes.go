package routing

import (
	// External Dependency
	"github.com/gorilla/mux"
	// Controllers
	c "lendr/internal/controllers"
	"lendr/internal/middleware"
	// Middleware
)

func Register() *mux.Router {
	app := mux.NewRouter()
	// Apply logging middleware
	app.Use(middleware.LogRequests)
	// *******************
	// PUBLIC Endpoints
	// *******************
	app.Methods("GET").Handler(
		RegisterPublicEndpoints())
	// *******************
	// PRIVATE Endpoints
	// *******************
	app.Methods("POST").
		Handler(middleware.ValidateSlackRequests(
			RegisterPrivateEndpoints()))
	return app
}

func RegisterPublicEndpoints() *mux.Router {
	app := mux.NewRouter()
	app.Path("/auth").HandlerFunc(c.InitializeTeamSetupProcedure)
	app.Path("/auth/setup").HandlerFunc(c.CollectAuthToken)
	return app
}

func RegisterPrivateEndpoints() *mux.Router {
	app := mux.NewRouter()
	// *******************
	// MIGRATIONS
	// *******************
	app.Path("/lendrs/migrate/").
		HandlerFunc(c.MigrateLendrs)
	app.Path("/lendrs/revert/").
		HandlerFunc(c.RevertLendrs)
	app.Path("/lendables/migrate/").
		HandlerFunc(c.MigrateLendables)
	app.Path("/lendables/revert/").
		HandlerFunc(c.RevertLendables)
	// *******************
	// LENDRS
	// *******************
	app.Path("/lendrs/register/{LendrId}").
		HandlerFunc(c.RegisterLendr)
	app.Path("/lendrs/unregister/{LendrId}").
		HandlerFunc(c.UnregisterLendr)
	app.Path("/lendrs/list").
		HandlerFunc(c.ShowAllLendrs)
	// *******************
	// MANAGE LENDABLES
	// *******************
	app.Path("/manage/index").
		HandlerFunc(c.ManagementIndex)
	return app
}
