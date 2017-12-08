package api

import (
	"web-status/db"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	PostWebAddressEndpoint endpoint.Endpoint
	GetWebAddressEndpoint  endpoint.Endpoint
	GetResultEndpoint      endpoint.Endpoint
	PostResultEndpoint     endpoint.Endpoint
	GetCountEndpoint       endpoint.Endpoint
}

func MakeServerEndpoints(s Service, handler db.Handler) Endpoints {
	return Endpoints{
		PostWebAddressEndpoint: MakePostWebAddressEndpoint(s, handler),
		GetWebAddressEndpoint:  MakeGetWebAddressEndpoint(s, handler),
		GetResultEndpoint:      MakeGetResultEndpoint(s, handler),
		PostResultEndpoint:     MakePostResultEndpoint(s, handler),
		GetCountEndpoint:       MakeGetCountEndpoint(s, handler),
	}
}
