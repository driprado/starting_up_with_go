package account

import "context"

// User struct to represent a user inside our business logic
type User struct {
	ID       string `json:"id,omitempty"` // since we will be using a http API, we need to convert this data to and from json, omitempty means id can be empty.
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Repository will interface with the DB Layer whereas the interface in `service.go` is for exposing the microservice to transport.
type Repository interface {
	CreateUser(ctx context.Context, user User) error // Takes ctx and a User struct
	GetUser(ctx context.Context, id string) (string, error)
}
