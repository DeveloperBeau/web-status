package api

import (
	"context"
	"errors"
	"web-status/db/model"
)

type Service interface {
	PostWebAddress(ctx context.Context, w model.Url) error
	GetWebAddress(ctx context.Context, id string) (model.Url, error)
	PostResult(ctx context.Context, id string, r model.Result) error
	GetResult(ctx context.Context, id string) (model.Result, error)
	GetCount(ctx context.Context, id string) (model.Check, error)
}

var (
	ErrCorruptData   = errors.New("data does not fit required json schema")
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)
