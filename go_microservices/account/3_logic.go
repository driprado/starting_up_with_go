package account

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

// Implements our Service interface into a Structure
type service struct {
	repository Repository // interface with DB, defined in "user.go"
	logger     log.Logger // Log things so we can see what is happening inside our microservices
}

// NewService allow us to create a service of type Service defined in service.go
func NewService(rep Repository, logger log.Logger) Service { // Returns a service interface wich will be implemented in the service struct
	return &service{
		repository: rep,
		logger:     logger,
	}
}

// Add Createuser method to be implemented in the service struct
func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser") //Use logger from service struct and modify it to log the method being executed)
	// Instanciate the user struct defined in "user.go"
	// This is the user info we are giving to the current function: Createuser
	uuid, _ := uuid.NewV4() // Generates uniq user id
	id := uuid.String()     // stringify uuid
	user := User{
		ID:       id, // stringified uuid above
		Email:    email,
		Password: password,
	}

	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id) // If there are no error, log the method being run

	return "Success", nil // Return success if function ran

}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser") // Log the current method
	email, err := s.repository.GetUser(ctx, id)
	if err != nil { // If there's an error return error
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get user", id) // If there's no error, log the user ID
	return email, nil          // Return user email if no errors and nil for the error
}
