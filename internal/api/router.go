package api

import (
	"net/http"
)

//NewRouter provide a handler API service
func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()

	return router, nil
}