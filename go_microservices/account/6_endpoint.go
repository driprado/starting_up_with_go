// Endpoint Structure contains the methods we want to expose to the public
// CreateUser and GetUser are converted into endpoint.Endpoint type from upstream package

package account

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreareUser endpoint.Endpoint // Type Endpoint is a function that takes a context and an interface and returns an interface or an error.
	GetUser    endpoint.Endpoint // Each endpoint represents a single RPC method
}

func MakeEndpoints(s Service) Endpoints { // Takes a Service interface and returns an Endpoints structure
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s), // Convert CreateUser method into endpoint
		GetUser:    makeGetUserEndpoint(s),	// // Convert GetUser method into endpoint
}

// makeCreateUserEndpoint takes a Service interface and converts it's method "CreateUser into an endpoint (which is an item inside the Endpoint structure)"
func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) { // Anonymus function that works as a wrapper for our existing methods, since endpoint.Endpoint is a function, it needs to return a function
		req := request.(CreateUserRequest) // request interface{} from above function converted into type CreateUserRequest (structure containing email and password) defined in 5_reqresp.go
		ok, err := s.CreateUser(ctx, req.Email, req.Password) // since req is of type CreateUserRequest, it has Email and Password fields
		return CreateUserResponse{OK: ok}, err
	 } // If no error, return a CreateUserResponse struct with OK value ok
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint { // Takes a Service interface, return and Endpoint for the service method
	return func (ctx context.Context, request  interface{}) (interface{}, error) {// since Endpoint is a function it has to return a function
		req := request.(GetUserRequest) // request interface above transformed in to GetUserRequest(structure)
		email, err := s.GetUser(ctx, req.Id)
		return GetUserResponse{
			Email: email,
		}, err
		}
}