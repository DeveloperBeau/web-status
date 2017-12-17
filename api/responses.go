package api

import (
	"web-status/db/model"
)

type postWebAddressRequest struct {
	Url string
}

type postWebAddressResponse struct {
	Err error `json:"err,omitempty"`
}

func (r postWebAddressResponse) error() error { return r.Err }

type getResultRequest struct {
	ID string
}

type getResultResponse struct {
	R   model.Result `json:"result,omitempty"`
	Err error        `json:"err,omitempty"`
}

func (r getResultResponse) error() error { return nil }

type postResultRequest struct {
	ID  string
	Url string
}

type postResultResponse struct {
	Rid string `json:"result_id,omitempty"`
	Err error `json:"err,omitempty"`
}

func (r postResultResponse) error() error { return nil }

type getCountRequest struct {
	Url string
}

type getCountResponse struct {
	C   model.CheckResponse `json:"count"`
	Err error               `json:"err,omitempty"`
}

func (r getCountResponse) error() error { return nil }
