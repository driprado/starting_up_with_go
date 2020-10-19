// The code must specify that the service will be served over http
// The methods to be exposed must be turned into endpoints.
// to do this structs must be created to represent the request and response for each method.
package account

// All methods and structure items must be public, therefore they all start with capital letters
type (
	CreateUserRequest struct {
		Email    string `json: "email"`
		Password string `json: "password"`
	}

	CreateUserResponse struct {
		ok string `json:"ok"`
	}

	GetUserReques struct {
		Id string `json:"id"`
	}

	GetUserResponse struct {
		Email string `json:"email"`
	}
)
