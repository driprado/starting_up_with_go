package account

// Create the actual server to serve out services

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer creates an actual server to serve our services
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()    // like a http gateway
	r.Use(commonMiddleware) // handler is in commonMiddleware

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer( // NewServer constructs a new server, which implements http.Handler and wraps the provided endpoint.
		endpoints.CreateUser,
		decodeUserReq,  // json decoder function defined in `5_reqresp.go`
		encodeResponse, // json encoder function defined in `5_reqresp.go`
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeEmailReq,
		encodeResponse,
	))

	return r
}

// commonMiddleware verifies that all requests and responses are of type json
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r) // call the next http.Handler from function above
	})
}
