package api

import (
	"net/http"
	"web-status/db"

	"github.com/gorilla/mux"
)

func Run(endpoint string, db db.Handler) error {
	r := mux.NewRouter()
	RunAPIOnRouter(r, db)
	return http.ListenAndServe(endpoint, r)
}

func RunAPIOnRouter(r *mux.Router, db db.Handler) {
	handler := MakeAPIHandler(db)

	apirouter := r.PathPrefix("/v1/").Subrouter()

	apirouter.Methods("GET").Path("/{Operation}/{id}").HandlerFunc(handler.searchHandler)
	apirouter.Methods("POST").PathPrefix("/{Operation}/{value}").HandlerFunc(handler.addHandler)
}
