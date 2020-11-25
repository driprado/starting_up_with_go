// // Example -1: Server without Context
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	http.HandleFunc("/", handler)                         // handle everything in / with the function handler
// 	log.Fatal(http.ListenAndServe("127.0.2.1:8080", nil)) // ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux.
// }

// func handler(w http.ResponseWriter, r *http.Request) { // As http.HandleFunc uses this handler in "/", when we "Get" from 127.0.2.1:8080 we see "Bula" in the stdout (check client code)
// 	log.Printf("handler started")
// 	defer log.Printf("handler ended")
// 	time.Sleep(5 * time.Second)
// 	fmt.Fprintln(w, "Bula")
// }

// Example -2: Server with Context
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	http.HandleFunc("/", handler)                         // handle everything in / with the function handler
// 	log.Fatal(http.ListenAndServe("127.0.2.1:8080", nil)) // ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux.
// }

// func handler(w http.ResponseWriter, r *http.Request) { // As http.HandleFunc uses this handler in "/", when we "Get" from 127.0.2.1:8080 we see "Bula" in the stdout (check client code)
// 	ctx := r.Context() // Creates a context for request
// 	log.Printf("handler started")
// 	defer log.Printf("handler ended")

// 	select {
// 	case <-time.After(5 * time.Second):
// 		fmt.Fprintln(w, "Bula")
// 	case <-ctx.Done():
// 		err := ctx.Err()
// 		log.Print(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)

// 	}
// }

// Example -3: Using the log package with context
package main

import (
	"fmt"
	"github.com/driprado/starting_up_with_go/go_context/log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))       // wrap the handler in the Decorate function
	panic(http.ListenAndServe("127.0.2.1:8080", nil)) // Chanched fatal for panic here
}

func handler(w http.ResponseWriter, r *http.Request) { // As http.HandleFunc uses this handler in "/", when we "Get" from 127.0.2.1:8080 we see "Bula" in the stdout (check client code)
	ctx := r.Context()                  // Creates a context for request
	log.Println(ctx, "handler started") //log.Print now requires ctx
	defer log.Println(ctx, "handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Bula")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}
