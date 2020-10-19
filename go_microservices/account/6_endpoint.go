package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct contain the methods to be exposed to public
type Endpoints struct {
	// Endpoint type is a function, takes a context and a request interface, returns a response interface or error
	CreateUser endpoint.Endpoint // convert CreateUser to Endpoint type
	GetUser    endpoint.Endpoint // convert GetUser to Endpoint type
}

// MakeEndpoints takes a service interface and return an Endpoint structure with the methods inside converted from methods to endpoints
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s), // converts CreateUser into endpoint.Endpoint
		GetUser:    makeGetUserEndpoint(s),    // converts GetUser into endpoint.Endpoint
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	// Since Endpoint type is a function, it has to return a function
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)                    // take the request interface coming in from above function and cast it as CreateUserRequest type defined in 5_reqresp.go
		ok, err := s.CreateUser(ctx, req.Email, req.Password) // now we can pass Email and Password from CreateUserRequest struct into the CreateUser method
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)      // take the request interface coming in from above function and cast it as GetUserRequest type defined in 5_reqresp.go
		email, err := s.GetUser(ctx, req.ID) // now we can pass Id CreateUserRequest struct into the GetUser method
		return GetUserResponse{
			Email: email,
		}, err
	}
}
