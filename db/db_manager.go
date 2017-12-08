package db

import "web-status/db/model"

type Handler interface {
	AddUrl(model.Url) error
	// AddResult(u model.Url, result bool) error
	// updateCheck(model.Url) error
}

//Database Handler factory function
func MakeDatabaseHandler(connection string) (Handler, error) {
	return NewSQLHandler(connection)
}
