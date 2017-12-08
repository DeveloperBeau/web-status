package api

import (
	"flag"
	"net/http"
	"web-status/db"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log"
)

var s Service

func Run(endpoint string, db db.Handler) error {
	r := RunAPIOnRouter(db)
	return http.ListenAndServe(endpoint, r)
}

func RunAPIOnRouter(db db.Handler) http.Handler {
	log.Println("starting api router")
	r := mux.NewRouter()
	e := MakeServerEndpoints(s, db)
	flag.Parse()

	//apirouter := r.PathPrefix("/dev/").Subrouter()

	r.Methods("POST").Path("/addUrl/{webAddress}").Handler(httptransport.NewServer(
		e.PostWebAddressEndpoint,
		decodePostWebAddressRequest,
		encodeResponse,
	))

	return r
}
