// // Example -1: Client without Context
// package main

// import (
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	res, err := http.Get("http://127.0.2.1:8080") // Clean Get into servers "/"
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Body.Close() // Always close body
// 	if res.StatusCode != http.StatusOK {
// 		log.Fatal(res.Status)
// 	}
// 	io.Copy(os.Stdout, res.Body) // Copy response body to stdout
// }

// Example - 2: Client with Context
package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()                            // new context
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // Cancel the context after 2 seconds
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://127.0.2.1:8080", nil) // func NewRequest(method, url string, body io.Reader) (*Request, error)
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(ctx) // make request into request with context

	res, err := http.DefaultClient.Do(req) // Do(req *Request) (*Response, error) // Takes a request returns a response
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close() // Always close body
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body) // Copy response body to stdout
}

// Notes
// Contexts should be used for request related information only
// Everything you put into a context is invisible to the rest of the code, don't add here things that are gonna be heavily used outside

// HERE => https://youtu.be/LSzR0VEraWw?t=1500
