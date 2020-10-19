package account

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
