package app

import (
	"fmt"
	"net/http"
)

func (app *App) createEmployee(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello World")
}
