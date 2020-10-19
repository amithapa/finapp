package api

import (
	"net/http"

	v1 "github.com/amithapa/finapp/internal/api/v1"
	"github.com/gorilla/mux"
)

//NewRouter provide a handler API service
func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/version", v1.VersionHandler)

	return router, nil
}
