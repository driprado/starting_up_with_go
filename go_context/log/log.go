package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

const requestIDKey = 101

// Println prints message with context identifier
func Println(ctx context.Context, msg string) { // receives a context and message to print
	id, ok := ctx.Value(requestIDKey).(int64) // get value from the context to identify what is being printed
	if !ok {
		log.Println("Could not find request ID in context")
		return
	}
	log.Println("%d %s", id, msg)
}

// Decorate is a handle wrapper, it takes a handle function and returns a handle function
// Receiving a context, add value, return to context
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // gets context from request
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id) // pass generated id value to context under requestIDKey key
		f(w, r.WithContext(ctx))                       // call above function usinf the same writer "w" and request with context
	}
}
