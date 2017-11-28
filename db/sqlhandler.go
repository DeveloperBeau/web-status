package db

import (
	"database/sql"
	"fmt"
	"log"
	"web-status/db/model"
)

type SQLHandler struct {
	*sql.DB
}

func (handler *SQLHandler) AddUrl(u model.Url) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into url (url) values ('%s')", u.Url))
	return err
}

func (handler *SQLHandler) AddResult(u model.Url, result bool) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into result (url_id,result) values ('%d','%t')", u.Id, result))
	return err
}
func (handler *SQLHandler) updateCheck(u model.Url) error {
	row := handler.DB.QueryRow("select * from check where url_id = $1", u.Id)
	c := model.Check{}
	err := row.Scan(&c.Id, &c.Url_id, &c.Count)
	if err != nil {
		log.Println(err)
	}
	return err
}
