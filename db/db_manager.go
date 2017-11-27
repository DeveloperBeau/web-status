package db

import "web-status/db/model"

type dbHandler interface {
	AddUrl(model.Url)          //error
	AddResult(model.Url, bool) //error
	updateCheck(model.Url)     //error
}

//factory function
func GetDatabaseHandler(connection string) (dbHandler, error) {
	return NewPQHandler(connection)
}
