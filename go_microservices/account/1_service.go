// Go Kit is a bundle of libraries that can be used to build microservices
// 3 Major components in Gokit
// Transport Layer
// - How microservices communicate with each other and the outside world.

// Endpoint Layer
// - Microservice functions when exposed are called endpoints
// - Each endpoint represent a single RPC (Remote Procedure Call) method. Which is exposed to the transport layer

// Service Layer
// - Service layer is the business logic of the microservice itself.

// Example - 1 A Simple User Module
// Use a web API to insert a user into a DB
// Use the web API to fech the user by using a unique ID generated in the business logic.

package account

import "context"

// Service is our interface to expose methods to the transport layer, this interface will also be used to implement busines logic
type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}
