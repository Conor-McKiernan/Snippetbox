package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Writes error msg and stack trace to errorLog, then sends generic 500 Internal Server Error response to user
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Sends specific status code + corresponding desc to user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency. Simply convenience wrapper around clientError whichs sends 404 Not Found to user
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
