package routes

import (
	"encoding/json"
	"strings"

	//"fmt"
	"github.com/gorilla/mux"
	"github.com/nlopes/slack"
	"golend/internal/controllers/database"
	"golend/internal/controllers/lenders"
	"golend/internal/controllers/manage"
	"golend/internal/controllers/registration"
	"golend/pkg/middleware"
	"log"
	"net/http"
)

func Register() *mux.Router {
	app := mux.NewRouter()
	// Apply logging middleware
	app.Use(middleware.LogRequests)
	// *******************
	// PUBLIC Endpoints
	// *******************
	app.Methods("GET").Handler(
		registerPublicEndpoints())
	// *******************
	// PRIVATE Endpoints
	// *******************
	app.Methods("POST").
		Handler(middleware.ValidateSlackRequests(
			registerPrivateEndpoints()))
	return app
}

func registerPublicEndpoints() *mux.Router {
	app := mux.NewRouter()
	app.Path("/registration/begin").HandlerFunc(registration.EnrollmentUrl)
	app.Path("/registration/finish").HandlerFunc(registration.CollectAuthTokens)
	return app
}

func registerPrivateEndpoints() *mux.Router {
	app := mux.NewRouter()
	// *******************
	// MIGRATIONS
	// *******************
	app.Path("/database/migrate/").
		HandlerFunc(database.Migrate)
	app.Path("/database/revert/").
		HandlerFunc(database.Revert)
	// *******************
	// LENDRS
	// *******************
	app.Path("/lenders/show_all").
		HandlerFunc(lenders.ShowAll)
	// *******************
	// MANAGE LENDABLES
	// *******************
	app.Path("/manage/index").
		HandlerFunc(manage.Index)
	// *******************
	// CALLBACK ROUTES
	// *******************
	app.Path("/interactivity").
		HandlerFunc(interactivity)
	app.Path("/dynamic_menus").
		HandlerFunc(dynamicMenus)

	return app
}

func interactivity(w http.ResponseWriter, r *http.Request) {
	var requestValues slack.InteractionCallback
	var err error = nil
	err = r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}
	str := r.Form["payload"][0]
	err = json.Unmarshal([]byte(str), &requestValues)
	if err != nil {
		log.Fatal(err)
		return
	}

	if strings.HasPrefix(requestValues.CallbackID, "manage.") {
		manage.CallbackHandler(requestValues, w, r)
	}
}
func dynamicMenus(w http.ResponseWriter, r *http.Request) {
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	log.Printf( "%s", s )

}
