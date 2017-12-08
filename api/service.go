package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"log"
	"web-status/db"
	"web-status/db/model"

	"github.com/gorilla/mux"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

type Service interface {
	PostWebAddress(ctx context.Context, handler db.Handler, url string) error
	GetWebAddress(ctx context.Context, handler db.Handler, id string) (model.Url, error)
	PostResult(ctx context.Context, handler db.Handler, id string, r model.Result) error
	GetResult(ctx context.Context, handler db.Handler, id string) (model.Result, error)
	GetCount(ctx context.Context, handler db.Handler, id string) (model.Check, error)
}

var (
	ErrCorruptData   = errors.New("Data error - Please verify your details are correct")
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

type e interface {
	error() error
}

//Decoders

func decodePostWebAddressRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	log.Println("recieved post decode request")
	vars := mux.Vars(r)
	url, ok := vars["webAddress"]
	if !ok {
		return nil, ErrBadRouting
		log.Println("returned post decode request failed")
	}
	if !IsValidUrl(url) {
		return nil, ErrCorruptData
	} else {
		log.Println("returned post decode request")
		return postWebAddressRequest{Url: url}, nil
	}
}

//Encoders

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("recieved post encode request")
	if e, ok := response.(e); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		log.Println("recieved post encode error request")
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	log.Println("recieved post encode success request")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrCorruptData:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
