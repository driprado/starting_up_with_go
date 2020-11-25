package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// The code must specify that the service will be served over http
// The methods to be exposed must be turned into endpoints.
// to do this structs must be created to represent the request and response for each method.

// All methods and structure items must be public, therefore they all start with capital letters

type (
	// CreateUserRequest represents the CreateUser method's request
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// CreateUserResponse represents the CreateUser method's response
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}
	// GetUserRequest represents the GetUser method's request
	GetUserRequest struct {
		ID string `json:"id"`
	}
	// GetUserResponse represents the GetUser method's response
	GetUserResponse struct {
		Email string `json:"email"`
	}
)

// encodeResponse encodes the response into json
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// decodeUserReq decodes the request from json
func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// decodeEmailReq decodes the email field out of the request method, it is needed because `r.Methos("GET")` in `7_server.go` grabs from a path ("/user/{id}") and not from a peace of jason
func decodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r) // get back a map that contain all variables inside of out path `/user/{id}`
	req = GetUserRequest{
		ID: vars["id"], // use the key `id` from out path variable and put into GetUserRequest struct
	}
	return req, nil
}
