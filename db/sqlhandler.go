package db

import (
	"database/sql"
	"web-status/db/model"
)

type SQLHandler struct {
	*sql.DB
}

func (handler *SQLHandler) AddUrl(model.Url) /*error*/ {

}

func (handler *SQLHandler) AddResult(model.Url, bool) /*error*/ {

}
func (handler *SQLHandler) updateCheck(model.Url) /*error*/ {

}
