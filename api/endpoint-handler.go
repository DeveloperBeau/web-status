package api

import "web-status/db"

type APIHandler struct {
	handler db.Handler
}

//API Handler Factory Function
func MakeAPIHandler(db db.Handler) *APIHandler {
	return &APIHandler{
		handler: db,
	}
}
