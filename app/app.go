package app

import (
	"github.com/gorilla/mux"
)

// App registers all the dependencies and initilizes
type App struct {
	Router *mux.Router
}

// Init is initilizing app instance
func (app *App) Init() {
	router := mux.NewRouter()

	// router.Path("/employees").Methods("POST").Handler(http.HandlerFunc(a.createEmployee))
	router.Path("/employees").Methods("POST").HandlerFunc(app.createEmployee)
}

// Run is starts app
func (app *App) Run() {

}
