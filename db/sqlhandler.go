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

// adds a URL to the database and sets up count for the web address.
func (handler SQLHandler) AddUrl(u model.Url) error {
	url, err := handler.getUrl(u.Url)
	if url == nil {
		_, err = handler.Exec(fmt.Sprintf("Insert into url (url) values ('%s')", u.Url))
		row, queryErr := handler.getUrl(u.Url)
		if queryErr == nil {
			err = nil
			url = row
		}
	}
	if err == nil {
		handler.addCheckRowForNewUrl(*url)
	}
	return err
}

// adds a result to the database and updates count related to the result
func (handler *SQLHandler) AddResult(u model.Url, result bool) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into result (url_id,result) values ('%d','%t')", u.Id, result))
	if err == nil {
		handler.UpdateCheck(u)
	}
	return err
}

//Checks if Url exists
func (handler *SQLHandler) GetUrl(u string) *model.Url {
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

func (handler *SQLHandler) UpdateCheck(u model.Url) error {
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

func (handler *SQLHandler) getUrl(s string) (*model.Url, error) {
	var url model.Url
	queryErr := handler.DB.QueryRow("select * from url where url = $1", s).Scan(&url.Id, &url.Url, &url.IsEnabled)
	if queryErr == nil {
		return &url, nil
	} else {
		return nil, queryErr
	}
}

func (handler *SQLHandler)addCheckRowForNewUrl(u model.Url) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into count (url_id,count) values ('%d',0)", u.Id))
	return err
}
