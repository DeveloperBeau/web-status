package db

import "web-status/db/model"

type Handler interface {
	AddUrl(model.Url) error
	AddResult(model.Url, bool) error
	updateCheck(model.Url) error
	getUrl(u string) (bool, *model.Url)
}

//Database factory function
func MakeDatabaseHandler(connection string) (Handler, error) {
	return NewPQHandler(connection)
}
