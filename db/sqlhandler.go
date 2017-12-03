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

// adds a URL to the database and sets up count
func (handler *SQLHandler) AddUrl(u model.Url) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into url (url) values ('%s')", u.Url))
	if err == nil {
		_, err := handler.Exec(fmt.Sprintf("Insert into count (url,count) values ('%s',0)", u.Url))
		return err
	}
	return err
}

// adds a result to the database and updates count related to the result
func (handler *SQLHandler) AddResult(u model.Url, result bool) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into result (url_id,result) values ('%d','%t')", u.Id, result))
	if err == nil {
		handler.updateCheck(u)
	}
	return err
}

//Checks if Url exists
func (handler *SQLHandler) getUrl(u string) *model.Url {
	var url model.Url
	err := handler.DB.QueryRow("select * from url where url = $1", u).Scan(&url.Id, &url.Url)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		} else {
			log.Fatal(err)
		}
	}
	return &url
}

func (handler *SQLHandler) updateCheck(u model.Url) error {
	row := handler.DB.QueryRow("select * from check where url_id = $1", u.Id)
	c := model.Check{}
	err := row.Scan(&c.Id, &c.Url_id, &c.Count)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Success")
	}
	return err
}
