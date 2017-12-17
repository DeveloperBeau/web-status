package api

import (
	"context"
	"log"
	"web-status/db"

	"github.com/go-kit/kit/endpoint"
)

// MakePostProfileEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakePostWebAddressEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postWebAddressRequest)
		e := s.PostWebAddress(ctx, handler, req.Url)
		return postWebAddressResponse{Err: e}, nil
	}
}

// MakePutProfileEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakePostResultEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return request, nil
	}
}

// MakePatchProfileEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetResultEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return request, nil
	}
}

// MakeDeleteProfileEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetCountEndpoint(s Service, handler db.Handler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return request, nil
	}
}
