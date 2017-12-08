package api

import (
	"context"
	"log"
	"web-status/db"
	"web-status/db/model"
)

type serviceHandler struct{}

func (sh serviceHandler) PostWebAddress(ctx context.Context, handler db.Handler, url string) error {
	log.Println("recieved post service request")
	webAddress := model.Url{Url: url}
	err := handler.AddUrl(webAddress)
	if err != nil {
		log.Println(err)
	}
	return err
}
func (sh serviceHandler) GetWebAddress(ctx context.Context, handler db.Handler, id string) (model.Url, error) {
	return model.Url{}, nil
}
func (sh serviceHandler) PostResult(ctx context.Context, handler db.Handler, id string, r model.Result) error {
	return nil
}
func (sh serviceHandler) GetResult(ctx context.Context, handler db.Handler, id string) (model.Result, error) {
	return model.Result{}, nil
}
func (sh serviceHandler) GetCount(ctx context.Context, handler db.Handler, id string) (model.Check, error) {
	return model.Check{}, nil
}
